package controllers

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	baseController
}

func (c *MainController) NestPrepare() {
	if c.CurrentUser == nil {
		c.Ctx.Redirect(302, "/login")
		return
	}
}

func (c *MainController) Get() {
	web.ReadFromRequest(&c.Controller)

	c.Layout = "layouts/application.html"
	c.TplName = "index.html"
}
