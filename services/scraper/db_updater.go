package scraper

import (
	"encoding/json"

	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/core/logs"
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

	keyword := service.Keyword
	keyword.NonAdwordLinksCount = nonAdwordLinksCount
	keyword.NonAdwordLinks = string(nonAdwordLinks)
	keyword.AdwordLinksCount = adwordLinksCount
	keyword.AdwordLinks = string(adwordLinks)
	keyword.ShopAdwordLinksCount = shopAdwordLinksCount
	keyword.ShopAdwordLinks = string(shopAdwordLinks)
	keyword.LinksCount = totalCount
	keyword.HtmlCode = service.parsingResult.HTMLCode

	keyword.Status = models.Processed
	err = models.UpdateKeyword(keyword)
	if err != nil {
		return err
	}

	return err
}

func (service *ScraperService) updateKeywordStatus(status models.KeywordStatus) {
	service.Keyword.Status = status
	err := models.UpdateKeyword(service.Keyword)
	if err != nil {
		logs.Error("Failed to update status: %v", err)
	}
}
