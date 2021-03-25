package apiv1controllers_test

import (
	"net/http"

	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HealthCheckController", func() {
	Describe("GET", func() {
		Context("given a valid request", func() {
			It("returns status OK", func() {
				response := MakeAuthenticatedRequest("GET", "/api/v1/health_check", nil, nil, nil)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("returns correct response", func() {
				response := MakeAuthenticatedRequest("GET", "/api/v1/health_check", nil, nil, nil)

				expectedResponse := `
				{
					"data": {
						"type": "health_check",
						"id": "0",
						"attributes": {
							"success": true
						}
					}
				}
				`

				Expect(response.Body).To(MatchJSON(expectedResponse))
			})
		})
	})
})
