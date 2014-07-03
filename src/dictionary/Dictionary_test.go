package dictionary

import (
	"testing"
)

func TestCorrectTranslation(t *testing.T) {
	var d Dictionary = Make("../../data/de-en.json")
	fixtures := map[string]string{
		"zeit":    "time",
		"fünf":    "five",
		"zukunft": "future",
		"apfel":   "apple",
	}

	for word, correctMeaning := range fixtures {
		if meaning := d.Lookup(word); correctMeaning != meaning {
			t.Error("expected meaning for '" + word + "' is '" + correctMeaning + "' but got '" + meaning + "'")
		}
	}
}
