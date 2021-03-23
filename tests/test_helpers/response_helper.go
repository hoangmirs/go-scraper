package test_helpers

import (
	"io/ioutil"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
)

func GetBody(response *httptest.ResponseRecorder) string {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Fail("Error when reading reponse bod: " + err.Error())
	}

	return string(body)
}
