package fabricators

import (
	"github.com/hoangmirs/go-scraper/models"
)

func FabricateKeyword(keyword string, user *models.User) (*models.Keyword, error) {
	keywordObject := models.Keyword{
		Keyword: keyword,
		User:    user,
	}

	_, err := models.CreateKeyword(&keywordObject)
	if err != nil {
		return nil, err
	}

	return &keywordObject, nil
}
