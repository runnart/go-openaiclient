package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/runnart/go-openaiclient"
)

func main() {
	// Replace SERVER_URL and API_TOKEN with your actual SERVER_URL and API_TOKEN values
	client, err := openai.NewBearerAuthClient(os.Getenv("SERVER_URL"), os.Getenv("API_TOKEN"))
	if err != nil {
		panic(err)
	}

	// Generate a text completion
	prompt := "Introduce himself please"
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
