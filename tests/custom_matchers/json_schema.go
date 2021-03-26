package custom_matchers

import (
	"fmt"
	"net/http/httptest"

	"github.com/hoangmirs/go-scraper/helpers"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/xeipuuv/gojsonschema"
)

type LoaderType int

const (
	StringLoader LoaderType = iota
	ReferenceLoader
)

// MatchJSONSchema checks if JSON document matches
// the given JSON Schema.
func MatchJSONSchema(fileName string) types.GomegaMatcher {
	return &matchJSONSchemaMatcher{
		schemaLoader: getSchemaLoader(fileName),
	}
}

type matchJSONSchemaMatcher struct {
	schemaLoader gojsonschema.JSONLoader
	responseBody string
	errorMessage string
}

func (m *matchJSONSchemaMatcher) Match(actual interface{}) (success bool, err error) {
	response, ok := actual.(*httptest.ResponseRecorder)

	if !ok {
		return false, fmt.Errorf("MatchJSONSchema must be passed a *httptest.ResponseRecorder. Got\n%s", format.Object(actual, 1))
	}

	m.responseBody = response.Body.String()
	documentLoader := gojsonschema.NewStringLoader(m.responseBody)

	result, err := gojsonschema.Validate(m.schemaLoader, documentLoader)
	if err != nil {
		return false, fmt.Errorf("Failed to validate JSON: %s", err.Error())
	}

	if !result.Valid() {
		for _, desc := range result.Errors() {
			m.errorMessage += fmt.Sprintf("- %s\n", desc)
		}
	}
	return result.Valid(), nil

}

func (m *matchJSONSchemaMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%v\nto match schema\n\t%#v\nSee the errors:\n%s", m.responseBody, m.schemaLoader.JsonSource(), m.errorMessage)
}

func (m *matchJSONSchemaMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%v\nnot to match schema\n\t%#v\nSee the errors:\n%s", m.responseBody, m.schemaLoader.JsonSource(), m.errorMessage)
}

func getSchemaLoader(fileName string) gojsonschema.JSONLoader {
	filePathFormat := "file://%s/tests/api/schemas/%s.json"
	filePath := fmt.Sprintf(filePathFormat, helpers.RootDir(), fileName)
	jsonLoader := gojsonschema.NewReferenceLoader(filePath)

	return jsonLoader
}
