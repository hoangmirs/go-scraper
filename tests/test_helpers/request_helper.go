package test_helpers

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/beego/beego/v2/server/web"
	"github.com/onsi/ginkgo"
)

// MakeRequest make a HTTP request and return response
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
