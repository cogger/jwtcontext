package jwtcontext

import "golang.org/x/net/context"

//Server interface that defines how a jwt token accesses the server for client id keys and gets users
type Server interface {
	GetClient(context.Context, string) ([]byte, error)
	GetUser(context.Context, string, interface{}) error
}
