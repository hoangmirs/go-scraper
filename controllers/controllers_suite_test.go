package controllers_test

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/hoangmirs/go-scraper/bootstrap"
	"github.com/hoangmirs/go-scraper/helpers"
	_ "github.com/hoangmirs/go-scraper/routers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	appPath := fmt.Sprintf("%s", helpers.RootDir())
	web.TestBeegoInit(appPath)
	bootstrap.SetUpDB()
})
