package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/hoangmirs/go-scraper/helpers"
)

type NestPreparer interface {
	NestPrepare()
}

// baseController implements global settings for all other controllers
type baseController struct {
	web.Controller
}

// Prepare implements Prepare method for baseController
func (c *baseController) Prepare() {
	// Setting properties
	helpers.SetControllerAttributes(&c.Controller)

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["FlashMessage"] = "shared/flash_message.html"

	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
