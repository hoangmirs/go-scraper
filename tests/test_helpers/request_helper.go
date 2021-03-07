package test_helpers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	"github.com/beego/beego/v2/server/web"
	"github.com/onsi/ginkgo"
)

// MakeRequest makes a HTTP request and returns response
func MakeRequest(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		ginkgo.Fail("Create request failed: " + err.Error())
	}

	if body != nil {
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	response := httptest.NewRecorder()
	web.BeeApp.Handlers.ServeHTTP(response, request)

	return response
}

func createUploadRequest(pathToFile string) (*http.Request, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(pathToFile))
	if err != nil {
		return nil, err
	}
	_, _ = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}
	// the body is the only important data for creating a new request with the form data attached
	req, err := http.NewRequest("POST", "", body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}
