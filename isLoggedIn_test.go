package jwtcontext_test

import (
	"bytes"
	"net/http"

	. "github.com/cogger/jwtcontext"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IsLoggedIn", func() {
	It("should panic when no context is added", func() {
		Expect(func() { IsLoggedIn(ctx) }).To(Panic())
	})
	It("should return true when the request is authorized", func() {
		for _, email := range emails {
			req, err := http.NewRequest("GET", "/", &bytes.Buffer{})
			Expect(err).NotTo(HaveOccurred())

			req = authorizeRequest(req, email)
			ctx := jwtTestContext(email)(ctx, req)

			Expect(IsLoggedIn(ctx)).To(BeTrue())
		}
	})

	It("should return false when the request is not authorized", func() {
		for _, email := range emails {
			req, err := http.NewRequest("GET", "/", &bytes.Buffer{})
			Expect(err).NotTo(HaveOccurred())

			ctx := jwtTestContext(email)(ctx, req)

			Expect(IsLoggedIn(ctx)).To(BeFalse())
		}
	})
})
