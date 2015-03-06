package jwt

import "golang.org/x/net/context"

//IsLoggedIn returns if a jwt token is valid
func IsLoggedIn(ctx context.Context) bool {
	return jc(ctx).Valid
}
