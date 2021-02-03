package controllers_test

import (
	"net/http"
	"net/url"
	"strings"

	. "github.com/hoangmirs/go-scraper/tests/custom_matchers"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegistrationController", func() {
	AfterEach(func() {
		TruncateTables("user")
	})

	Describe("GET", func() {
		It("renders registration#get template", func() {
			response := MakeRequest("GET", "/register", nil)

			Expect(response).To(RenderTemplate("registration#get"))
		})
	})

	Describe("POST", func() {
		Context("given valid params", func() {
			It("returns status FOUND", func() {
				form := url.Values{
					"email":                 {"hoang.mirs@gmail.com"},
					"password":              {"123456"},
					"password_confirmation": {"123456"},
				}
				body := strings.NewReader(form.Encode())
				response := MakeRequest("POST", "/register", body)

				Expect(response.Code).To(Equal(http.StatusFound))
			})
		})

		Context("given invalid params", func() {
			It("returns status OK", func() {
				form := url.Values{
					"email":                 {""},
					"password":              {""},
					"password_confirmation": {""},
				}
				body := strings.NewReader(form.Encode())
				response := MakeRequest("POST", "/register", body)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("returns error flash message", func() {
				form := url.Values{
					"email":                 {""},
					"password":              {""},
					"password_confirmation": {""},
				}
				body := strings.NewReader(form.Encode())
				response := MakeRequest("POST", "/register", body)

				flashMessage := GetFlash(response.Result().Cookies())

				Expect(flashMessage.Data["error"]).To(Equal("Email can not be empty"))
			})
		})
	})
})
