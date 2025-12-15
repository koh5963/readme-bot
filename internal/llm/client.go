// Package llm provides a LLM client with minimal API methods.
// This package is used by the README-generation agent to analyze GitHub diff.
package llm

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"os"

	"github.com/koh5963/readme-bot/internal/constants"
	openai "github.com/sashabaranov/go-openai"
)

//go:embed templates/*.txt
var templateFS embed.FS

// CallLLM sends a diff and rule to an LLM and returns its analysis result.
// diff is the GitHub diff context.
// rule is a custom rule used to analyze the diff.
// TODO: Support non-OpenAI compatible LLMs (e.g. via langchaingo).
func CallLLM(botType constants.BotType, diff, rule string) (string, error) {

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", errors.New("missing OPENAI_API_KEY environment variable")
	}

	template, err := getPromptTemplate(botType)
	if err != nil {
		return "", err
	}

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			// "ResponseFormat" settings forces LLM to Response Format Json.
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf(template, diff, rule),
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("api call failed: %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}

func getPromptTemplate(botType constants.BotType) (string, error) {
	var fileName string
	switch botType {
	case constants.BotTypeReadme:
		fileName = "templates/readme_prompt_template.txt"
	case constants.BotTypeReview:
		fileName = "templates/review_prompt_template.txt"
	default:
		return "", fmt.Errorf("unsupported bot type: %s", botType)
	}

	template, err := templateFS.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("prompt template read failed(%s): %w", fileName, err)
	}

	return string(template), nil
}
