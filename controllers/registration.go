package controllers

import (
	"html/template"
	"net/http"

	"github.com/hoangmirs/go-scraper/forms"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type Registration struct {
	baseController
}

func (c *Registration) NestPrepare() {
	c.requireGuestUser = true
	c.Layout = "layouts/authentication.html"
}

func (c *Registration) Get() {
	web.ReadFromRequest(&c.Controller)

	c.renderNewRegistrationView()
}

func (c *Registration) Post() {
	registrationForm := forms.RegistrationForm{}
	flash := web.NewFlash()

	err := c.ParseForm(&registrationForm)
	if err != nil {
		logs.Error("Error when parsing data: %v", err)
	}

	_, formError := registrationForm.CreateUser()
	if formError != nil {
		flash.Error(formError.Error())
		flash.Store(&c.Controller)

		c.Data["Form"] = registrationForm

		c.renderNewRegistrationView()
	} else {
		flash.Success("Account created successfully")
		flash.Store(&c.Controller)

		c.Ctx.Redirect(http.StatusFound, "/login")
	}
}

func (c *Registration) renderNewRegistrationView() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "registration/new.html"

	c.Data["Title"] = "Create a new account"
}
