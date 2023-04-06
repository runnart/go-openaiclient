package openaiclient

import (
	"github.com/ogen-go/ogen/middleware"
)

// BearerTokenMiddleware is a middleware function that adds a bearer token to outgoing requests.
func BearerTokenMiddleware(token string) Middleware {
	return func(request middleware.Request, next middleware.Next) (middleware.Response, error) {
		request.Raw.Header.Set("Authorization", "Bearer "+token)
		return next(request)
	}
}

// WithClientMiddleware specifies middlewares to use.
func WithClientMiddleware(m ...Middleware) ClientOption {
	return optionFunc[clientConfig](func(_ *clientConfig) {
		middleware.ChainMiddlewares(m...)
	})
}
