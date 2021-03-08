package controllers_test

import (
	"net/http"

	. "github.com/hoangmirs/go-scraper/tests/custom_matchers"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordController", func() {
	AfterEach(func() {
		TruncateTables("user", "keyword")
	})

	Describe("GET", func() {
		It("returns status OK", func() {
			userInfo := &UserInfo{
				Email:    faker.Email(),
				Password: faker.Password(),
			}

			response := MakeAuthenticatedRequest("GET", "/keyword", nil, nil, userInfo)

			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("has body data", func() {
			userInfo := &UserInfo{
				Email:    faker.Email(),
				Password: faker.Password(),
			}

			response := MakeAuthenticatedRequest("GET", "/keyword", nil, nil, userInfo)

			Expect(response).To(RenderTemplate("keyword#get"))
		})
	})

	Describe("POST", func() {
		Context("given valid params", func() {
			It("returns status FOUND", func() {
				userInfo := &UserInfo{
					Email:    faker.Email(),
					Password: faker.Password(),
				}

				headers, body := CreateMultipartRequestInfo("tests/fixtures/files/valid.csv", "text/csv")

				response := MakeAuthenticatedRequest("POST", "/keyword", headers, body, userInfo)

				Expect(response.Code).To(Equal(http.StatusFound))
			})

			It("displays a success flash message", func() {
				userInfo := &UserInfo{
					Email:    faker.Email(),
					Password: faker.Password(),
				}

				headers, body := CreateMultipartRequestInfo("tests/fixtures/files/valid.csv", "text/csv")

				response := MakeAuthenticatedRequest("POST", "/keyword", headers, body, userInfo)

				flashMessage := GetFlash(response.Result().Cookies())
				Expect(flashMessage.Data["success"]).To(Equal("Processing uploaded keywords"))
			})
		})

		Context("given invalid params", func() {
			It("returns status OK", func() {
				userInfo := &UserInfo{
					Email:    faker.Email(),
					Password: faker.Password(),
				}

				headers, body := CreateMultipartRequestInfo("tests/fixtures/files/text.txt", "text/plain")

				response := MakeAuthenticatedRequest("POST", "/keyword", headers, body, userInfo)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("displays an error flash message", func() {
				userInfo := &UserInfo{
					Email:    faker.Email(),
					Password: faker.Password(),
				}

				headers, body := CreateMultipartRequestInfo("tests/fixtures/files/text.txt", "text/plain")

				response := MakeAuthenticatedRequest("POST", "/keyword", headers, body, userInfo)

				flashMessage := GetFlash(response.Result().Cookies())
				Expect(flashMessage.Data["error"]).To(Equal("File type is not supported"))
			})
		})
	})
})
