package v1serializers

import "github.com/hoangmirs/go-scraper/models"

type OAuthToken struct {
	OAuthToken *models.OAuthToken
}

type oAuthTokenResponse struct {
	ID           string `jsonapi:"primary,oauth_token"`
	AccessToken  string `jsonapi:"attr,access_token"`
	RefreshToken string `jsonapi:"attr,refresh_token"`
	TokenType    string `jsonapi:"attr,token_type"`
	ExpiresIn    uint64 `jsonapi:"attr,expires_in"`
}

func (serializer *OAuthToken) Data() *oAuthTokenResponse {
	data := &oAuthTokenResponse{
		ID:           serializer.OAuthToken.AccessToken,
		AccessToken:  serializer.OAuthToken.AccessToken,
		RefreshToken: serializer.OAuthToken.RefreshToken,
		TokenType:    serializer.OAuthToken.TokenType,
		ExpiresIn:    serializer.OAuthToken.ExpiresIn,
	}

	return data
}
