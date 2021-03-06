package test_helpers

import (
	"github.com/hoangmirs/go-scraper/conf"
	"github.com/hoangmirs/go-scraper/database"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocraft/work"
)

func GetWorkerClient() *work.Client {
	return work.NewClient(conf.GetString("workerNamespace"), database.GetRedisPool())
}

func DeleteRedisJobs() {
	_, err := database.GetRedisPool().Get().Do("DEL", redisKeyJobs(conf.GetString("workerNamespace"), conf.GetString("scraperJobName")))
	if err != nil {
		logs.Error("Error when deleting redis jobs: %v", err)
	}
}

func redisKeyJobs(namespace, jobName string) string {
	return redisKeyJobsPrefix(namespace) + jobName
}

func redisKeyJobsPrefix(namespace string) string {
	return redisNamespacePrefix(namespace) + "jobs:"
}

func redisNamespacePrefix(namespace string) string {
	l := len(namespace)
	if (l > 0) && (namespace[l-1] != ':') {
		namespace = namespace + ":"
	}
	return namespace
}
