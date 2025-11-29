package main

import (
	"fmt"

	_ "github.com/koh5963/readme-bot/internal/github"
	"github.com/koh5963/readme-bot/internal/llm"
	"github.com/koh5963/readme-bot/internal/rules"
)

func main() {
	fmt.Println("Hello Readme Bot!")
	rule, err := rules.LoadRules("readme")
	if err != nil {
		fmt.Println(err)
	}

	// LLM API CALL
	resp, llmerr := llm.CallLlm(rule)
	if llmerr != nil {
		fmt.Println(llmerr)
		return
	}
	fmt.Println(resp)
}
