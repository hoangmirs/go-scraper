package apiv1controllers

import (
	"fmt"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/hoangmirs/go-scraper/forms"
	"github.com/hoangmirs/go-scraper/services/oauth"
)

type OAuthToken struct {
	baseController
}

func (c *OAuthToken) Post() {
	server := oauth.GetOAuthServer()

	server.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)

	err := server.HandleTokenRequest(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusForbidden)
	}
}

func passwordAuthorizationHandler(email string, password string) (string, error) {
	sessionForm := forms.SessionForm{
		Email:    email,
		Password: password,
	}
	user, err := sessionForm.Authenticate()
	if err != nil {
		return "", errors.ErrInvalidClient
	}

	return fmt.Sprint(user.Id), nil
}
