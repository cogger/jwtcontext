# jwtcontext 
[![GoDoc](https://godoc.org/github.com/cogger/jwtcontext?status.png)](http://godoc.org/github.com/cogger/jwtcontext)
[![Build Status](https://travis-ci.org/cogger/jwtcontext.svg?branch=master)](https://travis-ci.org/cogger/jwtcontext)
[![Coverage Status](https://coveralls.io/repos/cogger/jwtcontext/badge.svg?branch=master)](https://coveralls.io/r/cogger/jwtcontext?branch=master)  
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)

jwtcontext addes the ability to authenticate jwt tokens

## Installation

The import path for the package is *gopkg.in/cogger/jwtcontext.v1*.

To install it, run:

    go get gopkg.in/cogger/jwtcontext.v1

## Usage
~~~ go
// main.go
package main

import (
	"net/http"
	"gopkg.in/cogger/cogger.v1"
	"gopkg.in/cogger/jwtcontext.v1"
	"golang.org/x/net/context"
)

type User struct{}

func foo(ctx context.Context, w http.ResponseWriter, r *http.Request) int{
	if !jwtcontext.IsLoggedIn(ctx){
		return http.StatusUnauthorized
	}

	var user User
	err := jwtcontext.GetUser(ctx, &user)
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
	fooHandler.AddContext(jwtcontext.Add(jwtServer{}))

	fooHandler.SetHandler(foo)

  	http.Handle("/foo", fooHandler)
}

~~~
