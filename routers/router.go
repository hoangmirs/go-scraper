package routers

import (
	"github.com/hoangmirs/go-scraper/controllers"
	apiv1controllers "github.com/hoangmirs/go-scraper/controllers/api/v1"
	"github.com/hoangmirs/go-scraper/middlewares"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.Keyword{})
	web.Router("/keyword", &controllers.Keyword{})
	web.Router("/keyword/:id", &controllers.Keyword{}, "get:Show")
	web.Router("/keyword/:id/html", &controllers.Keyword{}, "get:ShowHTML")
	web.Router("/register", &controllers.Registration{})
	web.Router("/login", &controllers.Session{})
	web.Router("/logout", &controllers.Session{}, "get:Delete")

	// API V1
	ns := web.NewNamespace("/api/v1",
		web.NSRouter("/health_check", &apiv1controllers.HealthCheck{}),
		web.NSRouter("/users", &apiv1controllers.User{}),

		web.NSNamespace("/oauth",
			web.NSRouter("/client", &apiv1controllers.OAuthClient{}),
			web.NSRouter("/token", &apiv1controllers.OAuthToken{}),
			web.NSRouter("/revoke", &apiv1controllers.OAuthToken{}, "post:Revoke"),
		),

		web.NSNamespace("/keywords",
			web.NSRouter("/", &apiv1controllers.Keyword{}),
			web.NSRouter("/:id", &apiv1controllers.Keyword{}, "get:Show"),
		),
	)

	web.AddNamespace(ns)

	// Filters
	web.InsertFilter("/api/v1/oauth/client", web.BeforeRouter, middlewares.BasicAuthenticationMiddleware())
}
