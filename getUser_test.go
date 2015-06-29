package jwtcontext

import (
	"bytes"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetUser", func() {
	It("should panic when no context is added", func() {
		Expect(func() {
			var u user
			GetUser(ctx, &u)
		}).To(Panic())
	})

	It("should load the user model", func() {
		for _, email := range emails {
			req, err := http.NewRequest("GET", "/", &bytes.Buffer{})
			Expect(err).NotTo(HaveOccurred())

			req = authorizeRequest(req, email)
			ctx := jwtTestContext(email)(ctx, req)
			var u user
			err = GetUser(ctx, &u)
			Expect(err).NotTo(HaveOccurred())
			Expect(u.Email).To(Equal(email))
		}
	})

	It("should return an error when the request is not authorized", func() {
		for _, email := range emails {
			req, err := http.NewRequest("GET", "/", &bytes.Buffer{})
			Expect(err).NotTo(HaveOccurred())

			ctx := jwtTestContext(email)(ctx, req)
			var u user
			err = GetUser(ctx, &u)
			Expect(err).To(HaveOccurred())
			Expect(u.Email).To(BeEmpty())
		}
	})
})
