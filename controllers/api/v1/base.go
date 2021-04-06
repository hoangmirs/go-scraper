package apiv1controllers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/hoangmirs/go-scraper/models"
	"github.com/hoangmirs/go-scraper/services/oauth"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/google/jsonapi"
)

type baseController struct {
	web.Controller

	requireAuthenticatedUser bool
	CurrentUser              *models.User
	TokenInfo                *oauth2.TokenInfo
}

type NestPreparer interface {
	NestPrepare()
}

func (c *baseController) Prepare() {
	c.Ctx.Output.Header("Content-Type", jsonapi.MediaType)

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

	tokenInfo, err := server.ValidationBearerToken(c.Ctx.Request)
	if err != nil {
		return false
	}
	c.TokenInfo = &tokenInfo

	userID, err := strconv.ParseUint(tokenInfo.GetUserID(), 10, 64)
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

func (c *baseController) ensureAuthenticatedClient() error {
	clientID, clientSecret, err := oauth.GetOAuthServer().ClientInfoHandler(c.Ctx.Request)
	if err != nil {
		return err
	}

	client, err := oauth.GetClientStore().GetByID(context.TODO(), clientID)
	if err != nil || client.GetSecret() != clientSecret {
		return errors.New("Client authentication failed")
	}

	return nil
}

func (c *baseController) renderJSON(data interface{}) error {
	err := jsonapi.MarshalPayload(c.Ctx.ResponseWriter, data)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func (c *baseController) renderListJSON(data interface{}, meta *jsonapi.Meta, links *jsonapi.Links) error {
	response, err := jsonapi.Marshal(data)
	if err != nil {
		return err
	}

	payload, ok := response.(*jsonapi.ManyPayload)
	if !ok {
		return err
	}

	payload.Meta = meta
	payload.Links = links

	return json.NewEncoder(c.Ctx.ResponseWriter).Encode(payload)
}

func (c *baseController) renderGenericError(gErr error) {
	c.renderError("Generic error", gErr.Error(), "generic_error", http.StatusBadRequest, nil)
}

func (c *baseController) renderError(title string, detail string, code string, status int, meta *map[string]interface{}) {
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

	c.StopRun()
}
