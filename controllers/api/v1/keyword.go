package apiv1controllers

import (
	"net/http"

	"github.com/hoangmirs/go-scraper/forms"

	"github.com/beego/beego/v2/core/logs"
)

type Keyword struct {
	baseController
}

func (c *Keyword) NestPrepare() {
	c.requireAuthenticatedUser = true
}

func (c *Keyword) Post() {
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
		c.renderGenericError(err)
	} else {
		c.Ctx.Output.Status = http.StatusCreated
	}
}
