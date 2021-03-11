package jobenqueuer_test

import (
	"testing"

	"github.com/hoangmirs/go-scraper/bootstrap"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJobenqueuer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JobEnqueuer Suite")
}

var _ = BeforeSuite(func() {
	bootstrap.SetUp()
})
