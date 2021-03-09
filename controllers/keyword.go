package controllers

import (
	"html/template"
	"net/http"

	"github.com/hoangmirs/go-scraper/forms"
	"github.com/hoangmirs/go-scraper/models"

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

	keywords, err := models.GetKeywords(c.CurrentUser)
	if err != nil {
		logs.Error("Error when fetching keywords: %v", err)
	}

	c.Data["Keywords"] = keywords

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
		User:       c.CurrentUser,
	}

	err = keywordForm.Save()
	if err != nil {
		flash.Error(err.Error())
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
