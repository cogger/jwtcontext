package jwtcontext_test

import (
	"bytes"
	"net/http"

	. "github.com/cogger/jwtcontext"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetClaims", func() {
	It("should panic when no context is added", func() {
		Expect(func() { GetClaims(ctx) }).To(Panic())
	})

	It("should return an map of interfaces", func() {
		for _, email := range emails {
			req, err := http.NewRequest("GET", "/", &bytes.Buffer{})
			Expect(err).NotTo(HaveOccurred())

			req = authorizeRequest(req, email)
			ctx := jwtTestContext(email)(ctx, req)
			claims := GetClaims(ctx)
			Expect(claims).ToNot(BeNil())

			sub, ok := claims["sub"].(string)
			Expect(ok).To(BeTrue())
			Expect(sub).To(Equal(email))
		}
	})
})
