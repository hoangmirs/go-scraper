package middlewares

import (
	"github.com/hoangmirs/go-scraper/conf"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/auth"
)

func BasicAuthenticationMiddleware() web.FilterFunc {
	return auth.Basic(conf.GetString("basicAuthenticationUsername"), conf.GetString("basicAuthenticationPassword"))
}
