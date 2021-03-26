package apiv1controllers_test

import (
	"net/http"

	. "github.com/hoangmirs/go-scraper/tests/custom_matchers"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HealthCheckController", func() {
	Describe("GET", func() {
		Context("given a valid request", func() {
			It("returns status OK", func() {
				response := MakeRequest("GET", "/api/v1/health_check", nil)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("matches the correct API schema", func() {
				response := MakeRequest("GET", "/api/v1/health_check", nil)

				Expect(response).To(MatchJSONSchema("health_check/valid"))
			})
		})
	})
})
