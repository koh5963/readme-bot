package llm

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

// TODO: langchaingoなどでOPENAI互換のLLM以外にも対応していく
func CallLlm(rule string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", errors.New("missing OPENAI_API_KEY environment variable")
	}
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini, // TODO: 将来的には環境変数に設定可とする
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: rule, // TODO: プロンプトテンプレートをJSONで用意する（モノレポ対応準備）
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("api call is failed %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}
