package fabricators

import (
	"context"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/hoangmirs/go-scraper/models"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	oauthmodels "github.com/go-oauth2/oauth2/v4/models"
	"github.com/onsi/ginkgo"
)

func FabricateToken(user *models.User) *oauthmodels.Token {
	client := FabricateOAuthClient(faker.UUIDDigit(), faker.UUIDDigit())
	userID := fmt.Sprint(user.Id)
	data := &oauth2.GenerateBasic{
		Client:   client,
		UserID:   userID,
		CreateAt: time.Now(),
	}

	gen := generates.NewAccessGenerate()
	access, refresh, err := gen.Token(context.Background(), data, true)
	if err != nil {
		ginkgo.Fail("Generate token failed:" + err.Error())
	}

	token := oauthmodels.NewToken()
	token.SetClientID(client.ID)
	token.SetUserID(userID)
	token.SetAccess(access)
	token.SetRefresh(refresh)
	token.SetAccessCreateAt(time.Now())

	return token
}
