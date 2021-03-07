package test_helpers

import (
	"fmt"
	"mime/multipart"

	"github.com/hoangmirs/go-scraper/helpers"
)

func GetMultipartInfoFromFile(pathToFile string, contentType string) (multipart.File, *multipart.FileHeader, error) {
	realPath := fmt.Sprintf("%s/%s", helpers.RootDir(), pathToFile)
	req, err := createUploadRequest(realPath)
	if err != nil {
		return nil, nil, err
	}

	file, fileHeader, err := req.FormFile("file")
	fileHeader.Header.Set("Content-Type", contentType)

	return file, fileHeader, err
}
