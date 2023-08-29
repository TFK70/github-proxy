package github

import (
	"context"
	"encoding/base64"
	"github-proxy/pkg/config"
	"github-proxy/pkg/utils"
	"github.com/google/go-github/v54/github"
	"golang.org/x/oauth2"
)

type GetFileContent = func(path string) string

func CreateRepoCrawler(owner string, repo string) GetFileContent {
	ctx := context.Background()
	token := config.Config.Token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	getFileContent := func(path string) string {
		content, _, _, err := client.Repositories.GetContents(ctx, owner, repo, path, nil)

		utils.Check(err)

		decoded, _ := base64.StdEncoding.DecodeString(*content.Content)

		return string(decoded)
	}

	return getFileContent
}
