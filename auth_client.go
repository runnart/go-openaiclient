package go_openaiclient

import (
	"fmt"
	"net/http"
)

// bearerAuthClient implements proxy http client with bearer authorization.
type bearerAuthClient struct {
	token string
	http  *http.Client
	*Client
}

// NewBearerAuthClient creates new client ready for doing requests with bearer token.
func NewBearerAuthClient(serverURL, token string) (*bearerAuthClient, error) {
	authClient := &bearerAuthClient{
		token: token,
		http:  &http.Client{},
	}

	client, err := NewClient(serverURL, WithClient(authClient))
	if err != nil {
		return nil, err
	}
	authClient.Client = client

	return authClient, nil
}

func (b *bearerAuthClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", b.token))
	return b.http.Do(req)
}
