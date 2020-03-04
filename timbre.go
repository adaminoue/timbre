package timbre

import (
	"regexp"
	"strings"
	"unicode"
)

// -----------------------------------
// ---------- PLURALIZATION ----------
// -----------------------------------

// REGEX
var regexPluralizeIStoES = regexp.MustCompile("([0-9A-Za-z_]*)(is)$")            // e.g. analysis  -> analyses
var regexPluralizeES     = regexp.MustCompile("([0-9A-Za-z_]*)(ch|o|sh|s|x|z)$") // e.g. bench     -> benches
var regexPluralizeIES    = regexp.MustCompile("([0-9A-Za-z_]*[^aeiou])(y)$")     // e.g. city      -> cities
var regexPluralizeING    = regexp.MustCompile("([0-9A-Za-z_]*)(ing)$")           // e.g. running   -> running

// EXPORTED FUNC
func Pluralize(w string) string {
	// if word is all caps, it's probably an abbreviation so make no change
	if strings.ToUpper(w) == w {
		return w
	}

	// then make string lowercase
	w = strings.ToLower(w)

	// if word is key on irregular list, return associated value
	irregulars := irregulars("pluralize")
	v, ok := irregulars[w]
	if ok {
		return v
	}

	// check against each regex
	m := regexPluralizeIStoES.FindStringSubmatch(w)
	if m != nil {
		return m[1] + "es"
	}

	m = regexPluralizeES.FindStringSubmatch(w)
	if m != nil {
		return w + "es"
	}

	m = regexPluralizeIES.FindStringSubmatch(w)
	if m != nil {
		return m[1] + "ies"
	}

	m = regexPluralizeING.FindStringSubmatch(w)
	if m != nil {
		return w
	}

	return w + "s"
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
	for index, value := range s {
		return string(unicode.ToLower(value)) + s[index+1:]
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