package models_test

import (
	"testing"

	"github.com/hoangmirs/go-scraper/bootstrap"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var _ = BeforeSuite(func() {
	bootstrap.SetUp()
})
