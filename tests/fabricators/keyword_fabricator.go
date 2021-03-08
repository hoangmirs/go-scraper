package fabricators

import (
	"github.com/hoangmirs/go-scraper/models"

	"github.com/onsi/ginkgo"
)

func FabricateKeyword(keyword string, user *models.User) *models.Keyword {
	keywordObject := &models.Keyword{
		Keyword: keyword,
		User:    user,
	}

	_, err := models.CreateKeyword(keywordObject)
	if err != nil {
		ginkgo.Fail("Fabricate user error:" + err.Error())
	}

	return keywordObject
}
