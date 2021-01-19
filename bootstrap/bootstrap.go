package bootstrap

import (
	_ "github.com/hoangmirs/go-scraper/models"  // Models
	_ "github.com/hoangmirs/go-scraper/routers" // Routers

	"github.com/beego/beego/v2/core/logs"
)

func init() {
	logs.Info("Setup DB")
	SetUpDB()
}
