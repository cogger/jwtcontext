package jwtcontext

import (
	"bytes"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetUserID", func() {
	It("should panic when no context is added", func() {
		Expect(func() { GetUserID(ctx) }).To(Panic())
	})

	It("should return blank when the token is not valid", func() {

		req, err := http.NewRequest("GET", "/", &bytes.Buffer{})
		Expect(err).NotTo(HaveOccurred())

		// req.Header.Set(http.CanonicalHeaderKey("Authorization"), fmt.Sprintf("Bearer %s", "badtoken"))

		req = authorizeRequest(req, "badtoken")
		token := req.Header.Get(http.CanonicalHeaderKey("Authorization"))
		req.Header.Set(http.CanonicalHeaderKey("Authorization"), token+"a")
		ctx := jwtTestContext("badtoken")(ctx, req)

		Expect(GetUserID(ctx)).To(Equal(""))

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
