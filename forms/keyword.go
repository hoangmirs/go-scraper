package forms

import (
	"mime/multipart"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type KeywordForm struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

func (keywordForm *KeywordForm) Valid(v *validation.Validation) {
	if keywordForm.FileHeader.Header.Get("Content-Type") != "text/csv" {
		_ = v.SetError("File", "File type is not supported")
	}
}

func (keywordForm *KeywordForm) Save() error {
	valid := validation.Validation{}

	success, err := valid.Valid(keywordForm)
	if err != nil {
		logs.Error("Validation error:", err)
	}

	if !success {
		for _, err := range valid.Errors {
			return err
		}
	}

	return nil
}
