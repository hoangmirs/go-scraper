package controllers

import (
	"github.com/hoangmirs/go-scraper/helpers"
	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// baseController implements global settings for all other controllers
type baseController struct {
	web.Controller

	CurrentUser *models.User
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

	c.Data[currentUserSessionKey] = c.GetCurrentUser()

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["FlashMessage"] = "shared/flash_message.html"

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}
}

func (c *baseController) SetCurrentUser(user *models.User) {
	err := c.SetSession(currentUserSessionKey, user.Id)
	if err != nil {
		logs.Error("Cannot set session:", err)
	}

	c.CurrentUser = user
}

func (c *baseController) GetCurrentUser() *models.User {
	if c.GetSession(currentUserSessionKey) == nil {
		return nil
	}

	userID := c.GetSession(currentUserSessionKey).(uint)
	user := &models.User{Base: models.Base{Id: userID}}
	o := orm.NewOrm()
	err := o.Read(user)

	if err != nil {
		return nil
	}

	c.CurrentUser = user
	return c.CurrentUser
}
