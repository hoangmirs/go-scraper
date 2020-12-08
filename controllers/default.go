package controllers

import (
	beego "github.com/astaxie/beego/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "layouts/application.html"
	c.TplName = "index.html"
}
