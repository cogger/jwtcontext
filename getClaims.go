package jwt

import "golang.org/x/net/context"

//GetClaims gets the claims associated with a token
func GetClaims(ctx context.Context) map[string]interface{} {
	return jc(ctx).Claims
}
