package jwt

import (
	"net/http"
	"sync"

	"golang.org/x/net/context"
)

type jwtContextKey struct{}

//Add adds jwt processing to a context
func Add(server Server) func(context.Context, *http.Request) context.Context {
	return func(ctx context.Context, req *http.Request) context.Context {
		auth := req.Header.Get(http.CanonicalHeaderKey("Authorization"))

		return context.WithValue(ctx, jwtContextKey{}, jwtContext{
			Auth:   auth,
			Server: server,
			Load:   &sync.Once{},
		})
	}
}
