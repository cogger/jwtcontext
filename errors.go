package jwt

import "errors"

//ErrNotAuthenticated is passed when jwt failed to authenticate
var ErrNotAuthenticated = errors.New("jwt token is invalid")

//ErrNoJWTContext is passed when the jwt context is not added
var ErrNoJWTContext = errors.New("jwt content not added")

//ErrMalformedJWT is passed when a jwt is malformed
var ErrMalformedJWT = errors.New("jwt token is malformed")
