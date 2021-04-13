package apiv1controllers

import (
	"net/http"

	"github.com/hoangmirs/go-scraper/conf"
	"github.com/hoangmirs/go-scraper/forms"
	"github.com/hoangmirs/go-scraper/models"
	"github.com/hoangmirs/go-scraper/presenters"
	v1serializers "github.com/hoangmirs/go-scraper/serializers/v1"

	"github.com/beego/beego/v2/server/web/pagination"
)

type Keyword struct {
	baseController
}

func (c *Keyword) NestPrepare() {
	c.requireAuthenticatedUser = true
}

func (c *Keyword) Get() {
	query := map[string]interface{}{
		"user_id": c.CurrentUser.Id,
		"order":   "-id",
		"keyword__icontains": c.GetString("keyword"),
	}

	keywordsCount, err := models.GetKeywordsCount(query)
	if err != nil {
		c.renderGenericError(err)
	}

	keywordsPerPage := conf.GetInt("perPage")
	paginator := pagination.SetPaginator(c.Ctx, keywordsPerPage, keywordsCount)
	query["limit"] = keywordsPerPage
	query["offset"] = paginator.Offset()

	keywords, err := models.GetKeywords(query)
	if err != nil {
		c.renderGenericError(err)
	}

	keywordsSerializer := v1serializers.KeywordList{
		Keywords: keywords,
		Pagination: v1serializers.Pagination{
			Paginator: paginator,
		},
	}

	err = c.renderListJSON(keywordsSerializer.Data(), keywordsSerializer.Meta(), keywordsSerializer.Links())
	if err != nil {
		c.renderGenericError(err)
	}
}

func (c *Keyword) Show() {
	query := map[string]interface{}{
		"id":      c.Ctx.Input.Param(":id"),
		"user_id": c.CurrentUser.Id,
	}

	keyword, err := models.GetKeywordByQuery(query)
	if err != nil {
		c.renderError("Not found", err.Error(), "record_not_found", http.StatusNotFound, nil)
	}

	keywordPresenter := presenters.KeywordPresenter{Keyword: keyword}
	keywordPresenter.ConvertKeywordLinks()

	keywordSerializer := v1serializers.Keyword{
		Keyword: keyword,
		Links:   keywordPresenter.Links,
	}

	err = c.renderJSON(keywordSerializer.Data())
	if err != nil {
		c.renderGenericError(err)
	}
}

func (c *Keyword) Post() {
	file, fileHeader, err := c.GetFile("file")
	if err != nil {
		c.renderError("File required", err.Error(), "validation_error", http.StatusUnprocessableEntity, nil)
	}

	keywordForm := forms.KeywordForm{
		File:       file,
		FileHeader: fileHeader,
		User:       c.CurrentUser,
	}

	err = keywordForm.Save()
	if err != nil {
		c.renderError(err.Error(), err.Error(), "validation_error", http.StatusUnprocessableEntity, nil)
	} else {
		c.Ctx.Output.Status = http.StatusCreated
	}
}
