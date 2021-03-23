package database

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
	"github.com/soveran/redisurl"
)

var RedisPool *redis.Pool

func SetupRedisPool() {
	pool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial:      GetRedisConnection,
	}

	RedisPool = pool
}

func GetRedisConnection() (redis.Conn, error) {
	// Use redisurl to connect redis on production
	if web.AppConfig.DefaultString("runmode", "dev") == "prod" {
		return redisurl.Connect()
	}

	return redis.Dial("tcp", web.AppConfig.DefaultString("redisUrl", ""))
}

func GetRedisPool() *redis.Pool {
	return RedisPool
}
