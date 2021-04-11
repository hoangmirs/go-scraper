package forms_test

import (
	"github.com/hoangmirs/go-scraper/forms"
	"github.com/hoangmirs/go-scraper/tests/fabricators"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SessionForm", func() {
	AfterEach(func() {
		TruncateTables("user")
	})

	Describe("Authenticate", func() {
		Context("given valid attributes", func() {
			Context("given an existing user", func() {
				It("returns the user without error", func() {
					email := "hoang@nimblehq.co"
					password := "123456"
					_ = fabricators.FabricateUser(email, password)

					sessionForm := forms.SessionForm{
						Email:    email,
						Password: password,
					}

					user, err := sessionForm.Authenticate()
					Expect(user.Id).To(BeNumerically(">", 0))
					Expect(err).To(BeNil())
				})
			})

			Context("given a non-existing user", func() {
				It("returns the error", func() {
					sessionForm := forms.SessionForm{
						Email:    "hoang@nimblehq.co",
						Password: "123456",
					}

					user, err := sessionForm.Authenticate()
					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Incorrect username or password"))
				})
			})
		})

		Context("given invalid attributes", func() {
			Context("given an invalid email", func() {
				It("returns the error", func() {
					sessionForm := forms.SessionForm{
						Email:    "",
						Password: "123456",
					}

					user, err := sessionForm.Authenticate()
					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Email cannot be empty"))
				})

				It("returns the error", func() {
					sessionForm := forms.SessionForm{
						Email:    "invalid",
						Password: "123456",
					}

					user, err := sessionForm.Authenticate()
					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Email must be a valid email address"))
				})
			})

			Context("given an invalid password", func() {
				It("returns the error", func() {
					sessionForm := forms.SessionForm{
						Email:    "hoang@nimblehq.co",
						Password: "",
					}

					user, err := sessionForm.Authenticate()
					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Password cannot be empty"))
				})

				It("returns the error", func() {
					sessionForm := forms.SessionForm{
						Email:    "hoang@nimblehq.co",
						Password: "1",
					}

					user, err := sessionForm.Authenticate()
					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Password minimum size is 6"))
				})
			})

			Context("given wrong password", func() {
				It("returns the error", func() {
					email := "hoang@nimblehq.co"
					password := "123456"
					_ = fabricators.FabricateUser(email, password)

					sessionForm := forms.SessionForm{
						Email:    email,
						Password: "wrongpass",
					}

					user, err := sessionForm.Authenticate()
					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Incorrect username or password"))
				})
			})
		})
	})
})
