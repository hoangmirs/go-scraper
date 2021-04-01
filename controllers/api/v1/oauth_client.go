package apiv1controllers

import (
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"
	oauthservice "github.com/hoangmirs/go-scraper/services/oauth"

	"github.com/beego/beego/v2/core/logs"
)

type OAuthClient struct {
	baseController
}

func (c *OAuthClient) Post() {
	oauthClient, err := oauthservice.GenerateClient()
	if err != nil {
		err = c.renderGenericError(err)
		if err != nil {
			logs.Error("Error: %v", err.Error())
		}
		return
	}

	oauthClientSerializer := v1serializers.OAuthClient{
		OAuthClient: oauthClient,
	}

	c.Data["json"] = oauthClientSerializer.Data()
	err = c.ServeJSON()
	if err != nil {
		logs.Error("Error: %v", err.Error())
	}
}
