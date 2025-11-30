// Package llm provides a LLM client with minimal API methods.
// This package is used by the README-generation agent to analyze GitHub diff.
package llm

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	response "github.com/koh5963/readme-bot/internal/models/readme"
	openai "github.com/sashabaranov/go-openai"
)

//go:embed templates/readme_prompt_template.txt
var Template string

// CallLLM sends a diff and rule to an LLM and returns its analysis result.
// diff is the GitHub diff context.
// rule is a custom rule used to analyze the diff.
// TODO: Support non-OpenAI compatible LLMs (e.g. via langchaingo).
func CallLLM(diff, rule string) (response.Response, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return response.Response{}, errors.New("missing OPENAI_API_KEY environment variable")
	}
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(Template, diff, rule),
				},
			},
		},
	)

	if err != nil {
		return response.Response{}, fmt.Errorf("api call is failed %w", err)
	}

	var resJson response.Response
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &resJson); err != nil {
		return response.Response{}, fmt.Errorf("failed to parse LLM JSON: %w", err)
	}
	return resJson, nil
}
