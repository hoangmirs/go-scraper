package bootstrap

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"

	"github.com/beego/beego/v2/server/web"
)

var templateFunctions = map[string]interface{}{
	"render_file": renderFile,
	"render_icon": renderIcon,
}

// SetUpTemplateFunction register additional template functions
func SetUpTemplateFunction() {
	for name, fn := range templateFunctions {
		err := web.AddFuncMap(name, fn)
		if err != nil {
			log.Fatal("Failed to add template function", err.Error())
		}
	}
}

func renderFile(path string) template.HTML {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return web.Str2html(string(content))
}

func renderIcon(iconName string, options ...string) template.HTML {
	iconTemplate := `<svg class="icon %s" viewBox="0 0 16 16">
		<use xlink:href="#%s" />
	</svg>`

	var classList string
	if len(options) > 0 {
		classList = options[0]
	}

	htmlString := fmt.Sprintf(iconTemplate, classList, iconName)

	return web.Str2html(htmlString)
}
