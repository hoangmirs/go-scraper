package controllers

import "github.com/beego/beego/v2/server/web"

type RegistrationController struct {
	web.Controller
}

func (c *RegistrationController) Get() {
	web.ReadFromRequest(&c.Controller)

	c.Layout = "layouts/authentication.html"
	c.TplName = "registration/new.html"

	c.Data["Title"] = "Create a new account"
}
