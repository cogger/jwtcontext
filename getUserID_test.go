package jwtcontext_test

import (
	"bytes"
	"net/http"

	. "github.com/cogger/jwtcontext"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetUserID", func() {
	It("should panic when no context is added", func() {
		Expect(func() { GetUserID(ctx) }).To(Panic())
	})
	It("should return the id of the user", func() {
		for _, email := range emails {
			req, err := http.NewRequest("GET", "/", &bytes.Buffer{})
			Expect(err).NotTo(HaveOccurred())

			req = authorizeRequest(req, email)
			ctx := jwtTestContext(email)(ctx, req)
			Expect(GetUserID(ctx)).To(Equal(email))
		}
	})
})
