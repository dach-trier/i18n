package i18n_html

import (
	"bytes"
	"fmt"
	"html/template"

	"golang.org/x/text/language"

	"github.com/dach-trier/i18n"
)

func FuncMap(bundle i18n.Bundle) template.FuncMap {
	return template.FuncMap{
		"t": func(message string, lang language.Tag, data any) (template.HTML, error) {
			if translation, ok := i18n.Translate(message, lang, bundle); ok {
				var html bytes.Buffer
				var tmpl *template.Template
				var err error

				if tmpl, err = template.New("").Parse(translation); err != nil {
					return "", err
				}

				if err = tmpl.Execute(&html, data); err != nil {
					return "", err
				}

				return template.HTML(html.String()), nil
			}

			return "", fmt.Errorf("missing translation for %q in language %q", message, lang)
		},
	}
}
