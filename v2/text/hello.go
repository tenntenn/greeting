package text

import (
	"os"

	"golang.org/x/text/language"
)

func DefaultLang() language.Tag {
	for _, k := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if env := os.Getenv(k); env != "" {
			return language.Make(env)
		}
	}
	return language.English
}

func GoodMorning(lang language.Tag) string {
	switch lang {
	case language.Japanese:
		return "おはよう"
	}
	return "Good Morning"
}

func Hello(lang language.Tag) string {
	switch lang {
	case language.Japanese:
		return "こんにちは"
	}
	return "Hello"
}

func GoodEvening(lang language.Tag) string {
	switch lang {
	case language.Japanese:
		return "こんばんは"
	}
	return "Good Evening"
}
