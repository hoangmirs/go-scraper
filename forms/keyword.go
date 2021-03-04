package forms

import (
	"encoding/csv"
	"errors"
	"io"
	"mime/multipart"

	"github.com/hoangmirs/go-scraper/models"
	keywordservice "github.com/hoangmirs/go-scraper/services/keyword"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type KeywordForm struct {
	File       multipart.File        `valid:"Required"`
	FileHeader *multipart.FileHeader `valid:"Required"`
	Keywords   []string
	User       *models.User `valid:"Required"`
}

func (keywordForm *KeywordForm) Valid(v *validation.Validation) {
	err := keywordForm.validateCSVFileType()
	if err != nil {
		_ = v.SetError("File", err.Error())
	}

	err = keywordForm.readCSVFile()
	if err != nil {
		_ = v.SetError("File", err.Error())
	}

	err = keywordForm.validateCSVLength()
	if err != nil {
		_ = v.SetError("File", err.Error())
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

	keywordService := keywordservice.KeywordService{
		Keywords: keywordForm.Keywords,
		User:     keywordForm.User,
	}
	err = keywordService.Run()
	if err != nil {
		logs.Error("Run error:", err)
	}

	return nil
}

func (keywordForm *KeywordForm) readCSVFile() error {
	reader := csv.NewReader(keywordForm.File)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.New("File cannot be opened")
		}
		keywordForm.Keywords = append(keywordForm.Keywords, row[0])
	}

	return nil
}

func (keywordForm *KeywordForm) validateCSVLength() error {
	keywordLength := len(keywordForm.Keywords)
	if keywordLength <= 0 || keywordLength > 1000 {
		return errors.New("CSV file only accepts from 1 to 1000 keywords")
	}

	return nil
}

func (keywordForm *KeywordForm) validateCSVFileType() error {
	if keywordForm.FileHeader.Header.Get("Content-Type") != "text/csv" {
		return errors.New("File type is not supported")
	}

	return nil
}
