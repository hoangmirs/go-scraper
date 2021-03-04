package db

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func SetupRedisPool() {
	pool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", web.AppConfig.DefaultString("redisUrl", ""))
		},
	}

	RedisPool = pool
}

func GetRedisPool() *redis.Pool {
	return RedisPool
}
