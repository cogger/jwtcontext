package jwtcontext

import "golang.org/x/net/context"

//GetUser gets a user from a context
func GetUser(ctx context.Context, dst interface{}) error {
	jwt := jc(ctx)
	if !jwt.Valid {
		return ErrNotAuthenticated
	}

	return jwt.Server.GetUser(ctx, jwt.Claims["sub"].(string), dst)
}
