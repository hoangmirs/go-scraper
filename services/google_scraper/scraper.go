package google_scraper

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/hoangmirs/go-scraper/models"
)

type ParsingResult struct {
	Keyword         string
	HTMLCode        string
	LinksCount      int
	NonAdwordLinks  []string
	AdwordLinks     []string
	ShopAdwordLinks []string
	User            *models.User
}

var selectors = map[string]string{
	"nonAds":      "#search .yuRUbf > a",
	"ads":         "#tads .d5oMvf > a",
	"shopAds":     ".mnr-c.pla-unit .pla-unit-title a.pla-unit-title-link",
	"mobileLinks": ".ezO2md a.fuLhoc.ZWRArf",
	"mobileAds":   "span.dloBPe",
}

const urlPattern = "https://www.google.com/search?q=%s"

func Search(keyword string, user *models.User) {
	parsingResult := ParsingResult{
		Keyword: keyword,
		User:    user,
	}

	collector := colly.NewCollector()

	extensions.RandomUserAgent(collector)

	collector.OnRequest(func(r *colly.Request) {
		logs.Info("Visiting: %s user-agent: %v", r.URL, r.Headers.Get("User-Agent"))
	})

	collector.OnResponse(func(r *colly.Response) {
		logs.Info("HTML: %v", r.StatusCode)
		// parsingResult.HTMLCode = string(r.Body[:])
		// _ = ioutil.WriteFile("file.html", r.Body, 0644)
	})

	collector.OnHTML(selectors["nonAds"], func(e *colly.HTMLElement) {
		link := e.Attr("href")
		parsingResult.NonAdwordLinks = append(parsingResult.NonAdwordLinks, link)
	})

	collector.OnHTML(selectors["ads"], func(e *colly.HTMLElement) {
		link := e.Attr("href")
		parsingResult.AdwordLinks = append(parsingResult.AdwordLinks, link)
	})

	collector.OnHTML(selectors["shopAds"], func(e *colly.HTMLElement) {
		link := e.Attr("href")
		parsingResult.ShopAdwordLinks = append(parsingResult.ShopAdwordLinks, link)
	})

	collector.OnHTML(selectors["mobileLinks"], func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(e.DOM.Find(selectors["mobileAds"]).Nodes) > 0 {
			parsingResult.AdwordLinks = append(parsingResult.AdwordLinks, link)
		} else {
			parsingResult.NonAdwordLinks = append(parsingResult.NonAdwordLinks, link)
		}
	})

	collector.OnScraped(func(r *colly.Response) {
		err := saveToDB(parsingResult)
		logs.Error("Error when saving to DB:", err)
	})

	url := fmt.Sprintf(urlPattern, url.QueryEscape(keyword))
	_ = collector.Visit(url)
}

func saveToDB(parsingResult ParsingResult) error {
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
