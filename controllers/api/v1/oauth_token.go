package apiv1controllers

import (
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
