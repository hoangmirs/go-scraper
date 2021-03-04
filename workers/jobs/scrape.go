package jobs

import (
	"github.com/hoangmirs/go-scraper/models"
	"github.com/hoangmirs/go-scraper/services/scraper"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocraft/work"
)

type Context struct{}

const MaxFails = 3

func (c *Context) ScraperLog(job *work.Job, next work.NextMiddlewareFunc) error {
	logs.Info("Starting %v job for keywordID: %v", job.Name, job.ArgString("keywordID"))

	return next()
}

func (c *Context) PerformScrape(job *work.Job) error {
	keywordID := job.ArgInt64("keywordID")
	err := job.ArgError()
	if err != nil {
		return err
	}

	keyword, err := models.GetKeywordByID(keywordID)
	if err != nil {
		return err
	}

	scraperService := scraper.ScraperService{
		Keyword: keyword,
	}
	err = scraperService.Run()
	if err != nil {
		logs.Error("Error when scraping keyword: %v", err.Error())

		return err
	}

	return nil
}
