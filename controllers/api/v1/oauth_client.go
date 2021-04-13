package apiv1controllers

import (
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"
	oauthservice "github.com/hoangmirs/go-scraper/services/oauth"
)

type OAuthClient struct {
	baseController
}

func (c *OAuthClient) Post() {
	oauthClient, err := oauthservice.GenerateClient()
	if err != nil {
		c.renderGenericError(err)
	}

	oauthClientSerializer := v1serializers.OAuthClient{
		OAuthClient: oauthClient,
	}

	err = c.renderJSON(oauthClientSerializer.Data())
	if err != nil {
		c.renderGenericError(err)
	}
}
