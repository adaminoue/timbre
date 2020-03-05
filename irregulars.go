package timbre

// file contains all lists of irregular words, named by function

func irregulars(c string) map[string]string {
	if c == "pluralize" {
		return irregularsPluralize
	}

	return map[string]string{}
}

var irregularsPluralize = map[string]string {
	"addendum": "addenda",
	"advice": "advice",
	"air": "air",
	"aircraft": "aircraft",
	"alumna": "alumnae",
	"alumnus": "alumni",
	"antenna": "antennae",
	"apex": "apexes",
	"appendix": "appendices",
	"auto": "autos",
	"bacillus": "bacilli",
	"bacteria": "bacteria",
	"bacterium": "bacteria",
	"beau": "beaus",
	"beef": "beef",
	"beginning": "beginnings",
	"being": "beings",
	"bison": "bison",
	"bombing": "bombings",
	"building": "buildings",
	"bureau": "bureaus",
	"cactus": "cacti",
	"calf": "calves",
	"cash": "cash",
	"casino": "casinos",
	"ceiling": "ceilings",
	"child": "children",
	"clothes": "clothes",
	"codex": "codices",
	"corpus": "corpora",
	"criteria": "criteria",
	"criterion": "criteria",
	"curriculum": "curricula",
	"data": "data",
	"datum": "data",
	"diabetes": "diabetes",
	"debris": "debris",
	"deer": "deer",
	"die": "dice",
	"drawing": "drawings",
	"dwarf": "dwarves",
	"earnings": "earnings",
	"economics": "economics",
	"ego": "egos",
	"electronics": "electronics",
	"elf": "elves",
	"erratum": "errata",
	"ethics": "ethics",
	"evening": "evenings",
	"fame": "fame",
	"feeling": "feelings",
	"fez": "fezzes",
	"finding": "findings",
	"fish": "fish",
	"focus": "focuses",
	"foot": "feet",
	"fungus": "fungi",
	"gathering": "gatherings",
	"genus": "genera",
	"german": "german",
	"goose": "geese",
	"graffito": "graffiti",
	"grandchild": "grandchildren",
	"guilt": "guilt",
	"half": "halves",
	"happiness": "happiness",
	"hardware": "hardware",
	"headquarters": "headquarters",
	"health": "health",
	"heat": "heat",
	"help": "help",
	"hoof": "hooves",
	"homework": "homework",
	"human": "humans",
	"hunger": "hunger",
	"ice": "ice",
	"index": "indices",
	"japanese": "japanese",
	"jeans": "jeans",
	"jewelry": "jewelry",
	"king": "kings",
	"landing": "landings",
	"larva": "larvae",
	"loaf": "loaves",
	"locus": "loci",
	"louse": "lice",
	"mathematics": "mathematics",
	"matrix": "matrices",
	"meaning": "meanings",
	"media": "media",
	"medium": "media",
	"minutia": "minutiae",
	"moose": "moose",
	"morning": "mornings",
	"mouse": "mice",
	"music": "music",
	"nebula": "nebulae",
	"news": "news",
	"north": "north",
	"northeast": "northeast",
	"northwest": "northwest",
	"nucleus": "nuclei",
	"odds": "odds",
	"offering": "offerings",
	"opening": "openings",
	"opus": "opera",
	"ovum": "ova",
	"ox": "oxen",
	"painting": "paintings",
	"phenomenon": "phenomena",
	"photo": "photos",
	"phylum": "phyla",
	"physics": "physics",
	"piano": "pianos",
	"poetry": "poetry",
	"politics": "politics",
	"portfolio": "portfolios",
	"pro": "pros",
	"quiz": "quizzes",
	"radio": "radios",
	"radius": "radii",
	"rating": "ratings",
	"ratio": "ratios",
	"recording": "recordings",
	"referendum": "referenda",
	"ring": "rings",
	"ruling": "rulings",
	"sales": "sales",
	"salmon": "salmon",
	"scarf": "scarves",
	"scenario": "scenarios",
	"screening": "screenings",
	"self": "selves",
	"series": "series",
	"setting": "settings",
	"sheep": "sheep",
	"shorts": "shorts",
	"shrimp": "shrimp",
	"sibling": "siblings",
	"species": "species",
	"statistics": "statistics",
	"stimulus": "stimuli",
	"stomach": "stomachs",
	"stratum": "strata",
	"string": "strings",
	"studio": "studios",
	"surveillance": "surveillance",
	"swine": "swine",
	"swing": "swings",
	"syllabus": "syllabuses",
	"tennis": "tennis",
	"terrorism": "terrorism",
	"thanks": "thanks",
	"thief": "thieves",
	"thing": "things",
	"timing": "timings",
	"tobacco": "tobaccos",
	"tooth": "teeth",
	"tourism": "tourism",
	"trout": "trout",
	"tuna": "tuna",
	"two": "twos",
	"unemployment": "unemployment",
	"vertebra": "vertebrae",
	"vertex": "vertices",
	"video": "videos",
	"violence": "violence",
	"vita": "vitae",
	"warning": "warnings",
	"wedding": "weddings",
	"wharf": "wharves",
	"wife": "wives",
	"wildlife": "wildlife",
	"wing": "wings",
	"wolf": "wolves",
	"works": "works",
	"writing": "writings",
}