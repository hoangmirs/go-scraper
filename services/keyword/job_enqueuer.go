package keyword

import (
	"github.com/hoangmirs/go-scraper/db"
	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocraft/work"
)

var enqueuer *work.Enqueuer

func EnqueueKeyword(keyword models.Keyword) error {
	setUpEnqueuer()
	job, err := enqueuer.Enqueue("scrape", work.Q{"keywordID": keyword.Id})

	if err != nil {
		return err
	}

	logs.Info("Enqueued %v job for keyword %v", job.Name, job.ArgString("keyword"))

	return nil
}

func setUpEnqueuer() {
	if enqueuer == nil {
		enqueuer = work.NewEnqueuer("go-scraper", db.GetRedisPool())
	}
}
