package apiv1controllers_test

import (
	"net/http"

	. "github.com/hoangmirs/go-scraper/tests/test_helpers"
	. "github.com/hoangmirs/go-scraper/tests/custom_matchers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OAuthClientController", func() {
	Describe("POST", func() {
		// TODO : Revisit this to send request with basic authentication. The current implementation does NOT work
		XContext("given a request with basic authentication", func() {
			It("returns status OK", func() {
				response := MakeRequestWithBasicAuthentication("POST", "/api/v1/oauth/client", nil)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("returns correct response", func() {
				response := MakeRequestWithBasicAuthentication("POST", "/api/v1/oauth/client", nil)

				Expect(response).To(MatchJSONSchema("oauth/client/valid"))
			})
		})

		Context("given a request WITHOUT basic authentication", func() {
			It("returns status Unauthorized", func() {
				response := MakeRequest("POST", "/api/v1/oauth/client", nil)

				Expect(response.Code).To(Equal(http.StatusUnauthorized))
			})
		})
	})
})
