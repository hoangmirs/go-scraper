package controllers

import (
	"html/template"
	"net/http"

	"github.com/hoangmirs/go-scraper/conf"
	"github.com/hoangmirs/go-scraper/forms"
	"github.com/hoangmirs/go-scraper/models"
	"github.com/hoangmirs/go-scraper/presenters"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/pagination"
)

type Keyword struct {
	baseController
}

func (c *Keyword) NestPrepare() {
	c.requireAuthenticatedUser = true
}

func (c *Keyword) Get() {
	web.ReadFromRequest(&c.Controller)
	flash := web.NewFlash()

	c.renderKeywordView(flash)

	flash.Store(&c.Controller)
}

func (c *Keyword) Show() {
	keyword, err := c.getKeyword()
	if err != nil {
		logs.Error("Error when getting keyword: %v", err)
		c.Redirect("/", http.StatusNotFound)
		return
	}

	keywordPresenter := presenters.KeywordPresenter{Keyword: keyword}
	keywordPresenter.ConvertKeywordLinks()

	c.Data["KeywordPresenter"] = keywordPresenter
	c.TplName = "keyword/show.html"
}

func (c *Keyword) ShowHTML() {
	keyword, err := c.getKeyword()
	if err != nil {
		logs.Error("Error when getting keyword: %v", err)
		c.Redirect("/", http.StatusNotFound)
		return
	}

	err = c.Ctx.Output.Body([]byte(keyword.HtmlCode))
	if err != nil {
		logs.Error("Error when setting body: %v", err)
	}
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

		c.Data["Form"] = keywordForm

		c.renderKeywordView(flash)

		flash.Store(&c.Controller)
	} else {
		flash.Success("Processing uploaded keywords")
		flash.Store(&c.Controller)

		c.Ctx.Redirect(http.StatusFound, "/keyword")
	}
}

func (c *Keyword) renderKeywordView(flash *web.FlashData) {
	var keyword string
	err := c.Ctx.Input.Bind(&keyword, "keyword")
	if err != nil {
		logs.Error("Error when getting keyword input: %v", err.Error())
	}

	query := map[string]interface{}{
		"user_id":            c.CurrentUser.Id,
		"order":              "-id",
		"keyword__icontains": keyword,
	}

	keywordsCount, err := models.GetKeywordsCount(query)
	if err != nil {
		logs.Error("Error when getting keywords count: %v", err)
		flash.Error(err.Error())
	}

	keywordsPerPage := conf.GetInt("perPage")
	paginator := pagination.SetPaginator(c.Ctx, keywordsPerPage, keywordsCount)
	query["limit"] = keywordsPerPage
	query["offset"] = paginator.Offset()

	keywords, err := models.GetKeywords(query)
	if err != nil {
		logs.Error("Error when fetching keywords: %v", err)
		flash.Error(err.Error())
	}

	c.Data["Keyword"] = keyword
	c.Data["Keywords"] = keywords
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "keyword/index.html"

	c.Data["Title"] = "Keyword"
}

func (c *Keyword) getKeyword() (*models.Keyword, error) {
	keywordId := c.Ctx.Input.Param(":id")
	query := map[string]interface{}{
		"id":      keywordId,
		"user_id": c.CurrentUser.Id,
	}

	return models.GetKeywordByQuery(query)
}
