package models_test

import (
	"github.com/hoangmirs/go-scraper/models"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	AfterEach(func() {
		TruncateTables("user")
	})

	Describe("#CreateUser", func() {
		Context("given valid attributes", func() {
			It("returns the user ID", func() {
				user := models.User{
					Email:             "dev@nimblehq.co",
					EncryptedPassword: "hashedPassword",
				}

				userID, err := models.CreateUser(&user)
				if err != nil {
					Fail("Failed to create user: " + err.Error())
				}

				Expect(userID).To(BeNumerically(">", 0))
			})

			It("does NOT return error", func() {
				user := models.User{
					Email:             "dev@nimblehq.co",
					EncryptedPassword: "hashedPassword",
				}
				_, err := models.CreateUser(&user)

				Expect(err).To(BeNil())
			})
		})

		Context("given invalid attributes", func() {
			Context("given email that already exist in database", func() {
				It("returns an error", func() {
					_, err := fabricators.FabricateUser("dev@nimblehq.co", faker.Password())
					if err != nil {
						Fail("Failed to fabricate user: " + err.Error())
					}

					user := models.User{
						Email:             "dev@nimblehq.co",
						EncryptedPassword: "password",
					}
					userID, err := models.CreateUser(&user)

					Expect(err.Error()).To(Equal(`pq: duplicate key value violates unique constraint "user_email_key"`))
					Expect(userID).To(Equal(int64(0)))
				})
			})
		})
	})

})
