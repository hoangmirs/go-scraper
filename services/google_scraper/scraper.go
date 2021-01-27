package google_scraper

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type ParsingResult struct {
	HTMLCode        string
	LinksCount      int
	NonAdwordLinks  []string
	AdwordLinks     []string
	ShopAdwordLinks []string
}

var selectors = map[string]string{
	"nonAds":      "#search .yuRUbf > a",
	"ads":         "#tads .d5oMvf > a",
	"shopAds":     ".mnr-c.pla-unit .pla-unit-title a.pla-unit-title-link",
	"mobileLinks": ".ezO2md a.fuLhoc.ZWRArf",
}

const urlPattern = "https://www.google.com/search?q=%s"

func Search(keyword string) {
	parsingResult := ParsingResult{}

	collector := colly.NewCollector()

	extensions.RandomUserAgent(collector)

	collector.OnRequest(func(r *colly.Request) {
		logs.Info("Visiting: %s user-agent:", r.URL, r.Headers.Get("User-Agent"))
	})

	collector.OnResponse(func(r *colly.Response) {
		logs.Info("HTML: %v", r.StatusCode)
		parsingResult.HTMLCode = string(r.Body[:])
		_ = ioutil.WriteFile("file.html", r.Body, 0644)
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
		if len(e.DOM.Find("span.dloBPe").Nodes) > 0 {
			parsingResult.AdwordLinks = append(parsingResult.AdwordLinks, link)
		} else {
			parsingResult.NonAdwordLinks = append(parsingResult.NonAdwordLinks, link)
		}
	})

	collector.OnScraped(func(r *colly.Response) {
		logs.Info("Non adword count: %v Adword count: %v Shop Adword count: %v ", len(parsingResult.NonAdwordLinks), len(parsingResult.AdwordLinks), len(parsingResult.ShopAdwordLinks))
	})

	url := fmt.Sprintf(urlPattern, url.QueryEscape(keyword))
	_ = collector.Visit(url)
}
