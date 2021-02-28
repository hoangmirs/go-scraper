package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/hoangmirs/go-scraper/forms"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type Keyword struct {
	baseController
}

func (c *Keyword) NestPrepare() {
	c.requireAuthenticatedUser = true
}

func (c *Keyword) Get() {
	web.ReadFromRequest(&c.Controller)

	c.renderKeywordView()
}

func (c *Keyword) Post() {
	flash := web.NewFlash()

	file, fileHeader, err := c.GetFile("file")
	if err != nil {
		logs.Error("Error when getting file: %v", err)
	}

	keywordForm := forms.KeywordForm{
		File:       file,
		FileHeader: fileHeader,
	}

	err = keywordForm.Save()
	if err != nil {
		flash.Error(fmt.Sprint(err))
		flash.Store(&c.Controller)

		c.Data["Form"] = keywordForm

		c.renderKeywordView()
	} else {
		flash.Success("Processing uploaded keywords")
		flash.Store(&c.Controller)

		c.Ctx.Redirect(http.StatusFound, "/keyword")
	}
}

func (c *Keyword) renderKeywordView() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/application.html"
	c.TplName = "keyword/index.html"

	c.Data["Title"] = "Keyword"
}
