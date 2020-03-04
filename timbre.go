package timbre

import (
	"regexp"
	"strings"
	"unicode"
)

// TODO: complete this list and implement routines using this function to singularize and pluralize
// list of all words not to treat ordinarily
var irregulars = map[string]string {
	"addendum": "addenda",
	"aircraft": "aircraft",
	"alumna": "alumnae",
	"alumnus": "alumni",
	"antenna": "antennae",
	"apex": "apexes",
	"appendix": "appendices",
	"bacillus": "bacilli",
	"bacterium": "bacteria",
	"beau": "beaus",
	"bison": "bison",
	"bureau": "bureaus",
	"cactus": "cacti",
	"calf": "calves",
	"child": "children",
	"codex": "codices",
	"corpus": "corpora",
	"curriculum": "curricula",
	"datum": "data",
	"deer": "deer",
	"die": "dice",
	"dwarf": "dwarves",
	"elf": "elves",
	"erratum": "errata",
	"fey": "fey",
	"fez": "fezzes",
	"fish": "fish",
	"foot": "feet",
	"fungus": "fungi",
	"gas": "gasses",
	"genus": "genera",
	"goose": "geese",
	"graffito": "graffiti",
	"half": "halves",
	"hoof": "hooves",
	"index": "indices",
	"larva": "larvae",
	"loaf": "loaves",
	"locus": "loci",
	"louse": "lice",
	"man": "men",
	"matrix": "matrices",
	"minutia": "minutiae",
	"moose": "moose",
	"mouse": "mice",
	"nebula": "nebulae",
	"nucleus": "nuclei",
	"offspring": "offspring",
	"opus": "opera",
	"ovum": "ova",
	"ox": "oxen",
	"person": "people",
	"phylum": "phyla",
	"quiz": "quizzes",
	"radius": "radii",
	"referendum": "referenda",
	"salmon": "salmon",
	"scarf": "scarves",
	"self": "selves",
	"series": "series",
	"sheep": "sheep",
	"shrimp": "shrimp",
	"species": "species",
	"stimulus": "stimuli",
	"stratum": "strata",
	"swine": "swine",
	"syllabus": "syllabuses",
	"thief": "thieves",
	"tooth": "teeth",
	"trout": "trout",
	"tuna": "tuna",
	"vertebra": "vertebrae",
	"vertex": "vertices",
	"vita": "vitae",
	"vortex": "vortices",
	"wharf": "wharves",
	"wife": "wives",
	"wolf": "wolves",
	"woman": "women",
}

// -----------------------------------
// ---------- PLURALIZATION ----------
// -----------------------------------

// REGEX
var matchForEndingES     = regexp.MustCompile("\b[0-9A-Za-z_]*(ch|sh|o|[sxz])\b") // e.g. bench     -> benches
var matchForEndingIES    = regexp.MustCompile("\b[0-9A-Za-z_]*[^aeiou](y)\b")     // e.g. city      -> cities
var matchForEndingIStoES = regexp.MustCompile("\b[0-9A-Za-z_]*(is)\b")            // e.g. analysis  -> analyses
var matchForEndingA      = regexp.MustCompile("\b[0-9A-Za-z_]*(on|um)\b")         // e.g. criterion -> criteria

// HELPER FUNCTIONS
func 

func Pluralize(w string) string {
	// match against each regex in sequence and make changes accordingly

	return ""
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