package apiv1controllers

import (
	"net/http"

	"github.com/hoangmirs/go-scraper/forms"
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
		c.renderError("File required", err.Error(), "validation_error", http.StatusUnprocessableEntity, nil)
		return
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
