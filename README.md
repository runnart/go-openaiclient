# OpenAI Golang SDK

This is an unofficial Golang SDK for the OpenAI API. It provides a simple and easy-to-use way to interact with the OpenAI API using Golang.

## Installation

To install the OpenAI Golang SDK, simply run:

```sh
go get github.com/runnart/openaiclient
```

## Usage

To use the OpenAI Golang SDK, you will need to have an API key for the OpenAI API. You can obtain an API key by [creating an account](https://beta.openai.com/signup/) on the OpenAI website.

Once you have your API key, you can create a new OpenAI client and start making requests to the API. Here's an example:

```go
package main

import (
	"context"
	"fmt"

	"github.com/runnart/openaiclient"
)

func main() {
	// Replace SERVER_URL with your actual SERVER_URL key
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


```

This example generates a text completion using the OpenAI API, based on a prompt of "Once upon a time,". The resulting completed text is printed to the console.

## Contributing

This SDK is open source, and contributions are welcome! If you have any bug reports, feature requests, or patches, please submit them through the GitHub issue tracker and pull request system.

## Licence

This SDK is released under the MIT License. See [LICENSE](https://github.com/runnart/openaiclient/LICENSE) for details.
