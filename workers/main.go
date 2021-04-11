package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hoangmirs/go-scraper/bootstrap"
	"github.com/hoangmirs/go-scraper/conf"
	"github.com/hoangmirs/go-scraper/database"
	"github.com/hoangmirs/go-scraper/workers/jobs"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

// Make a redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial:      database.GetRedisConnection,
}

func init() {
	bootstrap.SetUp()
}

func main() {
	pool := work.NewWorkerPool(jobs.Context{}, 5, conf.GetString("workerNamespace"), redisPool)

	pool.Middleware((*jobs.Context).ScraperLog)

	pool.JobWithOptions(conf.GetString("scraperJobName"), work.JobOptions{MaxFails: jobs.MaxFails}, (*jobs.Context).PerformScrape)

	pool.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	pool.Stop()
}
