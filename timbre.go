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
	"analysis": "analyses",
	"antenna": "antennae",
	"antithesis": "antitheses",
	"apex": "apexes",
	"appendix": "appendices",
	"axis": "axes",
	"bacillus": "bacilli",
	"bacterium": "bacteria",
	"basis": "bases",
	"beau": "beaus",
	"bison": "bison",
	"bureau": "bureaus",
	"cactus": "cacti",
	"calf": "calves",
	"child": "children",
	"codex": "codices",
	"corpus": "corpora",
	"crisis": "crises",
	"criterion": "criteria",
	"curriculum": "curricula",
	"datum": "data",
	"deer": "deer",
	"diagnosis": "diagnoses",
	"die": "dice",
	"dwarf": "dwarves",
	"elf": "elves",
	"ellipsis": "ellipses",
	"erratum": "errata",
	"fez": "fezzes",
	"fish": "fish",
	"focus": "focuses",
	"foot": "feet",
	"fungus": "fungi",
	"genus": "genera",
	"goose": "geese",
	"graffito": "graffiti",
	"half": "halves",
	"hoof": "hooves",
	"hypothesis": "hypotheses",
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
	"oasis": "oases",
	"offspring": "offspring",
	"opus": "opera",
	"ovum": "ova",
	"ox": "oxen",
	"parenthesis": "parentheses",
	"phenomenon": "phenomena",
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
	"synopsis": "synopses",
	"thesis": "theses",
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

// Regex matchers to enable functions that follow

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

var matchSnake = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

// Naming convention manipulation functions
func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func lowerFirstChar(s string) string {
	for index, value := range s {
		return string(unicode.ToLower(value)) + s[index+1:]
	}
	return ""
}

func toCamelCase(str string, capFirstLetter bool) string {
	str = matchSnake.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s,"_","",-1))
	})

	if !capFirstLetter {
		return lowerFirstChar(str)
	}

	return str
}