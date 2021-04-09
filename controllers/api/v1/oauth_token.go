package apiv1controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/hoangmirs/go-scraper/models"
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"
	"github.com/hoangmirs/go-scraper/services/oauth"
)

type OAuthToken struct {
	baseController
}

func (c *OAuthToken) Post() {
	writer := httptest.NewRecorder()
	server := oauth.GetOAuthServer()

	err := server.HandleTokenRequest(writer, c.Ctx.Request)
	if err != nil {
		c.renderGenericError(err)
	}
	body := writer.Body.Bytes()

	if writer.Code != http.StatusOK {
		oAuthError := &models.OAuthError{}
		err = json.Unmarshal(body, oAuthError)
		if err != nil {
			c.renderGenericError(err)
		}

		c.renderError("Unauthorized", oAuthError.Message, oAuthError.Code, http.StatusUnauthorized, nil)
	}

	oAuthToken := &models.OAuthToken{}
	err = json.Unmarshal(body, oAuthToken)

	if err != nil {
		c.renderGenericError(err)
	}

	oauthTokenSerializer := v1serializers.OAuthToken{
		OAuthToken: oAuthToken,
	}

	err = c.renderJSON(oauthTokenSerializer.Data())
	if err != nil {
		c.renderGenericError(err)
	}
}

func (c *OAuthToken) Revoke() {
	err := c.ensureAuthenticatedClient()
	if err != nil {
		c.renderError("Unauthorized client", err.Error(), "unauthorized_client", http.StatusUnauthorized, nil)
	}

	server := oauth.GetOAuthServer()
	tokenInfo, err := server.ValidationBearerToken(c.Ctx.Request)
	if err != nil {
		c.renderError("Invalid token", err.Error(), "invalid_token", http.StatusUnprocessableEntity, nil)
	}

	err = oauth.GetTokenStore().RemoveByAccess(context.TODO(), tokenInfo.GetAccess())
	if err != nil {
		c.renderGenericError(err)
	}

	c.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
}
