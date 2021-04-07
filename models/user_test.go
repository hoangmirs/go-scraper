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
					Email:             faker.Email(),
					EncryptedPassword: faker.Password(),
				}

				userID, err := models.CreateUser(&user)
				if err != nil {
					Fail("Failed to create user: " + err.Error())
				}

				Expect(userID).To(BeNumerically(">", 0))
			})

			It("does NOT return error", func() {
				user := models.User{
					Email:             faker.Email(),
					EncryptedPassword: faker.Password(),
				}
				_, err := models.CreateUser(&user)

				Expect(err).To(BeNil())
			})
		})

		Context("given invalid attributes", func() {
			Context("given email that already exist in database", func() {
				It("returns an error", func() {
					email := faker.Email()
					_ = fabricators.FabricateUser(email, faker.Password())

					user := models.User{
						Email:             email,
						EncryptedPassword: faker.Password(),
					}
					userID, err := models.CreateUser(&user)

					Expect(err.Error()).To(Equal(`pq: duplicate key value violates unique constraint "user_email_key"`))
					Expect(userID).To(Equal(int64(0)))
				})
			})
		})
	})

	Describe("#GetUserByID", func() {
		Context("given a valid userID", func() {
			It("returns the correct user", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())

				storedUser, err := models.GetUserByID(user.Id)
				if err != nil {
					Fail("Failed to get user: " + err.Error())
				}

				Expect(storedUser.Id).To(Equal(user.Id))
			})

			It("does NOT return error", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())

				_, err := models.GetUserByID(user.Id)

				Expect(err).To(BeNil())
			})
		})

		Context("given invalid attributes", func() {
			Context("given an invalid userID", func() {
				It("returns an error", func() {
					user, err := models.GetUserByID(999)

					Expect(err.Error()).To(Equal(`<QuerySeter> no row found`))
					Expect(user).To(BeNil())
				})
			})
		})
	})
})
