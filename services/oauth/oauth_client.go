package oauth

import (
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/google/uuid"
)

func GenerateClient() (*models.Client, error) {
	clientID := uuid.New().String()
	clientSecret := uuid.New().String()

	client := &models.Client{
		ID:     clientID,
		Secret: clientSecret,
	}

	err := GetClientStore().Create(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}
