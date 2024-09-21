package localizations

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
)

type Replacements map[string]interface{}

type Localizer struct {
	Locale          string
	FallbackLocale  string
	LocalizationSrc []string
	Localizations   map[string]string
}

func (t *Localizer) SetupLocalization() {

	dt := make(map[string]string)

	for _, inputDir := range t.LocalizationSrc {

		// read files
		files := getAllFiles(inputDir)

		localizations, err := generateLocalizations(files, inputDir)
		if err != nil {

			log.Printf("failed to generate localization from directory %s | %s ", inputDir, err.Error())
			continue
		}

		for k, v := range localizations {

			dt[k] = v
		}
	}

	t.Localizations = dt

}

func (t *Localizer) SetSource(src ...string) {

	t.LocalizationSrc = src
}

func (t *Localizer) SetDefaultLocale(locale, fallback string) {

	t.Locale = locale
	t.FallbackLocale = fallback

}

func (t *Localizer) SetLocales(locale, fallback string) *Localizer {

	t.Locale = locale
	t.FallbackLocale = fallback
	return t
}

func (t *Localizer) SetLocale(locale string) *Localizer {
	t.Locale = locale
	return t
}

func (t *Localizer) SetFallbackLocale(fallback string) *Localizer {

	t.FallbackLocale = fallback
	return t
}

func (t *Localizer) Translate(key string, replacements ...map[string]string) string {

	lkey := t.getLocalizationKey(t.Locale, key)
	log.Printf("got localization key %s ", lkey)

	str, ok := t.Localizations[lkey]
	if !ok {

		str, ok = t.Localizations[t.getLocalizationKey(t.FallbackLocale, key)]
		if !ok {

			return key
		}
	}

	// If the str doesn't have any substitutions, no need to
	// template.Execute.
	if strings.Index(str, "}}") == -1 {

		return str
	}

	mergedReplacements := make(map[string]string)

	for _, j := range replacements {

		for k, v := range j {

			mergedReplacements[k] = v
		}
	}

	return t.replace(str, mergedReplacements)
}

func (t *Localizer) getLocalizationKey(locale string, key string) string {

	/*
		// insert locale as the 2nd last item of the key
		parts := strings.Split(key,".")
		size := len(parts)

		var newArrKey []string

		for k,v := range parts {

			if k == size - 1 {

				newArrKey = append(newArrKey, locale)
			}

			newArrKey = append(newArrKey, v)

		}

		return strings.Join(newArrKey,".")

	*/
	return fmt.Sprintf("%v.%v", locale, key)
}

func (t *Localizer) replace(str string, replacements map[string]string) string {

	b := &bytes.Buffer{}
	tmpl, err := template.New("").Parse(str)
	if err != nil {

		return str
	}

	err = template.Must(tmpl, err).Execute(b, replacements)
	if err != nil {

		return str

	}

	buff := b.String()

	return buff
}
