package controllers

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	baseController
}

func (c *MainController) NestPrepare() {
	c.requireAuthenticatedUser = true
}

func (c *MainController) Get() {
	web.ReadFromRequest(&c.Controller)

	c.Layout = "layouts/application.html"
	c.TplName = "index.html"
}
