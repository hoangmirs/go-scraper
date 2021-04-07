package apiv1controllers_test

import (
	"net/http"
	"net/url"
	"strings"

	. "github.com/hoangmirs/go-scraper/tests/custom_matchers"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserController", func() {
	AfterEach(func() {
		TruncateTables("user", "oauth2_clients", "oauth2_tokens")
	})

	Describe("POST", func() {
		Context("given valid credentials", func() {
			Context("given a valid user information", func() {
				It("returns status Created", func() {
					client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					email := faker.Email()
					password := faker.Password()

					form := url.Values{
						"client_id":             {client.ID},
						"client_secret":         {client.Secret},
						"email":                 {email},
						"password":              {password},
						"password_confirmation": {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/users", body)

					Expect(response.Code).To(Equal(http.StatusCreated))
				})

				It("returns empty response", func() {
					client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					email := faker.Email()
					password := faker.Password()

					form := url.Values{
						"client_id":             {client.ID},
						"client_secret":         {client.Secret},
						"email":                 {email},
						"password":              {password},
						"password_confirmation": {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/users", body)

					Expect(response.Body.String()).To(BeEmpty())
				})

				It("creates a new user record", func() {
					client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					email := faker.Email()
					password := faker.Password()

					form := url.Values{
						"client_id":             {client.ID},
						"client_secret":         {client.Secret},
						"email":                 {email},
						"password":              {password},
						"password_confirmation": {password},
					}
					body := strings.NewReader(form.Encode())

					_ = MakeRequest("POST", "/api/v1/users", body)

					userCount, err := orm.NewOrm().QueryTable("user").Count()
					if err != nil {
						Fail("Failed to count users: " + err.Error())
					}

					Expect(userCount).To(BeNumerically("==", 1))
				})
			})

			Context("given an existing user information", func() {
				It("returns status UnprocessableEntity", func() {
					client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					email := faker.Email()
					password := faker.Password()
					_ = fabricators.FabricateUser(email, password)

					form := url.Values{
						"client_id":             {client.ID},
						"client_secret":         {client.Secret},
						"email":                 {email},
						"password":              {password},
						"password_confirmation": {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/users", body)

					Expect(response.Code).To(Equal(http.StatusUnprocessableEntity))
				})

				It("returns correct response", func() {
					client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					email := faker.Email()
					password := faker.Password()
					_ = fabricators.FabricateUser(email, password)

					form := url.Values{
						"client_id":             {client.ID},
						"client_secret":         {client.Secret},
						"email":                 {email},
						"password":              {password},
						"password_confirmation": {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/users", body)

					Expect(response).To(MatchJSONSchema("error/json_api"))
				})
			})

			Context("given an invalid user information", func() {
				It("returns status UnprocessableEntity", func() {
					client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())

					form := url.Values{
						"client_id":             {client.ID},
						"client_secret":         {client.Secret},
						"email":                 {"email"},
						"password":              {"password"},
						"password_confirmation": {""},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/users", body)

					Expect(response.Code).To(Equal(http.StatusUnprocessableEntity))
				})

				It("returns correct response", func() {
					client := fabricators.FabricateOAuthClient(uuid.New().String(), uuid.New().String())

					form := url.Values{
						"client_id":             {client.ID},
						"client_secret":         {client.Secret},
						"email":                 {"email"},
						"password":              {"password"},
						"password_confirmation": {""},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/users", body)

					Expect(response).To(MatchJSONSchema("error/json_api"))
				})
			})
		})

		Context("given INVALID credentials", func() {
			It("returns status Unauthorized", func() {
				form := url.Values{
					"client_id":     {"invalid"},
					"client_secret": {"invalid"},
				}
				body := strings.NewReader(form.Encode())

				response := MakeRequest("POST", "/api/v1/users", body)

				Expect(response.Code).To(Equal(http.StatusUnauthorized))
			})
		})
	})
})
