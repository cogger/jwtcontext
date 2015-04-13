package jwtcontext

import (
	"strings"
	"sync"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
)

type jwtContext struct {
	Auth   string
	Server Server
	Valid  bool
	Load   *sync.Once
	Claims map[string]interface{}
}

func jc(ctx context.Context) jwtContext {
	config, ok := ctx.Value(jwtContextKey{}).(*jwtContext)
	if !ok {
		panic(ErrNoJWTContext)
	}

	config.Load.Do(func() {
		parts := strings.Split(config.Auth, " ")
		if len(parts) != 2 {
			return
		}

		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			clientID, ok := token.Claims["cid"].(string)
			if !ok {
				return nil, ErrMalformedJWT
			}
			return config.Server.GetClient(ctx, clientID)
		})
		config.Claims = token.Claims
		config.Valid = (err == nil && token.Valid)
	})

	return *config
}
