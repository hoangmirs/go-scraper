package test_helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
)

func GetBody(response *httptest.ResponseRecorder) string {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Fail("Error when reading reponse body: " + err.Error())
	}

	return string(body)
}

func GetJSONBody(response *httptest.ResponseRecorder) map[string]interface{} {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Fail("Error when reading reponse body: " + err.Error())
	}
	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		Fail("Error when unmarshalling: " + err.Error())
	}

	return data
}
