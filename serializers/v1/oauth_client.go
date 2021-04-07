package v1serializers

import "github.com/go-oauth2/oauth2/v4/models"

type OAuthClient struct {
	OAuthClient *models.Client
}

type oAuthClientResponse struct {
	ID           string `jsonapi:"primary,oauth_client"`
	ClientID     string `jsonapi:"attr,client_id"`
	ClientSecret string `jsonapi:"attr,client_secret"`
}

func (serializer *OAuthClient) Data() *oAuthClientResponse {
	data := &oAuthClientResponse{
		ID:           serializer.OAuthClient.ID,
		ClientID:     serializer.OAuthClient.ID,
		ClientSecret: serializer.OAuthClient.Secret,
	}

	return data
}
