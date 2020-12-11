package main

import (
	beego "github.com/astaxie/beego/server/web"
	_ "github.com/hoangmirs/go-scraper/routers"
)

func main() {
	beego.Run()
}
