package forms_test

import (
	"github.com/hoangmirs/go-scraper/forms"
	"github.com/hoangmirs/go-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegistrationForm", func() {
	AfterEach(func() {
		tests.ClearUserTable()
	})

	Describe("CreateUser", func() {
		Context("given valid attributes", func() {
			It("returns nil", func() {
				registrationForm := forms.RegistrationForm{
					Email:                "hoang@nimblehq.co",
					Password:             "123456",
					PasswordConfirmation: "123456",
				}
				_, err := registrationForm.CreateUser()
				Expect(err).To(BeNil())
			})
		})

		Context("given invalid attributes", func() {
			Context("given a blank email", func() {
				It("returns the correct error message", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "",
						Password:             "123456",
						PasswordConfirmation: "123456",
					}
					_, err := registrationForm.CreateUser()
					Expect(err.Error()).To(Equal("Email can not be empty"))
				})
			})

			Context("given an existing email", func() {
				It("returns nil error", func() {
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

					_, err := registrationForm2.CreateUser()
					Expect(err.Error()).To(Equal("Email already exists"))
				})
			})

			Context("given an invalid email", func() {
				It("returns the correct error message", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang-nimblehq.co",
						Password:             "123456",
						PasswordConfirmation: "123456",
					}
					_, err := registrationForm.CreateUser()
					Expect(err.Error()).To(Equal("Email must be a valid email address"))
				})
			})

			Context("given a blank password", func() {
				It("returns the correct error message", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "",
						PasswordConfirmation: "",
					}
					_, err := registrationForm.CreateUser()
					Expect(err.Error()).To(Equal("Password can not be empty"))
				})
			})

			Context("given a short password", func() {
				It("returns the correct error message", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "123",
						PasswordConfirmation: "123",
					}
					_, err := registrationForm.CreateUser()
					Expect(err.Error()).To(Equal("Password minimum size is 6"))
				})
			})

			Context("given non-matched password and password confirmation", func() {
				It("returns the correct error message", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "123456",
						PasswordConfirmation: "111111",
					}
					_, err := registrationForm.CreateUser()
					Expect(err.Error()).To(Equal("Password confirmation does not match"))
				})
			})
		})
	})
})
