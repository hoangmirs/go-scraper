package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/hoangmirs/go-scraper/forms"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type Session struct {
	baseController
}

func (c *Session) NestPrepare() {
	c.requireGuestUser = true
}

func (c *Session) Get() {
	web.ReadFromRequest(&c.Controller)

	c.renderNewSessionView()
}

func (c *Session) Post() {
	sessionForm := forms.SessionForm{}
	flash := web.NewFlash()

	err := c.ParseForm(&sessionForm)
	if err != nil {
		logs.Error("Error when parsing data: %v", err)
	}

	user, err := sessionForm.Authenticate()
	if err != nil {
		flash.Error(fmt.Sprint(err))
		flash.Store(&c.Controller)

		c.Data["Form"] = sessionForm

		c.renderNewSessionView()
	} else {
		flash.Success("Logging in successfully")
		flash.Store(&c.Controller)
		c.SetCurrentUser(user)

		c.Ctx.Redirect(http.StatusFound, "/")
	}
}

func (c *Session) Delete() {
	c.SetCurrentUser(nil)

	flash := web.NewFlash()
	flash.Success("Logging out successfully")
	flash.Store(&c.Controller)

	c.Ctx.Redirect(http.StatusFound, "/")
}

func (c *Session) renderNewSessionView() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/authentication.html"
	c.TplName = "session/new.html"

	c.Data["Title"] = "Log in to Go Scraper"
}
