package keyword

import (
	"errors"

	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/core/logs"
)

type KeywordService struct {
	Keywords []string
	User     *models.User
}

func (service *KeywordService) Run() error {
	err := service.validate()
	if err != nil {
		return err
	}

	err = service.saveAndEnqueue()
	if err != nil {
		return err
	}

	return nil
}

func (service *KeywordService) validate() error {
	if len(service.Keywords) == 0 {
		return errors.New("Keywords are empty")
	}

	if service.User == nil {
		return errors.New("User object required")
	}

	return nil
}

func (service *KeywordService) saveAndEnqueue() error {
	for _, value := range service.Keywords {
		keyword := models.Keyword{Keyword: value, User: service.User}

		_, err := models.CreateKeyword(&keyword)

		if err != nil {
			logs.Error("Error when creating keyword: %v", err.Error())
			return err
		}

		err = EnqueueKeyword(keyword)
		if err != nil {
			logs.Error("Error when enqueuing keyword: %v", err.Error())
			return err
		}
	}

	return nil
}
