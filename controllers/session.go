package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/hoangmirs/go-scraper/forms"

	"github.com/beego/beego/v2/server/web"
)

type Session struct {
	baseController
}

func (c *Session) Get() {
	web.ReadFromRequest(&c.Controller)

	if c.CurrentUser != nil {
		c.Ctx.Redirect(http.StatusFound, "/")
	}

	c.renderNewSessionView()
}

func (c *Session) Post() {
	sessionForm := forms.SessionForm{}
	flash := web.NewFlash()

	_ = c.ParseForm(&sessionForm)

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

func (c *Session) renderNewSessionView() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/authentication.html"
	c.TplName = "session/new.html"

	c.Data["Title"] = "Log in to Go Scraper"
}
