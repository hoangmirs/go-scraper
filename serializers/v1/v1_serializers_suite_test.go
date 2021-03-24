package v1serializers_test

import (
	"testing"

	"github.com/hoangmirs/go-scraper/bootstrap"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "V1 Serializers Suite")
}

var _ = BeforeSuite(func() {
	bootstrap.SetUp()
})
