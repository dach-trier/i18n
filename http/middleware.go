package i18n_http

import (
	"context"
	"net/http"

	"golang.org/x/text/language"
)

func Middleware(matcher func(tags ...language.Tag) language.Tag) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				tags := []language.Tag{}

				if tag, err := language.Parse(r.FormValue("lang")); err == nil {
					tags = append(tags, tag)
				}

				if tag, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language")); err == nil {
					tags = append(tags, tag...)
				}

				ctx := context.WithValue(r.Context(), "lang", matcher(tags...))
				next.ServeHTTP(w, r.WithContext(ctx))
			},
		)
	}
}
