package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/hoangmirs/go-scraper/forms"
	"github.com/hoangmirs/go-scraper/tests"

	"github.com/beego/beego/v2/server/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SessionController", func() {
	AfterEach(func() {
		tests.ClearUserTable()
	})

	Describe("GET", func() {
		It("returns status OK", func() {
			request, _ := http.NewRequest("GET", "/login", nil)
			response := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(response, request)

			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("has body data", func() {
			request, _ := http.NewRequest("GET", "/login", nil)
			response := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(response, request)

			Expect(response.Body.Len()).To(BeNumerically(">", 0))
		})
	})

	Describe("POST", func() {
		Context("given valid params", func() {
			It("returns status FOUND", func() {
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

				request, _ := http.NewRequest("POST", "/login", body)
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

				response := httptest.NewRecorder()
				web.BeeApp.Handlers.ServeHTTP(response, request)

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
				request, _ := http.NewRequest("POST", "/login", body)
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				response := httptest.NewRecorder()
				web.BeeApp.Handlers.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("returns error flash message", func() {
				form := url.Values{
					"email":    {""},
					"password": {""},
				}
				body := strings.NewReader(form.Encode())
				request, _ := http.NewRequest("POST", "/login", body)
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				response := httptest.NewRecorder()
				web.BeeApp.Handlers.ServeHTTP(response, request)
				flashMessage := tests.GetFlash(response.Result().Cookies())

				Expect(flashMessage.Data["error"]).To(Equal("Email can not be empty"))
			})
		})
	})
})
