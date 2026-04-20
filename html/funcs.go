package i18n_html

import (
	"fmt"
	"html/template"

	"golang.org/x/text/language"

	"github.com/dach-trier/i18n"
)

func FuncMap(bundle i18n.Bundle) template.FuncMap {
	return template.FuncMap{
		"t": func(message string, lang language.Tag, args ...any) (template.HTML, error) {
			if translation, ok := i18n.Translate(message, lang, bundle); ok {
				return template.HTML(fmt.Sprintf(translation, args...)), nil
			}

			return "", fmt.Errorf("missing translation for %q in language %q", message, lang)
		},
	}
}
