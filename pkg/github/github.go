package github

import (
  "fmt"
	"context"
	"encoding/base64"
	"github-proxy/pkg/config"
	"github.com/google/go-github/v54/github"
	"golang.org/x/oauth2"
)

type GetFileContent = func(path string) (string, error)

func CreateRepoCrawler(owner string, repo string) GetFileContent {
	ctx := context.Background()
	token := config.Config.Token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	getFileContent := func(path string) (string, error) {
		content, _, _, err := client.Repositories.GetContents(ctx, owner, repo, path, nil)

    if err != nil {
      fmt.Println("Error occured while requesting file content:")
      fmt.Println(err)
      return "", err
    }

		decoded, _ := base64.StdEncoding.DecodeString(*content.Content)

		return string(decoded), nil
	}

	return getFileContent
}
