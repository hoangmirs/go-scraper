package scraper

import (
	"encoding/json"

	"github.com/hoangmirs/go-scraper/models"
)

func SaveToDB(parsingResult ParsingResult) error {
	nonAdwordLinks, err := json.Marshal(parsingResult.NonAdwordLinks)
	if err != nil {
		return err
	}

	adwordLinks, err := json.Marshal(parsingResult.AdwordLinks)
	if err != nil {
		return err
	}

	shopAdwordLinks, err := json.Marshal(parsingResult.ShopAdwordLinks)
	if err != nil {
		return err
	}

	nonAdwordLinksCount := len(parsingResult.NonAdwordLinks)
	adwordLinksCount := len(parsingResult.AdwordLinks)
	shopAdwordLinksCount := len(parsingResult.ShopAdwordLinks)
	totalCount := nonAdwordLinksCount + adwordLinksCount + shopAdwordLinksCount

	keywordResult := &models.KeywordResult{
		KeyWord:              parsingResult.Keyword,
		NonAdwordLinksCount:  nonAdwordLinksCount,
		NonAdwordLinks:       string(nonAdwordLinks),
		AdwordLinksCount:     adwordLinksCount,
		AdwordLinks:          string(adwordLinks),
		ShopAdwordLinksCount: shopAdwordLinksCount,
		ShopAdwordLinks:      string(shopAdwordLinks),
		LinksCount:           totalCount,
		HtmlCode:             parsingResult.HTMLCode,
		User:                 parsingResult.User,
	}

	_, err = models.CreateKeywordResult(keywordResult)
	if err != nil {
		return err
	}

	return err
}
