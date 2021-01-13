package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/hoangmirs/go-scraper/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegistrationController", func() {
	Describe("GET", func() {
		It("returns status OK", func() {
			request, _ := http.NewRequest("GET", "/register", nil)
			response := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(response, request)

			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("has body data", func() {
			request, _ := http.NewRequest("GET", "/register", nil)
			response := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(response, request)

			Expect(response.Body.Len()).To(BeNumerically(">", 0))
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

				request, _ := http.NewRequest("POST", "/register", body)
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

				response := httptest.NewRecorder()
				web.BeeApp.Handlers.ServeHTTP(response, request)

				logs.Trace("Code[%d]\n%s", response.Code, response.Body.String())
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
				request, _ := http.NewRequest("POST", "/register", body)
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				response := httptest.NewRecorder()
				web.BeeApp.Handlers.ServeHTTP(response, request)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("returns error flash message", func() {
				form := url.Values{
					"email":                 {""},
					"password":              {""},
					"password_confirmation": {""},
				}
				body := strings.NewReader(form.Encode())
				request, _ := http.NewRequest("POST", "/register", body)
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				response := httptest.NewRecorder()
				web.BeeApp.Handlers.ServeHTTP(response, request)
				flashMessage := tests.GetFlash(response.Result().Cookies())

				Expect(flashMessage.Data["error"]).To(Equal("Email Can not be empty"))
			})
		})
	})
})
