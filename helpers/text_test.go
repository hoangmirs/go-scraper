package helpers_test

import (
	"github.com/hoangmirs/go-scraper/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Text", func() {
	Describe("ToSnakeCase", func() {
		Context("given a blank string", func() {
			It("returns blank string", func() {
				Expect(helpers.ToSnakeCase("")).To(BeEmpty())
			})
		})

		Context("given a normal string", func() {
			It("returns that string", func() {
				Expect(helpers.ToSnakeCase("normal string")).To(Equal("normal string"))
			})
		})

		Context("given a camel case string", func() {
			It("returns converted string", func() {
				Expect(helpers.ToSnakeCase("camelCase")).To(Equal("camel_case"))
			})
		})
	})
})
