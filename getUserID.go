package jwtcontext

import "golang.org/x/net/context"

//GetUserID gets the userID associated with a token
func GetUserID(ctx context.Context) string {
	jwt := jc(ctx)
	if !jwt.Valid {
		return ""
	}
	return jwt.Claims["sub"].(string)
}
