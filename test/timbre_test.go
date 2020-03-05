package test

import (
	"github.com/adaminoue/timbre"
	"testing"
)

// tests for Pluralize (not exhaustive; concerned users are encouraged to test by inspection via word_list_test.go)
func TestPluralize(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"example", "examples"},
		{"criterion", "criteria"},
		{"potato", "potatoes"},
		{"FBI", "FBI"},
		{"bus", "buses"},
		{"running", "running"},
		{"man", "men"},
		{"policeman", "policemen"},
		{"city", "cities"},
		{"analysis", "analyses"},
		{"  ", "  "},
		{"trail  ", "trails  "},
		{"  leading", "  leading"},
		{"  &^", "  &^"},
		{"^trail ^ ", "^trails ^ "},
		{"^American", "^Americans"},
		{"^man", "^men"},
		{" ^man  **", " ^men  **"},
		{" % leading(", " % leading("},
		{"x", "xes"},
		{"X", "X"},
		{"Capitol", "Capitols"},
		{"Potato", "Potatoes"},
		{"ToMaTo", "ToMaToes"},
		{"CalTrain", "CalTrains"},
		{"PainTing", "Paintings"},
		{"citY", "cities"},
		{"CIty", "CIties"},
		{"ANalysis", "ANalyses"},
		{"PolicEman", "PolicEmen"},
		{"COding", "COding"},
	}
	for _, test := range tests {
		timbre.Pluralize(&test.input)
		if test.input != test.want {
			t.Log("Failed on output", test.input, "which should be", test.want)
			t.Fail()
		}
	}
}

func TestPluralizeNth(t *testing.T) {
	passTests := []struct {
		input string
		nth   int
		want  string
	}{
		{"The quick brown fox jumped over the lazy dog",3, "The quick brown foxes jumped over the lazy dog"},
		{" The quick brown fox jumped over the lazy dog. ", 8, "The quick brown fox jumped over the lazy dogs."},
		{"Raindrop keep falling on my head", 0, "Raindrops keep falling on my head"},
		{"Raindrop keep falling on my head", 1, "Raindrop keeps falling on my head"},
	}
	for _, test := range passTests {
		err := timbre.PluralizeNth(&test.input, test.nth)
		if err != nil {
			t.Error("pluralization error (output not checked)")
		}
		if test.input != test.want {
			t.Log("Failed on output", test.input, "which should be", test.want)
			t.Fail()
		}
	}
}

func TestPluralizeNthError(t *testing.T) {
	s := "only three words"
	err := timbre.PluralizeNth(&s, 3)
	if err == nil {
		t.Log("Failed on intended index error (index too high)")
		t.Fail()
	}
	err = timbre.PluralizeNth(&s, -1)
	if err == nil {
		t.Log("Failed on intended index error (negative index")
		t.Fail()
	}
}

// tests for snake_case and camelCase manipulation
func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"already_snake", "already_snake"},
		{"A", "a"},
		{"AA", "aa"},
		{"AaAa", "aa_aa"},
		{"HTTPRequest", "http_request"},
		{"BobLoblaw", "bob_loblaw"},
		{"bobLobLaw", "bob_lob_law"},
		{"Id0Value", "id0_value"},
		{"ID0Value", "id0_value"},
		{"speci0u$Intent1ons", "speci0u$_intent1ons"},
	}
	for _, test := range tests {
		have := timbre.ToSnakeCase(test.input)
			if have != test.want {
				t.Log("Failed on ", test.input, "which is", have, "and should be", test.want)
				t.Fail()
			}
	}
}

func TestToCamelCaseLowerFirst(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"currentlyCamel", "currentlyCamel"},
		{"currently_snake", "currentlySnake"},
		{"Aa", "aa"},
		{"aa_aa", "aaAa"},
		{"http_request", "httpRequest"},
		{"bob_lob_law", "bobLobLaw"},
		{"id0_value", "id0Value"},
		{"speci0u$_intent1ons", "speci0u$Intent1ons"},
	}
	for _, test := range tests {
		have := timbre.ToCamelCase(test.input, false)
		if have != test.want {
			t.Log("Failed on", test.input, "which is", have, "and should be", test.want)
			t.Fail()
		}
	}
}

func TestToCamelCaseCapFirst(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"currentlyCamel", "CurrentlyCamel"},
		{"currently_snake", "CurrentlySnake"},
		{"Aa", "Aa"},
		{"aa_aa", "AaAa"},
		{"http_request", "HttpRequest"},
		{"bob_lob_law", "BobLobLaw"},
		{"id0_value", "Id0Value"},
		{"speci0u$_intent1ons", "Speci0u$Intent1ons"},
	}
	for _, test := range tests {
		have := timbre.ToCamelCase(test.input, true)
		if have != test.want {
			t.Log("Failed on", test.input, "which is", have, "and should be", test.want)
			t.Fail()
		}
	}
}