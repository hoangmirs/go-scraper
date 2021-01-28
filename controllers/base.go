package controllers

import (
	"github.com/hoangmirs/go-scraper/helpers"

	"github.com/beego/beego/v2/server/web"
)

// NestPreparer : check the below docs
// https://beego.me/docs/mvc/controller/controller.md#custom-logic
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

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}
}
