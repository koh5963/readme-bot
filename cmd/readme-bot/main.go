package main

import (
	"fmt"

	_ "github.com/koh5963/readme-bot/internal/github"
	llm "github.com/koh5963/readme-bot/internal/llm"
	rules "github.com/koh5963/readme-bot/internal/rules"
)

func main() {
	fmt.Println("Hello Readme Bot!")
	// Load RULES.md
	rule, err := rules.LoadRules("readme")
	if err != nil {
		// 汎用モード実行
		fmt.Println("rule load warning, use default rule: ", err)
	}

	// TODO: Get Github diff
	diff := ""

	// LLM API CALL
	resp, llmerr := llm.CallLlm(diff, rule) // TODO: Diff param setting
	if llmerr != nil {
		fmt.Println(llmerr)
		return
	}
	fmt.Println(resp)

	// TODO: Rewrite README.md and CHANGELOG.md
}
