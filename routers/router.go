package routers

import (
	"github.com/hoangmirs/go-scraper/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.Keyword{})
	web.Router("/keyword", &controllers.Keyword{})
	web.Router("/register", &controllers.Registration{})
	web.Router("/login", &controllers.Session{})
	web.Router("/logout", &controllers.Session{}, "get:Delete")
}
