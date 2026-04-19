package i18n

import (
	"fmt"
	"io/fs"

	"github.com/goccy/go-yaml"
	"golang.org/x/text/language"
)

type entry struct {
	Message     string `yaml:"message"`
	Translation string `yaml:"translation"`
}

// Loads translations from a YAML file in the provided filesystem and adds them to the bundle.
//
// The file is expected to contain a list of entries, where each entry defines
// a message and its translation.
//
// If a translation collision occurs, this method returns an error. If
// translation collisions are not expected, use MustLoadYaml.
//
// Example YAML format:
//
//   - message:     "hello"
//     translation: "hola"
//   - message:     "goodbye"
//     translation: "adiós"
func (bundle Bundle) LoadYaml(filesystem fs.FS, path string, lang language.Tag) error {
	var entries []entry
	var data []byte
	var err error
	var ok bool

	if data, err = fs.ReadFile(filesystem, path); err != nil {
		return err
	}

	if err = yaml.Unmarshal(data, &entries); err != nil {
		return err
	}

	for _, entry := range entries {
		if ok = bundle.AddTranslation(entry.Message, entry.Translation, lang); !ok {
			return fmt.Errorf("translation collision for message %q in language %q", entry.Message, lang)
		}
	}

	return nil
}

// Like LoadYaml but panics on error or a translation collision.
func (bundle Bundle) MustLoadYaml(filesystem fs.FS, path string, lang language.Tag) {
	if err := bundle.LoadYaml(filesystem, path, lang); err != nil {
		panic(err)
	}
}
