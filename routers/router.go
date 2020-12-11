package routers

import (
	beego "github.com/astaxie/beego/server/web"
	"github.com/hoangmirs/go-scraper/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
