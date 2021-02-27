package scraper

import (
	"encoding/json"

	"github.com/hoangmirs/go-scraper/models"
)

func (service *ScraperService) saveToDB() error {
	nonAdwordLinks, err := json.Marshal(service.parsingResult.NonAdwordLinks)
	if err != nil {
		return err
	}

	adwordLinks, err := json.Marshal(service.parsingResult.AdwordLinks)
	if err != nil {
		return err
	}

	shopAdwordLinks, err := json.Marshal(service.parsingResult.ShopAdwordLinks)
	if err != nil {
		return err
	}

	nonAdwordLinksCount := len(service.parsingResult.NonAdwordLinks)
	adwordLinksCount := len(service.parsingResult.AdwordLinks)
	shopAdwordLinksCount := len(service.parsingResult.ShopAdwordLinks)
	totalCount := nonAdwordLinksCount + adwordLinksCount + shopAdwordLinksCount

	keywordResult := &models.KeywordResult{
		KeyWord:              service.Keyword,
		User:                 service.User,
		NonAdwordLinksCount:  nonAdwordLinksCount,
		NonAdwordLinks:       string(nonAdwordLinks),
		AdwordLinksCount:     adwordLinksCount,
		AdwordLinks:          string(adwordLinks),
		ShopAdwordLinksCount: shopAdwordLinksCount,
		ShopAdwordLinks:      string(shopAdwordLinks),
		LinksCount:           totalCount,
		HtmlCode:             service.parsingResult.HTMLCode,
	}

	_, err = models.CreateKeywordResult(keywordResult)
	if err != nil {
		return err
	}

	return err
}
