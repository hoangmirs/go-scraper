package main

import (
	"github.com/beego/beego/v2/server/web"

	_ "github.com/hoangmirs/go-scraper/bootstrap"
)

func main() {
	web.Run()
}
