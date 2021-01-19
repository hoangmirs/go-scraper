package helpers_test

import (
	"github.com/hoangmirs/go-scraper/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Path", func() {
	Describe("RootDir", func() {
		Context("given the current file", func() {
			It("returns root directory of this project", func() {
				Expect(helpers.RootDir()).To(ContainSubstring("go-scraper"))
			})
		})
	})
})
