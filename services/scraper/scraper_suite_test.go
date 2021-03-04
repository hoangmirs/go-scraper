package scraper_test

import (
	"testing"

	"github.com/hoangmirs/go-scraper/bootstrap"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestScraper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services/Scraper Suite")
}

var _ = BeforeSuite(func() {
	bootstrap.SetUp()
})
