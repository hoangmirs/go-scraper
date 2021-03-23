package jobenqueuer

import (
	"errors"

	"github.com/hoangmirs/go-scraper/conf"
	"github.com/hoangmirs/go-scraper/database"
	"github.com/hoangmirs/go-scraper/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocraft/work"
)

var enqueuer *work.Enqueuer

// SetUpEnqueuer sets up a new job enqueuer, will be run when initializing application
func SetUpEnqueuer() {
	if enqueuer == nil {
		enqueuer = work.NewEnqueuer(conf.GetString("workerNamespace"), database.GetRedisPool())
	}
}

// EnqueueKeyword enqueues a keyword
func EnqueueKeyword(keyword *models.Keyword) error {
	if keyword == nil {
		return errors.New("Keyword cannot be nil")
	}
	job, err := enqueuer.Enqueue(conf.GetString("scraperJobName"), work.Q{"keywordID": keyword.Id})

	if err != nil {
		return err
	}

	logs.Info("Enqueued %v job for keyword %v", job.Name, job.ArgString("keyword"))

	return nil
}
