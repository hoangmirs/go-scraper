package bootstrap

import (
	"fmt"

	"github.com/hoangmirs/go-scraper/database"
	"github.com/hoangmirs/go-scraper/helpers"
	_ "github.com/hoangmirs/go-scraper/routers" // Routers
	"github.com/hoangmirs/go-scraper/services/jobenqueuer"
	"github.com/hoangmirs/go-scraper/services/oauth"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	configPath := fmt.Sprintf("%s/conf/app.conf", helpers.RootDir())
	err := web.LoadAppConfig("ini", configPath)
	if err != nil {
		logs.Error("Error when parsing config: ", err)
	}
}

func SetUp() {
	database.SetUpDB()
	database.SetupRedisPool()
	jobenqueuer.SetUpEnqueuer()
	SetUpTemplateFunction()
	err := oauth.SetUpOAuthServer()
	if err != nil {
		logs.Error("Error when setting up OAuth server: %v", err.Error())
	}
}
