package i18n

import "golang.org/x/text/language"

// Returns the translation of the given message for the specified target language.
//
// It looks up the message in the provided bundle and then searches for a
// translation matching the target language tag.
//
// The returned boolean indicates whether a translation was found. If either the message
// or the language-specific translation is missing, it returns an empty string and false.
//
// Use MustTranslate if the translation is known to exist.
func Translate(message string, targetLanguage language.Tag, bundle Bundle) (string, bool) {
	if targetLanguage == language.English {
		return message, true
	}

	var translations map[language.Tag]string
	var translation string
	var ok bool

	if translations, ok = bundle[message]; !ok {
		return "", false
	}

	if translation, ok = translations[targetLanguage]; !ok {
		return "", false
	}

	return translation, true
}

// Returns the translation of the given message for the specified target language.
//
// This method panics if the translation is not found.
//
// Use Translate when a translation is not guaranteed to exist.
func MustTranslate(message string, targetLanguage language.Tag, bundle Bundle) string {
	if translation, ok := Translate(message, targetLanguage, bundle); ok {
		return translation
	}

	panic("translation missing")
}
