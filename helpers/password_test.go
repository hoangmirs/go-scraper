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
})
