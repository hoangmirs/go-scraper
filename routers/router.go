package routers

import (
	"github.com/hoangmirs/go-scraper/controllers"
	beego "github.com/astaxie/beego/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
