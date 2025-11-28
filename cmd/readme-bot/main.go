package main

import (
	"fmt"

	"github.com/koh5963/readme-bot/internal/rules"
)

func main() {
	fmt.Println("Hello, Readme Bot!")
	rules, err := rules.LoadRules()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rules)
}
