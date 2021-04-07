package apiv1controllers

import (
	"context"
	"errors"
	"net/http"

	"github.com/hoangmirs/go-scraper/services/oauth"
)

type OAuthToken struct {
	baseController
}

func (c *OAuthToken) Post() {
	server := oauth.GetOAuthServer()

	err := server.HandleTokenRequest(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func (c *OAuthToken) Revoke() {
	err := c.ensureAuthenticatedClient()
	if err != nil {
		c.renderError("Unauthorized client", err.Error(), "unauthorized_client", http.StatusUnauthorized, nil)
	}

	token := c.GetString("token")
	if token == "" {
		err = errors.New("Token is blank")
		c.renderError(err.Error(), err.Error(), "token_required", http.StatusUnprocessableEntity, nil)
	}

	err = oauth.GetTokenStore().RemoveByAccess(context.TODO(), token)
	if err != nil {
		c.renderGenericError(err)
	}

	c.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
}
