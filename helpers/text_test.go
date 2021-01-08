package helpers

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestText(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Text helper Suite")
}

var _ = Describe("Text", func() {
	Describe("ToSnakeCase", func() {
		Context("given a blank string", func() {
			It("returns blank string", func() {
				Expect(ToSnakeCase("")).To(BeEmpty())
			})
		})

		Context("given a normal string", func() {
			It("returns that string", func() {
				Expect(ToSnakeCase("normal string")).To(Equal("normal string"))
			})
		})

		Context("given a camel case string", func() {
			It("returns converted string", func() {
				Expect(ToSnakeCase("camelCase")).To(Equal("camel_case"))
			})
		})
	})
})
