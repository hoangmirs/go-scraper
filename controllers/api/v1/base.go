package apiv1controllers

import (
	"net/http"
	"strconv"

	"github.com/hoangmirs/go-scraper/models"
	"github.com/hoangmirs/go-scraper/services/oauth"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/google/jsonapi"
)

type baseController struct {
	web.Controller

	requireAuthenticatedUser bool
	CurrentUser              *models.User
}

type NestPreparer interface {
	NestPrepare()
}

func (c *baseController) Prepare() {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}

	if c.requireAuthenticatedUser && !c.ensureAuthenticatedUser() {
		c.renderError("Unauthorized", "Authentication is required", "unauthorized", http.StatusUnauthorized, nil)
	}
}

func (c *baseController) ensureAuthenticatedUser() bool {
	server := oauth.GetOAuthServer()

	oauth, err := server.ValidationBearerToken(c.Ctx.Request)
	if err != nil {
		return false
	}

	userID, err := strconv.ParseUint(oauth.GetUserID(), 10, 64)
	if err != nil {
		logs.Error("Error when getting UserID: %v", err.Error())
		return false
	}

	user, err := models.GetUserByID(uint(userID))
	if err != nil {
		logs.Error("Error when querying user: %v", err.Error())
		return false
	}

	c.CurrentUser = user

	return true
}

func (c *baseController) renderJSON(data interface{}) error {
	err := jsonapi.MarshalPayload(c.Ctx.ResponseWriter, data)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func (c *baseController) renderGenericError(gErr error) {
	c.renderError("Generic error", gErr.Error(), "generic_error", http.StatusBadRequest, nil)
}

func (c *baseController) renderError(title string, detail string, code string, status int, meta *map[string]interface{}) {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.ResponseWriter.WriteHeader(status)

	err := jsonapi.MarshalErrors(c.Ctx.ResponseWriter, []*jsonapi.ErrorObject{{
		Title:  title,
		Detail: detail,
		Code:   code,
		Meta:   meta,
	}})
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		if err != nil {
			logs.Error("Error when rendering error message: %v", err.Error())
		}
	}
}
