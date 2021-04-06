package apiv1controllers_test

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/hoangmirs/go-scraper/services/oauth"
	. "github.com/hoangmirs/go-scraper/tests/custom_matchers"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OAuthTokenController", func() {
	AfterEach(func() {
		TruncateTables("user", "oauth2_clients", "oauth2_tokens")
	})

	Describe("POST", func() {
		Context("given a valid request", func() {
			It("returns status OK", func() {
				client := fabricators.FabricateOAuthClient(faker.UUIDDigit(), faker.UUIDDigit())
				email := faker.Email()
				password := faker.Password()
				_ = fabricators.FabricateUser(email, password)
				form := url.Values{
					"client_id":     {client.ID},
					"client_secret": {client.Secret},
					"grant_type":    {"password"},
					"username":      {email},
					"password":      {password},
				}
				body := strings.NewReader(form.Encode())

				response := MakeRequest("POST", "/api/v1/oauth/token", body)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("returns correct response", func() {
				client := fabricators.FabricateOAuthClient(faker.UUIDDigit(), faker.UUIDDigit())
				email := faker.Email()
				password := faker.Password()
				_ = fabricators.FabricateUser(email, password)
				form := url.Values{
					"client_id":     {client.ID},
					"client_secret": {client.Secret},
					"grant_type":    {"password"},
					"username":      {email},
					"password":      {password},
				}
				body := strings.NewReader(form.Encode())

				response := MakeRequest("POST", "/api/v1/oauth/token", body)

				Expect(response).To(MatchJSONSchema("oauth/token/valid"))
			})
		})

		Context("given an INVALID request", func() {
			Context("given a blank client information", func() {
				It("returns status Unauthorized", func() {
					email := faker.Email()
					password := faker.Password()
					_ = fabricators.FabricateUser(email, password)
					form := url.Values{
						"client_id":     {""},
						"client_secret": {""},
						"grant_type":    {"password"},
						"username":      {email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response.Code).To(Equal(http.StatusUnauthorized))
				})

				It("returns correct response", func() {
					email := faker.Email()
					password := faker.Password()
					_ = fabricators.FabricateUser(email, password)
					form := url.Values{
						"client_id":     {""},
						"client_secret": {""},
						"grant_type":    {"password"},
						"username":      {email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response).To(MatchJSONSchema("oauth/token/invalid"))
				})
			})

			Context("given an INVALID client information", func() {
				It("returns status InternalServerError", func() {
					email := faker.Email()
					password := faker.Password()
					_ = fabricators.FabricateUser(email, password)
					form := url.Values{
						"client_id":     {"invalid"},
						"client_secret": {"invalid"},
						"grant_type":    {"password"},
						"username":      {email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response.Code).To(Equal(http.StatusInternalServerError))
				})

				It("returns correct response", func() {
					email := faker.Email()
					password := faker.Password()
					_ = fabricators.FabricateUser(email, password)
					form := url.Values{
						"client_id":     {"invalid"},
						"client_secret": {"invalid"},
						"grant_type":    {"password"},
						"username":      {email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response).To(MatchJSONSchema("oauth/token/invalid"))
				})
			})

			Context("given INVALID user credentials", func() {
				It("returns status Unauthorized", func() {
					client := fabricators.FabricateOAuthClient(faker.UUIDDigit(), faker.UUIDDigit())
					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {client.Secret},
						"grant_type":    {"password"},
						"username":      {"invalid"},
						"password":      {"invalid"},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response.Code).To(Equal(http.StatusUnauthorized))
				})

				It("returns correct response", func() {
					client := fabricators.FabricateOAuthClient(faker.UUIDDigit(), faker.UUIDDigit())
					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {client.Secret},
						"grant_type":    {"password"},
						"username":      {"invalid"},
						"password":      {"invalid"},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response).To(MatchJSONSchema("oauth/token/invalid"))
				})
			})
		})
	})

	Describe("POST#Revoke", func() {
		Context("given valid credentials", func() {
			It("returns status NoContent", func() {
				client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				token := fabricators.FabricateToken(user, client)

				form := url.Values{
					"client_id":     {client.ID},
					"client_secret": {client.Secret},
					"token":         {token.GetAccess()},
				}
				body := strings.NewReader(form.Encode())

				response := MakeRequest("POST", "/api/v1/oauth/revoke", body)

				Expect(response.Code).To(Equal(http.StatusNoContent))
			})

			It("returns empty response", func() {
				client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				token := fabricators.FabricateToken(user, client)

				form := url.Values{
					"client_id":     {client.ID},
					"client_secret": {client.Secret},
					"token":         {token.GetAccess()},
				}
				body := strings.NewReader(form.Encode())

				response := MakeRequest("POST", "/api/v1/oauth/revoke", body)

				Expect(response.Body.String()).To(BeEmpty())
			})

			It("deletes the token from database", func() {
				client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				token := fabricators.FabricateToken(user, client)

				form := url.Values{
					"client_id":     {client.ID},
					"client_secret": {client.Secret},
					"token":         {token.GetAccess()},
				}
				body := strings.NewReader(form.Encode())

				_ = MakeRequest("POST", "/api/v1/oauth/revoke", body)

				previousToken, err := oauth.GetTokenStore().GetByAccess(context.Background(), token.GetAccess())

				Expect(previousToken).To(BeNil())
				Expect(err.Error()).To(Equal("sql: no rows in result set"))
			})
		})

		Context("given INVALID credentials", func() {
			Context("given an empty token", func() {
				It("returns status UnprocessableEntity", func() {
					client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {client.Secret},
						"token":         {""},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/revoke", body)

					Expect(response.Code).To(Equal(http.StatusUnprocessableEntity))
				})

				It("returns correct response", func() {
					client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {client.Secret},
						"token":         {""},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/revoke", body)

					Expect(response).To(MatchJSONSchema("error/json_api"))
				})
			})

			Context("given an INVALID client info", func() {
				It("returns status Unauthorized", func() {
					client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {"invalid"},
						"token":         {"token"},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/revoke", body)

					Expect(response.Code).To(Equal(http.StatusUnauthorized))
				})
			})
		})
	})
})
