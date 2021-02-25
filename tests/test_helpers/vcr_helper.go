package test_helpers

import (
	"fmt"

	"github.com/hoangmirs/go-scraper/helpers"
)

func CassettePath(cassetteName string) string {
	return fmt.Sprintf("%s/tests/fixtures/vcr/%s", helpers.RootDir(), cassetteName)
}
