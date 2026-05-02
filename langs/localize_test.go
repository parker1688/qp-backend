package langs

import (
	"testing"
)

func TestEmbed(t *testing.T) {
	t.Log(locales.ReadDir("locales"))
}

func TestLang(t *testing.T) {
	t.Log(GetWithLocale("en", "hello"))
	t.Log(GetWithLocale("th", "hello_firstname_lastname", &Replacements{"firstname": "steve", "lastname": "steve"}))
}
