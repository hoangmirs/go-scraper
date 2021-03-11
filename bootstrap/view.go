package bootstrap

import (
	"html/template"
	"io/ioutil"
	"log"

	"github.com/beego/beego/v2/server/web"
)

// SetUpTemplateFunction register additional template functions
func SetUpTemplateFunction() {
	templateFunctions := map[string]interface{}{
		"render_file": renderFile,
		"render_icon": renderIcon,
	}

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

func renderIcon(iconName string) template.HTML {
	iconTemplate := `<svg class="icon" viewBox="0 0 16 16">
		<use xlink:href="#` + iconName + `" />
	</svg>`

	return web.Str2html(iconTemplate)
}
