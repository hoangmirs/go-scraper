package forms_test

import (
	"github.com/hoangmirs/go-scraper/forms"
	. "github.com/hoangmirs/go-scraper/tests/test_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegistrationForm", func() {
	AfterEach(func() {
		TruncateTables("user")
	})

	Describe("CreateUser", func() {
		Context("given valid attributes", func() {
			It("creates a new user and returns nil error", func() {
				registrationForm := forms.RegistrationForm{
					Email:                "hoang@nimblehq.co",
					Password:             "123456",
					PasswordConfirmation: "123456",
				}
				user, err := registrationForm.CreateUser()

				Expect(user.Id).NotTo(BeNil())
				Expect(err).To(BeNil())
			})
		})

		Context("given invalid attributes", func() {
			Context("given a blank email", func() {
				It("returns the correct error message and does NOT create new user", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "",
						Password:             "123456",
						PasswordConfirmation: "123456",
					}
					user, err := registrationForm.CreateUser()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Email cannot be empty"))
				})
			})

			Context("given an existing email", func() {
				It("returns the correct error message and does NOT create new user", func() {
					// TODO : Using fabricator
					registrationForm1 := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "123456",
						PasswordConfirmation: "123456",
					}
					_, _ = registrationForm1.CreateUser()

					registrationForm2 := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "123456",
						PasswordConfirmation: "123456",
					}

					user, err := registrationForm2.CreateUser()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Email already exists"))
				})
			})

			Context("given an invalid email", func() {
				It("returns the correct error message and does NOT create new user", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang-nimblehq.co",
						Password:             "123456",
						PasswordConfirmation: "123456",
					}
					user, err := registrationForm.CreateUser()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Email must be a valid email address"))
				})
			})

			Context("given a blank password", func() {
				It("returns the correct error message and does NOT create new user", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "",
						PasswordConfirmation: "",
					}
					user, err := registrationForm.CreateUser()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Password cannot be empty"))
				})
			})

			Context("given a short password", func() {
				It("returns the correct error message and does NOT create new user", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "123",
						PasswordConfirmation: "123",
					}
					user, err := registrationForm.CreateUser()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Password minimum size is 6"))
				})
			})

			Context("given password and password confirmation do NOT match", func() {
				It("returns the correct error message and does NOT create new user", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "123456",
						PasswordConfirmation: "111111",
					}
					user, err := registrationForm.CreateUser()

					Expect(user).To(BeNil())
					Expect(err.Error()).To(Equal("Password does not match password confirmation"))
				})
			})
		})
	})
})
