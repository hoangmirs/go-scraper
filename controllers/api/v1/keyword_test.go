package apiv1controllers_test

import (
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/bxcodec/faker/v3"
	. "github.com/hoangmirs/go-scraper/tests/custom_matchers"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordController", func() {
	Describe("POST", func() {
		// TODO : Revisit this to send request with authentication. The current implementation does NOT work
		XContext("given an authenticated request", func() {
			Context("given a valid file", func() {
				It("returns status OK", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					token := fabricators.FabricateToken(user)
					userInfo := &UserInfo{
						Token: token,
					}
					logs.Info("User info: %v", userInfo)
					headers, body := CreateMultipartRequestInfo("tests/fixtures/files/valid.csv", "text/csv")

					response := MakeAuthenticatedRequest("POST", "/api/v1/keywords", headers, body, userInfo)

					Expect(response.Code).To(Equal(http.StatusOK))
				})

				It("returns correct response", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					token := fabricators.FabricateToken(user)
					userInfo := &UserInfo{
						Token: token,
					}
					headers, body := CreateMultipartRequestInfo("tests/fixtures/files/valid.csv", "text/csv")

					response := MakeAuthenticatedRequest("POST", "/api/v1/keywords", headers, body, userInfo)

					Expect(response).To(MatchJSONSchema("oauth/client/valid"))
				})
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
