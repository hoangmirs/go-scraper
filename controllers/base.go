package controllers

import (
	"net/http"

	"github.com/hoangmirs/go-scraper/helpers"
	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// baseController implements global settings for all other controllers
type baseController struct {
	web.Controller

	CurrentUser              *models.User
	requireAuthenticatedUser bool
	requireGuestUser         bool
}

// NestPreparer : check the below docs
// https://beego.me/docs/mvc/controller/controller.md#custom-logic
type NestPreparer interface {
	NestPrepare()
}

const currentUserSessionKey = "CurrentUser"

// Prepare implements Prepare method for baseController
func (c *baseController) Prepare() {
	// Setting properties
	helpers.SetControllerAttributes(&c.Controller)

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "shared/header.html"
	c.LayoutSections["FlashMessage"] = "shared/flash_message.html"

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}

	if c.requireGuestUser && !c.ensureGuestUser() {
		c.Ctx.Redirect(http.StatusFound, "/")
	}

	if c.requireAuthenticatedUser && !c.ensureAuthenticatedUser() {
		c.Ctx.Redirect(http.StatusFound, "/login")
	}
}

func (c *baseController) SetCurrentUser(user *models.User) {
	err := c.SetSession(currentUserSessionKey, user.Id)
	if err != nil {
		logs.Error("Cannot set session:", err)
	}

	c.Data["CurrentUser"] = user
	c.CurrentUser = user
}

func (c *baseController) ensureAuthenticatedUser() bool {
	if c.GetSession(currentUserSessionKey) == nil {
		return false
	}

	userID := c.GetSession(currentUserSessionKey).(uint)

	user := &models.User{
		Base: models.Base{Id: userID},
	}

	// Check if user exists in DB
	o := orm.NewOrm()
	err := o.Read(user)
	if err != nil {
		return false
	}

	// Re-set current user with latest changes
	c.SetCurrentUser(user)

	return true
}

func (c *baseController) ensureGuestUser() bool {
	return c.GetSession(currentUserSessionKey) == nil
}
