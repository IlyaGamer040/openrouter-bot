package service

import (
	"context"
	"fmt"

	"github.com/IlyaGamer040/openrouter-bot/internal/config"
	"github.com/revrost/go-openrouter"
)

type ApiRouter struct {
	client *openrouter.Client
}

func New(cfg *config.Config) *ApiRouter {
	client := openrouter.NewClient(cfg.ApiToken)
	return &ApiRouter{client: client}
}

func (a *ApiRouter) SendTextMessage(ctx context.Context, message string) string {

	messageWithText := openrouter.UserMessage(message)

	resp, err := a.client.CreateChatCompletion(
		ctx,
		openrouter.ChatCompletionRequest{
			Model: "qwen/qwq-32b:free",
			Messages: []openrouter.ChatCompletionMessage{
				messageWithText,
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)

	}

	return resp.Choices[0].Message.Content.Text
}

func (a *ApiRouter) SendImageMessage(ctx context.Context, imageUrl string, messgae string) string {

	// imageByte, mimeType, err := pkg.DownloadImageWithValidation(imageUrl)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return ""
	// }

	// baseImage := pkg.BytesToBase64(imageByte)

	// imageStr := fmt.Sprintf("data:%s;base64,%s", mimeType, baseImage)
	// fmt.Println(imageStr)

	messageWithImage := openrouter.UserMessageWithImage(messgae, imageUrl)

	resp, err := a.client.CreateChatCompletion(
		ctx,
		openrouter.ChatCompletionRequest{
			Model: "mistralai/mistral-small-3.2-24b-instruct:free",
			Messages: []openrouter.ChatCompletionMessage{
				messageWithImage,
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)

	}

	return resp.Choices[0].Message.Content.Text
}

