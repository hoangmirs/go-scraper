package apiv1controllers_test

import (
	"testing"

	"github.com/hoangmirs/go-scraper/bootstrap"
	"github.com/hoangmirs/go-scraper/helpers"
	_ "github.com/hoangmirs/go-scraper/routers"

	"github.com/beego/beego/v2/server/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API V1 Controllers Suite")
}

var _ = BeforeSuite(func() {
	bootstrap.SetUp()
	web.TestBeegoInit(helpers.RootDir())
})
