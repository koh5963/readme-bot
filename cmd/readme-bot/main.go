// Main Package
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/koh5963/readme-bot/internal/constants"
	ghclient "github.com/koh5963/readme-bot/internal/github"
	llmclient "github.com/koh5963/readme-bot/internal/llm"
	"github.com/koh5963/readme-bot/internal/models/common"
	response "github.com/koh5963/readme-bot/internal/models/readme"
	"github.com/koh5963/readme-bot/internal/rules"
	"github.com/koh5963/readme-bot/internal/utils"
)

func main() {
	// Load RULES.md
	rule, err := rules.LoadRules("readme")
	if err != nil {
		// using general rule
		fmt.Println("rule load warning, use default rule: ", err)
	}

	// Get Github diff from Pull Request
	accessInfo, paramErr := getGitHubAccessInfo()
	if paramErr != nil {
		fmt.Println(paramErr)
		return
	}
	diff, diffErr := ghclient.GetDiff(accessInfo)
	if diffErr != nil {
		fmt.Println(diffErr)
		return
	}

	// LLM API CALL
	resp, llmErr := llmclient.CallLLM(constants.BotTypeReadme, diff, rule)
	if llmErr != nil {
		fmt.Println(llmErr)
		return
	}
	fmt.Println(resp)

	var resJson response.Response
	if err := json.Unmarshal([]byte(resp), &resJson); err != nil {
		fmt.Println(fmt.Errorf("failed to parse LLM JSON: %w", err))
		return
	}

	utils.RewriteReadme(resJson.ReadmeLatestChange, "## latest change")
}

func getGitHubAccessInfo() (common.GitHubAccessInfo, error) {
	owner := flag.String("owner", "", "repo owner")
	repo := flag.String("repo", "", "repo name")
	number := flag.Int("number", 0, "PR number")
	flag.Parse()

	token := os.Getenv("GITHUB_TOKEN")

	if *owner == "" || *repo == "" || *number == 0 {
		return common.GitHubAccessInfo{}, errors.New("invalid system parameters: owner/repo/number is required")
	}

	return common.GitHubAccessInfo{
		Owner:  *owner,
		Repo:   *repo,
		Number: *number,
		Token:  token,
	}, nil
}
