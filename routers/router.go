package routers

import (
	"github.com/beego/beego/v2/server/web"

	"github.com/hoangmirs/go-scraper/controllers"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Router("/register", &controllers.Registration{})
}
