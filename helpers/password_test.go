package helpers_test

import (
	"github.com/hoangmirs/go-scraper/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Password", func() {
	Describe("HashPassword", func() {
		Context("given a string", func() {
			It("returns the hashed string", func() {
				Expect(helpers.HashPassword("hello password")).To(ContainSubstring("$"))
			})
		})
	})

	Describe("ComparePassword", func() {
		Context("given a correct password", func() {
			It("returns nil", func() {
				hashedPassword := "$2a$10$9rnyG9shT0T49i9CBQcnYuICoTwgKvwFCY/EEET63PqjJTat1qHRW"
				correctPassword := "123456"
				Expect(helpers.ComparePassword(hashedPassword, correctPassword)).To(BeNil())
			})
		})

		Context("given an incorrect password", func() {
			It("returns error", func() {
				hashedPassword := "$2a$10$9rnyG9shT0T49i9CBQcnYuICoTwgKvwFCY/EEET63PqjJTat1qHRW"
				incorrectPassword := "111111"
				Expect(helpers.ComparePassword(hashedPassword, incorrectPassword)).NotTo(BeNil())
			})
		})
	})
})
