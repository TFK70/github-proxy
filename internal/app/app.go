package app

import (
	"fmt"
	"github-proxy/pkg/cache"
	"github-proxy/pkg/config"
	"github-proxy/pkg/github"
	"github-proxy/pkg/utils"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func Reconcile() {
	owner := config.Config.Owner
	repo := config.Config.Repo
	sources := config.Config.Sources

	getFileContent := github.CreateRepoCrawler(owner, repo)

	for service, source := range sources {
		fmt.Println(fmt.Sprintf("Reconciled envs for service %s", service))
		content := getFileContent(source.Path)
		cache.CacheEnv(service, content)
	}

	fmt.Println("Reconcilation finished")
}

func GetEnvHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	utils.Check(err)

	env := cache.GetCachedEnv(string(body))
	io.WriteString(w, string(env))
}

func ReconcileHandler(w http.ResponseWriter, r *http.Request) {
	Reconcile()
}

func CreateServer() {
	http.HandleFunc("/getenv", GetEnvHandler)
	http.HandleFunc("/reconcile", ReconcileHandler)

	fmt.Println("Listening on port 3000")
	err := http.ListenAndServe(":3000", nil)
	utils.Check(err)
}

func Bootstrap() {
	config.LoadConfig()

	Reconcile()

	ticker := time.NewTicker(10 * time.Minute)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				Reconcile()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	CreateServer()
}
