package github

import (
	"context"
	"errors"
	"fmt"

	github "github.com/google/go-github/v79/github"
	"github.com/koh5963/readme-bot/internal/models/common"
	"golang.org/x/oauth2"
)

// GetDiff gets GitHub diff from Pull Request and it returns diff string.
// info is GitHub access info struct.
func GetDiff(info common.GitHubAccessInfo) (string, error) {
	ctx := context.Background()
	client, conErr := createGithubClient(info, ctx)
	if conErr != nil {
		return "", conErr
	}

	diff, diffErr := getPullRequestDiff(info, client, ctx)
	if diffErr != nil {
		return "", diffErr
	}

	return diff, nil
}

func createGithubClient(info common.GitHubAccessInfo, ctx context.Context) (*github.Client, error) {
	token := info.Token
	if token == "" {
		return nil, errors.New("missing GITHUB_TOKEN")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client, nil
}

func getPullRequestDiff(info common.GitHubAccessInfo,
	client *github.Client,
	ctx context.Context) (string, error) {

	diff, _, err := client.PullRequests.GetRaw(
		ctx,
		info.Owner,
		info.Repo,
		info.Number,
		github.RawOptions{Type: github.Diff},
	)

	if err != nil {
		return "", fmt.Errorf("get diff failed: %w", err)
	}
	return diff, nil
}
