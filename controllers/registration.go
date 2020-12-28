package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/hoangmirs/go-scraper/forms"
)

type Registration struct {
	baseController
}

func (c *Registration) Get() {
	web.ReadFromRequest(&c.Controller)

	c.Layout = "layouts/authentication.html"
	c.TplName = "registration/new.html"

	c.Data["Title"] = "Create a new account"
}

func (c *Registration) Post() {
	userForm := forms.RegistrationForm{}
	flash := web.NewFlash()

	err := c.ParseForm(&userForm)
	if err != nil {
		logs.Error("Can not parse the form", err)
	}

	err = userForm.CreateUser()
	if err != nil {
		flash.Error(fmt.Sprint(err))
		flash.Store(&c.Controller)

		c.Layout = "layouts/authentication.html"
		c.TplName = "registration/new.html"

		c.Data["Title"] = "Create a new account"
	} else {
		flash.Success("Account created successfully")
		flash.Store(&c.Controller)

		c.Ctx.Redirect(http.StatusFound, "/register")
	}
}
