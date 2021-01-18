package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/hoangmirs/go-scraper/forms"
)

type Session struct {
	baseController
}

func (c *Session) Get() {
	web.ReadFromRequest(&c.Controller)

	if c.CurrentUser != nil {
		c.Ctx.Redirect(http.StatusFound, "/")
	}

	c.renderCreateView()
}

func (c *Session) Post() {
	sessionForm := forms.SessionForm{}
	flash := web.NewFlash()

	err := c.ParseForm(&sessionForm)

	if err != nil {
		logs.Error("Can not parse the form", err)
	}

	user, err := sessionForm.Authenticate()
	if err != nil {
		flash.Error(fmt.Sprint(err))
		flash.Store(&c.Controller)

		c.renderCreateView()
	} else {
		flash.Success("Logging in successfully")
		flash.Store(&c.Controller)
		c.SetCurrentUser(user)

		c.Ctx.Redirect(http.StatusFound, "/")
	}
}

func (c *Session) renderCreateView() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/authentication.html"
	c.TplName = "session/new.html"

	c.Data["Title"] = "Log in to Go Scraper"
}
