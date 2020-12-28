package bootstrap

import (
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/hoangmirs/go-scraper/models"  // Models
	_ "github.com/hoangmirs/go-scraper/routers" // Models
)

func init() {
	logs.Info("Setup DB")
	SetUpDB()
}
