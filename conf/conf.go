package conf

import "github.com/beego/beego/v2/server/web"

func GetString(key string) string {
	return web.AppConfig.DefaultString(key, "")
}

func GetInt(key string) int {
	return web.AppConfig.DefaultInt(key, 0)
}
