// Package common defines shared data models.
package common

// GithubAccessInfo represents GitHub repository access information.
type GitHubAccessInfo struct {
	Owner  string
	Repo   string
	Number int    // PR number
	Token  string // GITHUB_TOKEN
}
