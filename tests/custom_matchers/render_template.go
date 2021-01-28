package custom_matchers

import (
	"fmt"
	"net/http/httptest"
	"strings"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

func RenderTemplate(templateString string) types.GomegaMatcher {
	return &RenderTemplateMatcher{
		templateString: templateString,
	}
}

type RenderTemplateMatcher struct {
	templateString string
}

func (matcher *RenderTemplateMatcher) Match(actual interface{}) (bool, error) {
	response, ok := actual.(*httptest.ResponseRecorder)

	if !ok {
		return false, fmt.Errorf("RenderTemplate must be passed a *httptest.ResponseRecorder. Got\n%s", format.Object(actual, 1))
	}

	classList := strings.Join(strings.Split(matcher.templateString, "#"), " ")
	return strings.Contains(response.Body.String(), classList), nil
}

func (matcher *RenderTemplateMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected response to render `%s` template", matcher.templateString)
}

func (matcher *RenderTemplateMatcher) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected response not to render `%s` template", matcher.templateString)
}
