package i18n_template

import (
	"fmt"
	"golang.org/x/text/language"
	"html/template"

	"github.com/dach-trier/i18n"
)

func FuncMap(bundle i18n.Bundle) template.FuncMap {
	return template.FuncMap{
		"t": func(message string, lang language.Tag, args ...any) (template.HTML, error) {
			if translation, ok := i18n.Translate(message, lang, bundle); ok {
				return template.HTML(translation), nil
			}

			return "", fmt.Errorf("missing translation for %q in language %q", message, lang)
		},
	}
}
