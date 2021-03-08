package controllers_test

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/hoangmirs/go-scraper/forms"
	. "github.com/hoangmirs/go-scraper/tests/custom_matchers"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SessionController", func() {
	AfterEach(func() {
		TruncateTables("user")
	})

	Describe("GET", func() {
		It("returns status OK", func() {
			response := MakeRequest("GET", "/login", nil)

			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("has body data", func() {
			response := MakeRequest("GET", "/login", nil)

			Expect(response).To(RenderTemplate("session#get"))
		})
	})

	Describe("POST", func() {
		Context("given valid params", func() {
			It("returns status FOUND", func() {
				// TODO : Using fabricator
				registrationForm := forms.RegistrationForm{
					Email:                "hoang@nimblehq.co",
					Password:             "123456",
					PasswordConfirmation: "123456",
				}
				_, _ = registrationForm.CreateUser()

				form := url.Values{
					"email":    {"hoang@nimblehq.co"},
					"password": {"123456"},
				}
				body := strings.NewReader(form.Encode())
				response := MakeRequest("POST", "/login", body)

				Expect(response.Code).To(Equal(http.StatusFound))
			})
		})

		Context("given invalid params", func() {
			It("returns status OK", func() {
				form := url.Values{
					"email":    {""},
					"password": {""},
				}
				body := strings.NewReader(form.Encode())
				response := MakeRequest("POST", "/login", body)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("returns error flash message", func() {
				form := url.Values{
					"email":    {""},
					"password": {""},
				}
				body := strings.NewReader(form.Encode())
				response := MakeRequest("POST", "/login", body)

				flashMessage := GetFlash(response.Result().Cookies())

				Expect(flashMessage.Data["error"]).To(Equal("Email cannot be empty"))
			})
		})
	})

	Describe("DELETE", func() {
		It("returns status Found", func() {
			response := MakeRequest("GET", "/logout", nil)

			Expect(response.Code).To(Equal(http.StatusFound))
		})

		It("returns error flash message", func() {
			response := MakeRequest("GET", "/logout", nil)
			flashMessage := GetFlash(response.Result().Cookies())

			Expect(flashMessage.Data["success"]).To(Equal("Logging out successfully"))
		})
	})
})
