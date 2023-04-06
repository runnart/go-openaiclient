package main

import (
	"context"
	"fmt"

	"github.com/runnart/openaiclient"
)

func main() {
	// Replace SERVER_URL with your actual SERVER_URL
	client, err := openaiclient.NewClient("SERVER_URL")
	if err != nil {
		panic(err)
	}

	// Generate a text completion
	prompt := "Once upon a time,"
	req := &openaiclient.CreateChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openaiclient.ChatCompletionRequestMessage{
			{
				Role:    openaiclient.ChatCompletionRequestMessageRoleUser,
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
