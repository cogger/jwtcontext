package jwtcontext

import "golang.org/x/net/context"

//GetClaims gets the claims associated with a token
func GetUserID(ctx context.Context) string {
	jwt := jc(ctx)

	return jwt.Claims["sub"].(string)
}
