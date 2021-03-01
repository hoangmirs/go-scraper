package scraper

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type ScraperService struct {
	Keyword   *models.Keyword
	Collector *colly.Collector // is used for testing

	parsingResult *ParsingResult
}

type ParsingResult struct {
	HTMLCode        string
	NonAdwordLinks  []string
	AdwordLinks     []string
	ShopAdwordLinks []string
}

var selectors = map[string][]string{
	"nonAds":      {"#search .yuRUbf > a", ".ZINbbc.xpd .kCrYT > a"},
	"ads":         {"#tads .d5oMvf > a", ".uEierd .ZINbbc a.C8nzq.BmP5tf"},
	"shopAds":     {".mnr-c.pla-unit .pla-unit-title a.pla-unit-title-link", ".qvfQJe .M09uTc.VoEfsd a"},
	"mobileLinks": {".ezO2md a.fuLhoc.ZWRArf"},
	"mobileAds":   {"span.dloBPe"},
}

const urlPattern = "https://www.google.com/search?q=%s"

func (service *ScraperService) Run() error {
	err := service.validateAttributes()
	if err != nil {
		return err
	}

	parsingResult := ParsingResult{}
	if service.Collector == nil {
		service.Collector = colly.NewCollector()
	}

	extensions.RandomUserAgent(service.Collector)

	service.Collector.OnRequest(func(r *colly.Request) {
		logs.Info("Visiting: %s user-agent: %v", r.URL, r.Headers.Get("User-Agent"))
	})

	service.Collector.OnResponse(func(r *colly.Response) {
		parsingResult.HTMLCode = string(r.Body[:])
	})

	for _, pattern := range selectors["nonAds"] {
		service.Collector.OnHTML(pattern, func(e *colly.HTMLElement) {
			link := checkLink(e.Attr("href"))
			parsingResult.NonAdwordLinks = append(parsingResult.NonAdwordLinks, link)
		})
	}

	for _, pattern := range selectors["ads"] {
		service.Collector.OnHTML(pattern, func(e *colly.HTMLElement) {
			link := checkLink(e.Attr("href"))
			parsingResult.AdwordLinks = append(parsingResult.AdwordLinks, link)
		})
	}

	for _, pattern := range selectors["shopAds"] {
		service.Collector.OnHTML(pattern, func(e *colly.HTMLElement) {
			link := checkLink(e.Attr("href"))
			parsingResult.ShopAdwordLinks = append(parsingResult.ShopAdwordLinks, link)
		})
	}

	// Need to check child nodes to detect ads links
	service.Collector.OnHTML(selectors["mobileLinks"][0], func(e *colly.HTMLElement) {
		link := checkLink(e.Attr("href"))
		if len(e.DOM.Find(selectors["mobileAds"][0]).Nodes) > 0 {
			parsingResult.AdwordLinks = append(parsingResult.AdwordLinks, link)
		} else {
			parsingResult.NonAdwordLinks = append(parsingResult.NonAdwordLinks, link)
		}
	})

	service.Collector.OnScraped(func(r *colly.Response) {
		service.parsingResult = &parsingResult
		err := service.saveToDB()
		if err != nil {
			logs.Error("Error when saving to DB:", err)
		}
	})

	url := fmt.Sprintf(urlPattern, url.QueryEscape(service.Keyword.Keyword))
	return service.Collector.Visit(url)
}

func (service *ScraperService) GetParsingResult() ParsingResult {
	return *service.parsingResult
}

// Private
func (service *ScraperService) validateAttributes() error {
	if len(service.Keyword.Keyword) == 0 {
		return errors.New("Keyword required")
	}

	if service.Keyword.User == nil {
		return errors.New("User required")
	}

	return nil
}

func checkLink(link string) string {
	if strings.HasPrefix(link, "http") {
		return link
	} else {
		return "https://www.google.com" + link
	}
}
