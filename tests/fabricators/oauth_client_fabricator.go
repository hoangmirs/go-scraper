package fabricators

import (
	"github.com/hoangmirs/go-scraper/services/oauth"

	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/onsi/ginkgo"
)

func FabricateOAuthClient(id string, secret string) *models.Client {
	client := &models.Client{
		ID:     id,
		Secret: secret,
	}

	err := oauth.GetClientStore().Create(client)
	if err != nil {
		ginkgo.Fail("Fabricate OAuth client error:" + err.Error())
	}

	return client
}
