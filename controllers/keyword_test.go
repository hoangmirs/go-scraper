package controllers_test

import (
	"fmt"
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
		TruncateTables("user", "keyword")
	})

	Describe("GET", func() {
		Context("given NO parameters", func() {
			It("returns status OK", func() {
				userInfo := &UserInfo{
					Email:    faker.Email(),
					Password: faker.Password(),
				}

				response := MakeAuthenticatedRequest("GET", "/keyword", nil, nil, userInfo)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("renders keyword#get template", func() {
				userInfo := &UserInfo{
					Email:    faker.Email(),
					Password: faker.Password(),
				}

				response := MakeAuthenticatedRequest("GET", "/keyword", nil, nil, userInfo)

				Expect(response).To(RenderTemplate("keyword#get"))
			})

			It("returns correct records", func() {
				email := faker.Email()
				password := faker.Password()
				user := fabricators.FabricateUser(email, password)
				userInfo := &UserInfo{
					Id:       user.Id,
					Email:    email,
					Password: password,
				}
				savedKeyword := fabricators.FabricateKeyword("EXPECTED_KEYWORD", user)

				response := MakeAuthenticatedRequest("GET", "/keyword", nil, nil, userInfo)

				body := GetBody(response)

				Expect(body).To(ContainSubstring(savedKeyword.Keyword))
			})
		})

		Context("given keyword parameter", func() {
			Context("given a blank keyword parameter", func() {
				It("returns status OK", func() {
					userInfo := &UserInfo{
						Email:    faker.Email(),
						Password: faker.Password(),
					}

					response := MakeAuthenticatedRequest("GET", "/keyword?keyword=", nil, nil, userInfo)

					Expect(response.Code).To(Equal(http.StatusOK))
				})

				It("renders keyword#get template", func() {
					userInfo := &UserInfo{
						Email:    faker.Email(),
						Password: faker.Password(),
					}

					response := MakeAuthenticatedRequest("GET", "/keyword?keyword=", nil, nil, userInfo)

					Expect(response).To(RenderTemplate("keyword#get"))
				})

				It("returns all records", func() {
					email := faker.Email()
					password := faker.Password()
					user := fabricators.FabricateUser(email, password)
					userInfo := &UserInfo{
						Id:       user.Id,
						Email:    email,
						Password: password,
					}
					savedKeyword1 := fabricators.FabricateKeyword("EXPECTED_KEYWORD1", user)
					savedKeyword2 := fabricators.FabricateKeyword("EXPECTED_KEYWORD2", user)

					response := MakeAuthenticatedRequest("GET", "/keyword?keyword=", nil, nil, userInfo)

					body := GetBody(response)

					Expect(body).To(ContainSubstring(savedKeyword1.Keyword))
					Expect(body).To(ContainSubstring(savedKeyword2.Keyword))
				})
			})

			Context("given a valid keyword", func() {
				It("returns status OK", func() {
					userInfo := &UserInfo{
						Email:    faker.Email(),
						Password: faker.Password(),
					}

					response := MakeAuthenticatedRequest("GET", "/keyword?keyword=EXPECTED_KEYWORD", nil, nil, userInfo)

					Expect(response.Code).To(Equal(http.StatusOK))
				})

				It("renders keyword#get template", func() {
					userInfo := &UserInfo{
						Email:    faker.Email(),
						Password: faker.Password(),
					}

					response := MakeAuthenticatedRequest("GET", "/keyword?keyword=EXPECTED_KEYWORD", nil, nil, userInfo)

					Expect(response).To(RenderTemplate("keyword#get"))
				})

				It("returns only filtered records", func() {
					email := faker.Email()
					password := faker.Password()
					user := fabricators.FabricateUser(email, password)
					userInfo := &UserInfo{
						Id:       user.Id,
						Email:    email,
						Password: password,
					}
					savedKeyword1 := fabricators.FabricateKeyword("EXPECTED_KEYWORD1", user)
					savedKeyword2 := fabricators.FabricateKeyword("EXPECTED_KEYWORD2", user)

					response := MakeAuthenticatedRequest("GET", "/keyword?keyword=EXPECTED_KEYWORD1", nil, nil, userInfo)

					body := GetBody(response)

					Expect(body).To(ContainSubstring(savedKeyword1.Keyword))
					Expect(body).NotTo(ContainSubstring(savedKeyword2.Keyword))
				})
			})

			Context("given an invalid keyword", func() {
				It("returns status OK", func() {
					userInfo := &UserInfo{
						Email:    faker.Email(),
						Password: faker.Password(),
					}

					response := MakeAuthenticatedRequest("GET", "/keyword?keyword=INVALID", nil, nil, userInfo)

					Expect(response.Code).To(Equal(http.StatusOK))
				})

				It("renders keyword#get template", func() {
					userInfo := &UserInfo{
						Email:    faker.Email(),
						Password: faker.Password(),
					}

					response := MakeAuthenticatedRequest("GET", "/keyword?keyword=EXPECTED_KEYWORD", nil, nil, userInfo)

					Expect(response).To(RenderTemplate("keyword#get"))
				})

				It("returns NO records", func() {
					email := faker.Email()
					password := faker.Password()
					user := fabricators.FabricateUser(email, password)
					userInfo := &UserInfo{
						Id:       user.Id,
						Email:    email,
						Password: password,
					}
					_ = fabricators.FabricateKeyword("EXPECTED_KEYWORD", user)

					response := MakeAuthenticatedRequest("GET", "/keyword?keyword=INVALID_KEYWORD", nil, nil, userInfo)

					body := GetBody(response)

					Expect(body).To(ContainSubstring("No keywords were found"))
				})
			})
		})
	})

	Describe("GET /:id", func() {
		Context("given a valid keyword", func() {
			It("returns status OK", func() {
				email := faker.Email()
				password := faker.Password()
				user := fabricators.FabricateUser(email, password)
				userInfo := &UserInfo{
					Id:       user.Id,
					Email:    email,
					Password: password,
				}
				savedKeyword := fabricators.FabricateKeyword("keyword", user)
				keywordDetailPath := fmt.Sprintf("/keyword/%d", savedKeyword.Id)

				response := MakeAuthenticatedRequest("GET", keywordDetailPath, nil, nil, userInfo)

				Expect(response.Code).To(Equal(http.StatusOK))
			})

			It("renders keyword#show template", func() {
				email := faker.Email()
				password := faker.Password()
				user := fabricators.FabricateUser(email, password)
				userInfo := &UserInfo{
					Id:       user.Id,
					Email:    email,
					Password: password,
				}
				savedKeyword := fabricators.FabricateKeyword("keyword", user)
				keywordDetailPath := fmt.Sprintf("/keyword/%d", savedKeyword.Id)

				response := MakeAuthenticatedRequest("GET", keywordDetailPath, nil, nil, userInfo)

				Expect(response).To(RenderTemplate("keyword#show"))
			})
		})

		Context("given an invalid keyword", func() {
			It("returns status NotFound", func() {
				email := faker.Email()
				password := faker.Password()
				userInfo := &UserInfo{
					Email:    email,
					Password: password,
				}
				keywordDetailPath := fmt.Sprintf("/keyword/%s", "invalid")

				response := MakeAuthenticatedRequest("GET", keywordDetailPath, nil, nil, userInfo)

				Expect(response.Code).To(Equal(http.StatusNotFound))
			})
		})
	})

	Describe("GET /:id/html", func() {
		Context("given a valid keyword", func() {
			It("returns status OK", func() {
				email := faker.Email()
				password := faker.Password()
				user := fabricators.FabricateUser(email, password)
				userInfo := &UserInfo{
					Id:       user.Id,
					Email:    email,
					Password: password,
				}
				savedKeyword := fabricators.FabricateKeyword("keyword", user)
				keywordDetailPath := fmt.Sprintf("/keyword/%d/html", savedKeyword.Id)

				response := MakeAuthenticatedRequest("GET", keywordDetailPath, nil, nil, userInfo)

				Expect(response.Code).To(Equal(http.StatusOK))
			})
		})

		Context("given an invalid keyword", func() {
			It("returns status NotFound", func() {
				email := faker.Email()
				password := faker.Password()
				userInfo := &UserInfo{
					Email:    email,
					Password: password,
				}
				keywordDetailPath := fmt.Sprintf("/keyword/%s/html", "invalid")

				response := MakeAuthenticatedRequest("GET", keywordDetailPath, nil, nil, userInfo)

				Expect(response.Code).To(Equal(http.StatusNotFound))
			})
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
