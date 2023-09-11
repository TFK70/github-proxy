package cache

import (
	"fmt"
	"github-proxy/pkg/utils"
	"os"
)

func CacheEnv(service string, content string) {
	if _, err := os.Stat("/cache"); os.IsNotExist(err) {
		os.Mkdir("./cache", 0644)
	}

	err := os.WriteFile(fmt.Sprintf("./cache/%s", service), []byte(content), 0644)
	utils.Check(err)
}

func GetCachedEnv(service string) []byte {
	file, err := os.ReadFile(fmt.Sprintf("./cache/%s", service))
	utils.Check(err)
	return file
}
