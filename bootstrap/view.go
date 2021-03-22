package bootstrap

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"

	"github.com/beego/beego/v2/server/web"
)

var templateFunctions = map[string]interface{}{
	"render_file": renderFile,
	"render_icon": renderIcon,
	"args":        args,
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

func args(values ...interface{}) (map[string]interface{}, error) {
	if len(values) % 2 != 0 {
		return nil, errors.New("invalid args call")
	}

	args := make(map[string]interface{}, len(values) / 2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("args keys must be strings")
		}
		args[key] = values[i + 1]
	}

	return args, nil
}
