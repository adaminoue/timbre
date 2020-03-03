package timbre

import "testing"

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
		have := toSnakeCase(test.input)
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
		have := toCamelCase(test.input, false)
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
		have := toCamelCase(test.input, true)
		if have != test.want {
			t.Log("Failed on", test.input, "which is", have, "and should be", test.want)
			t.Fail()
		}
	}
}