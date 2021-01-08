package main

import (
	"github.com/hoangmirs/go-scraper/bootstrap"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	bootstrap.SetUp()

	web.Run()
}
