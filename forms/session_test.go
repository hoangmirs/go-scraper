package forms_test

import (
	"github.com/hoangmirs/go-scraper/forms"
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
					registrationForm := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "123456",
						PasswordConfirmation: "123456",
					}
					_, _ = registrationForm.CreateUser()

					sessionForm := forms.SessionForm{
						Email:    "hoang@nimblehq.co",
						Password: "123456",
					}

					user, err := sessionForm.Authenticate()
					Expect(user).NotTo(BeNil())
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
					Expect(err).NotTo(BeNil())
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
					Expect(err.Error()).To(Equal("Email can not be empty"))
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
					Expect(err.Error()).To(Equal("Password can not be empty"))
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
		})
	})
})
