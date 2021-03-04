package keyword

// import (
// 	"github.com/hoangmirs/go-scraper/bootstrap"
// 	"github.com/hoangmirs/go-scraper/models"

// 	"github.com/beego/beego/v2/core/logs"
// 	"github.com/gocraft/work"
// )

// var enqueuer = work.NewEnqueuer("go-scraper", bootstrap.GetRedisPool())

// func EnqueueKeyword(keyword models.Keyword) error {
// 	job, err := enqueuer.Enqueue("scraper", work.Q{"address": "test@example.com", "subject": "hello world", "customer_id": 4})

// 	if err != nil {
// 		return err
// 	}

// 	logs.Info("Enqueued %v job for keyword %v", job.Name, job.ArgString("keyword"))

// 	return nil
// }
