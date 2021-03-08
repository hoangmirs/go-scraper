package test_helpers

import (
	"io"
	"net/http"
	"net/http/httptest"
	neturl "net/url"
	"strings"

	"github.com/hoangmirs/go-scraper/tests/fabricators"

	"github.com/beego/beego/v2/server/web"
	"github.com/onsi/ginkgo"
)

type UserInfo struct {
	Email    string
	Password string
}

// MakeRequest makes a HTTP request and returns response
func MakeRequest(method string, url string, body io.Reader) *httptest.ResponseRecorder {
	return makeRequest(method, url, nil, body, nil)
}

// MakeAuthenticatedRequest makes a HTTP request with authenticated user and returns response
func MakeAuthenticatedRequest(method string, url string, headers http.Header, body io.Reader, userInfo *UserInfo) *httptest.ResponseRecorder {
	return makeRequest(method, url, headers, body, userInfo)
}

func makeRequest(method string, url string, headers http.Header, body io.Reader, userInfo *UserInfo) *httptest.ResponseRecorder {
	request, err := http.NewRequest(method, url, body)
	if headers != nil {
		request.Header = headers
	}

	if err != nil {
		ginkgo.Fail("Create request failed: " + err.Error())
	}

	if body != nil {
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	response := httptest.NewRecorder()

	if userInfo != nil {
		authenticationCookie := getAuthenticationCookie(userInfo)
		request.Header.Set("Cookie", authenticationCookie.String())
	}

	web.BeeApp.Handlers.ServeHTTP(response, request)

	return response
}

func getAuthenticationCookie(userInfo *UserInfo) *http.Cookie {
	_, err := fabricators.FabricateUser(userInfo.Email, userInfo.Password)
	if err != nil {
		ginkgo.Fail("Failed to fabricate user: " + err.Error())
	}

	form := neturl.Values{
		"email":    {userInfo.Email},
		"password": {userInfo.Password},
	}
	body := strings.NewReader(form.Encode())

	response := MakeRequest("POST", "/login", body)

	return response.Result().Cookies()[0]
}
