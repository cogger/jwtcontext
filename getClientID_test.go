package jwtcontext

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetClientID", func() {
	It("should panic when no context is added", func() {
		Expect(func() { GetClientID(ctx) }).To(Panic())
	})

	It("should the clientID", func() {
		for _, email := range emails {
			req, err := http.NewRequest("GET", "/", &bytes.Buffer{})
			Expect(err).NotTo(HaveOccurred())

			req = authorizeRequest(req, email)
			ctx := jwtTestContext(email)(ctx, req)
			Expect(GetClientID(ctx)).To(Equal(clientID))
		}
	})

	It("should panic when the clientID is bad", func() {
		for _, email := range emails {
			token := jwt.New(jwt.GetSigningMethod("RS256"))
			token.Claims["cid"] = nil
			token.Claims["exp"] = time.Now().Add(60 * time.Minute)
			token.Claims["sub"] = email
			accessToken, err := token.SignedString(testPrivateKey)
			if err != nil {
				panic(err)
			}
			req, err := http.NewRequest("GET", "/", &bytes.Buffer{})
			Expect(err).NotTo(HaveOccurred())
			req.Header.Set(http.CanonicalHeaderKey("Authorization"), fmt.Sprintf("Bearer %s", accessToken))
			ctx := jwtTestContext(email)(ctx, req)
			Expect(func() { GetClientID(ctx) }).To(Panic())
		}
	})
})
