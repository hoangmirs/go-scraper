package apiv1controllers_test

import (
	"net/http"

	. "github.com/hoangmirs/go-scraper/tests/custom_matchers"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordController", func() {
	AfterEach(func() {
		TruncateTables("user", "oauth2_clients", "oauth2_tokens")
	})

	Describe("POST", func() {
		Context("given an authenticated request", func() {
			Context("given a valid file", func() {
				It("returns status Created", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					token := fabricators.FabricateToken(user)
					userInfo := &UserInfo{
						Token: token,
					}
					headers, body := CreateMultipartRequestInfo("tests/fixtures/files/valid.csv", "text/csv")

					response := MakeAuthenticatedRequest("POST", "/api/v1/keywords", headers, body, userInfo)

					Expect(response.Code).To(Equal(http.StatusCreated))
				})

				It("returns the empty body", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					token := fabricators.FabricateToken(user)
					userInfo := &UserInfo{
						Token: token,
					}
					headers, body := CreateMultipartRequestInfo("tests/fixtures/files/valid.csv", "text/csv")

					response := MakeAuthenticatedRequest("POST", "/api/v1/keywords", headers, body, userInfo)

					Expect(response.Body.String()).To(BeEmpty())
				})
			})

			Context("given an invalid file", func() {
				It("returns status UnprocessableEntity", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					token := fabricators.FabricateToken(user)
					userInfo := &UserInfo{
						Token: token,
					}
					headers, body := CreateMultipartRequestInfo("tests/fixtures/files/text.txt", "text/plain")

					response := MakeAuthenticatedRequest("POST", "/api/v1/keywords", headers, body, userInfo)

					Expect(response.Code).To(Equal(http.StatusUnprocessableEntity))
				})

				It("returns correct response", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					token := fabricators.FabricateToken(user)
					userInfo := &UserInfo{
						Token: token,
					}
					headers, body := CreateMultipartRequestInfo("tests/fixtures/files/text.txt", "text/plain")

					response := MakeAuthenticatedRequest("POST", "/api/v1/keywords", headers, body, userInfo)

					Expect(response).To(MatchJSONSchema("error/json_api"))
				})
			})

			Context("given NO file", func() {
				It("returns status UnprocessableEntity", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					token := fabricators.FabricateToken(user)
					userInfo := &UserInfo{
						Token: token,
					}

					response := MakeAuthenticatedRequest("POST", "/api/v1/keywords", nil, nil, userInfo)

					Expect(response.Code).To(Equal(http.StatusUnprocessableEntity))
				})

				It("returns correct response", func() {
					user := fabricators.FabricateUser(faker.Email(), faker.Password())
					token := fabricators.FabricateToken(user)
					userInfo := &UserInfo{
						Token: token,
					}

					response := MakeAuthenticatedRequest("POST", "/api/v1/keywords", nil, nil, userInfo)

					Expect(response).To(MatchJSONSchema("error/json_api"))
				})
			})
		})

		Context("given an unauthenticated request", func() {
			It("returns status Unauthorized", func() {
				response := MakeRequest("POST", "/api/v1/keywords", nil)

				Expect(response.Code).To(Equal(http.StatusUnauthorized))
			})
		})
	})
})
