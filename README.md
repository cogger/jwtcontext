# jwt [![GoDoc](https://godoc.org/github.com/cogger/jwt?status.png)](http://godoc.org/github.com/cogger/jwt)

jwt addes the ability to authenticate jwt tokens

## Usage
~~~ go
// main.go
package main

import (
	"net/http"
	"github.com/cogger/cogger"
	"github.com/cogger/jwt"
	"golang.org/x/net/context"
)

type User struct{}

func foo(ctx context.Context, w http.ResponseWriter, r *http.Request) int{
	if !jwt.IsLoggedIn(ctx){
		return http.StatusUnauthorized
	}

	var user User
	err := jwt.GetUser(ctx, &user)
	if err != nil {
		return http.StatusUnauthorized
	}

	return http.StatusOK
}


type jwtServer struct{}

func (server jwtServer) GetClient(ctx context.Context, clientID string) ([]byte, error){
	return []byte("somebytes"),nil
}

func (server jwtServer) GetUser(ctx context.Context, userID string, user interface{}) error{
	//load user
	return nil
}

func init() {
	fooHandler := cogger.NewHandler()
	fooHandler.AddContext(jwt.Add(jwtServer{}))

	fooHandler.SetHandler(foo)

  	http.Handle("/foo", fooHandler)
}

~~~
