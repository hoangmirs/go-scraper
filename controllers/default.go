package controllers

import (
	beego "github.com/astaxie/beego/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "mirs.xyz"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
