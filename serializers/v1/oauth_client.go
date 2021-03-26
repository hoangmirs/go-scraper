package v1serializers

import "github.com/go-oauth2/oauth2/v4/models"

type OAuthClient struct {
	OAuthClient *models.Client
}

type oAuthClientResponse struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (serializer *OAuthClient) Data() *oAuthClientResponse {
	data := &oAuthClientResponse{
		ClientID:     serializer.OAuthClient.ID,
		ClientSecret: serializer.OAuthClient.Secret,
	}

	return data
}
