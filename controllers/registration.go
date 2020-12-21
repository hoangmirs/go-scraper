package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/hoangmirs/go-scraper/helpers"
)

type Registration struct {
	web.Controller
}

func (c *Registration) Get() {
	web.ReadFromRequest(&c.Controller)

	c.Layout = "layouts/authentication.html"
	c.TplName = "registration/new.html"

	helpers.SetControllerAttributes(&c.Controller)
	c.Data["Title"] = "Create a new account"
}
