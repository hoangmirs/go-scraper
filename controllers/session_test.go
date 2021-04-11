package controllers_test

import (
	"net/http"
	"net/url"
	"strings"

	. "github.com/hoangmirs/go-scraper/tests/custom_matchers"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
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

		It("renders session#get template", func() {
			response := MakeRequest("GET", "/login", nil)

			Expect(response).To(RenderTemplate("session#get"))
		})
	})

	Describe("POST", func() {
		Context("given valid params", func() {
			It("returns status FOUND", func() {
				email := "hoang@nimblehq.co"
				password := "123456"
				_ = fabricators.FabricateUser(email, password)

				form := url.Values{
					"email":    {email},
					"password": {password},
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
