package timbre

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// -----------------------------------
// ---------- PLURALIZATION ----------
// -----------------------------------

// REGEX
var regexPluralizeIStoES = regexp.MustCompile("([0-9A-Za-z]*)(is)$")            // e.g. analysis -> analyses
var regexPluralizeES     = regexp.MustCompile("([0-9A-Za-z]*)(ch|o|s|sh|x|z)$") // e.g. bench    -> benches
var regexPluralizeIES    = regexp.MustCompile("([0-9A-Za-z]*[^aeiou])(y)$")     // e.g. city     -> cities
var regexPluralizeING    = regexp.MustCompile("([0-9A-Za-z]*)(ing)$")           // e.g. running  -> running
var regexPluralizeMEN    = regexp.MustCompile("([0-9A-Za-z]*)(man)$")           // e.g. freshman -> freshmen

var regexPreserveEnds    = regexp.MustCompile("([^0-9A-Za-z]*)([0-9A-Za-z]*)([^0-9A-Za-z]*)")

// HELPER FUNC
// swaps first n characters of string s for contents of stem, where n = len(stem).
func swapStem(s string, stem string) string {
	return string(stem) + s[len(stem):]
}

// EXPORTED FUNC: Pluralize a word (in place).
func Pluralize(s *string) {
	ends := regexPreserveEnds.FindStringSubmatch(*s)
	h, w, t := ends[1], ends[2], ends[3]

	// if word is all caps, it's probably an abbreviation so make no change
	if strings.ToUpper(w) == w {
		return
	}

	// grab copy of word to preserve capitalization, then make word lowercase
	stem := w
	ltr := string(stem[0])
	w = strings.ToLower(w)

	// if word is key on irregular list, return associated value
	irr := Irregulars("pluralize")
	v, ok := irr[w]
	if ok {
		*s = h + swapStem(v, ltr) + t
		return
	}

	// check against each regex
	m := regexPluralizeING.FindStringSubmatch(w)
	if m != nil {
		return
	}

	m = regexPluralizeIStoES.FindStringSubmatch(w)
	if m != nil {
		*s = h + swapStem(m[1] + "es", stem[:len(m[1])]) + t
		return
	}

	m = regexPluralizeES.FindStringSubmatch(w)
	if m != nil {
		*s = h + swapStem(w + "es", stem) + t
		return
	}

	m = regexPluralizeIES.FindStringSubmatch(w)
	if m != nil {
		*s = h + swapStem(m[1] + "ies", stem[:len(m[1])]) + t
		return
	}

	m = regexPluralizeMEN.FindStringSubmatch(w)
	if m != nil {
		*s = h + swapStem(m[1] + "men", stem[:len(m[1])]) + t
		return
	}

	*s = h + swapStem(w + "s", stem) + t
	return
}

// Trims whitespace at both ends of string and pluralizes nth word of string (in place).
func PluralizeNth(s *string, n int) error {
	*s = strings.TrimSpace(*s)
	x := strings.Split(*s, " ")

	if n < 0 || n + 1 > len(x) {
		return errors.New("invalid index")
	}

	Pluralize(&x[n])
	*s = strings.Join(x, " ")

	return nil
}

// -----------------------------------
// -- NAMING CONVENTION MANIPULATION -
// -----------------------------------

// REGEX
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchSnake    = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

// HELPER FUNC
func lowerFirstChar(s string) string {
	for i, v := range s {
		return string(unicode.ToLower(v)) + s[i+1:]
	}
	return ""
}

// EXPORTED FUNC
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToCamelCase(str string, capFirstLetter bool) string {
	str = matchSnake.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s,"_","",-1))
	})

	if !capFirstLetter {
		return lowerFirstChar(str)
	}
	return str
}