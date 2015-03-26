package jwtcontext

import "golang.org/x/net/context"

//GetClientID gets the client Id from the ctx
func GetClientID(ctx context.Context) string {
	clientID, ok := jc(ctx).Claims["cid"].(string)
	if !ok {
		panic(ErrMalformedJWT)
	}
	return clientID
}
