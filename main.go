package main

import (
	_ "github.com/hoangmirs/go-scraper/routers"
	beego "github.com/astaxie/beego/server/web"
)

func main() {
	beego.Run()
}
