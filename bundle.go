package i18n

import "golang.org/x/text/language"

// A collection of messages mapped to their translations. The source language
// is implicitly English.
type Bundle map[string]map[language.Tag]string

func NewBundle() Bundle {
	return Bundle(make(map[string]map[language.Tag]string))
}

// Adds a translation of the original message in the specified language.
//
// Returns true if the translation was added. If a translation for the message
// in the provided language already exists, it is not modified and false is
// returned.
//
// Use MustAddTranslation if translation collisions are not expected.
func (bundle Bundle) AddTranslation(original string, translation string, lang language.Tag) bool {
	if _, ok := bundle[original]; !ok {
		bundle[original] = make(map[language.Tag]string)
	}

	if _, ok := bundle[original][lang]; !ok {
		bundle[original][lang] = translation
		return true
	}

	return false
}

// Adds a translation of the original message in the specified language.
//
// This method panics if a translation for the message in the provided language
// already exists.
//
// Use AddTranslation if translation collisions are expected.
func (bundle Bundle) MustAddTranslation(original string, translation string, lang language.Tag) {
	ok := bundle.AddTranslation(original, translation, lang)

	if !ok {
		panic("translation collision")
	}
}
