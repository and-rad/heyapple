package api

// Translator defines functions for text localization.
type Translator interface {
	Translate(input interface{}, lang string) string
	Default() string
	Get(lang string) map[string]string
	Set(map[string]map[string]string) error
}
