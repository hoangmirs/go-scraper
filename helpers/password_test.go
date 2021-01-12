package helpers_test

import (
	"github.com/hoangmirs/go-scraper/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Password", func() {
	Describe("EncryptPassword", func() {
		Context("given a string", func() {
			It("returns the encrypted string", func() {
				Expect(helpers.EncryptPassword([]byte("hello password"))).To(ContainSubstring("$"))
			})
		})
	})
})
