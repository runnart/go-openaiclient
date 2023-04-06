package main

import (
	"context"
	"fmt"

	openai "github.com/runnart/openaiclient"
)

func main() {
	// Replace SERVER_URL and API_TOKEN with your actual SERVER_URL and API_TOKEN values
	client, err := openai.NewClient("SERVER_URL", openai.WithClientMiddleware(openai.BearerTokenMiddleware("API_TOKEN")))
	if err != nil {
		panic(err)
	}

	// Generate a text completion
	prompt := "Once upon a time,"
	req := &openai.CreateChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionRequestMessage{
			{
				Role:    openai.ChatCompletionRequestMessageRoleUser,
				Content: prompt,
			},
		},
	}
	response, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the completed text
	fmt.Println(response.Choices[0].Message.Value.Content)
}
