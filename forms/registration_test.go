package forms_test

import (
	"testing"

	"github.com/beego/beego/v2/core/logs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRegistrationForm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Registration Form Suite")
}

var _ = Describe("RegistrationForm", func() {
	BeforeEach(func() {
		logs.Info("before each")
	})

	AfterEach(func() {
		logs.Info("after each")
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				logs.Info("test 1")
				Expect(1).To(Equal(1))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				logs.Info("test 2")
				Expect(1).To(Equal(1))
			})
		})
	})
})
