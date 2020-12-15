package routers

import (
	"github.com/astaxie/beego"
	"github.com/hoangmirs/go-scraper/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
