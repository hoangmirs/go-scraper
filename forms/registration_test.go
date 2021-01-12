package forms_test

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/hoangmirs/go-scraper/forms"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegistrationForm", func() {
	AfterEach(func() {
		o := orm.NewOrm()
		_, err := o.Raw("TRUNCATE TABLE \"user\"").Exec()
		if err != nil {
			// If table can't be truncated, rebuild all tables (CAUTION: Star and Message db are lost!)
			// This is only for absolute startup
			err := orm.RunSyncdb("default", true, true)
			if err != nil {
				logs.Error(err)
			}
		}
	})

	Describe("CreateUser", func() {
		Context("given valid attributes", func() {
			It("returns nil", func() {
				registrationForm := forms.RegistrationForm{
					Email:                "hoang@nimblehq.co",
					Password:             "123456",
					PasswordConfirmation: "123456",
				}
				err := registrationForm.CreateUser()
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
					err := registrationForm.CreateUser()
					Expect(err.Error()).To(Equal("Email Can not be empty"))
				})
			})

			Context("given an existing email", func() {
				It("returns nil", func() {
					registrationForm1 := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "123456",
						PasswordConfirmation: "123456",
					}
					_ = registrationForm1.CreateUser()

					registrationForm2 := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "123456",
						PasswordConfirmation: "123456",
					}

					err := registrationForm2.CreateUser()
					Expect(err.Error()).To(Equal("pq: duplicate key value violates unique constraint \"user_email_key\""))
				})
			})

			Context("given an invalid email", func() {
				It("returns the correct error message", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang-nimblehq.co",
						Password:             "123456",
						PasswordConfirmation: "123456",
					}
					err := registrationForm.CreateUser()
					Expect(err.Error()).To(Equal("Email Must be a valid email address"))
				})
			})

			Context("given a blank password", func() {
				It("returns the correct error message", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "",
						PasswordConfirmation: "",
					}
					err := registrationForm.CreateUser()
					Expect(err.Error()).To(Equal("Password Can not be empty"))
				})
			})

			Context("given an invalid password", func() {
				It("returns the correct error message", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "123",
						PasswordConfirmation: "123",
					}
					err := registrationForm.CreateUser()
					Expect(err.Error()).To(Equal("Password Minimum size is 6"))
				})
			})

			Context("given non-matched password and password confirmation", func() {
				It("returns the correct error message", func() {
					registrationForm := forms.RegistrationForm{
						Email:                "hoang@nimblehq.co",
						Password:             "123456",
						PasswordConfirmation: "111111",
					}
					err := registrationForm.CreateUser()
					Expect(err.Error()).To(Equal("Password confirmation does not match"))
				})
			})
		})
	})
})
