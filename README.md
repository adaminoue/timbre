# Timbre Inflector

__WORK IN PROGRESS:__
* Only exported functions are `Pluralize`, `ToSnakeCase`, `ToCamelCase`, `Irregulars`
* Next upcoming: `Singularize`

This library is a project inspired by the [Doctrine PHP inflector](https://github.com/doctrine/inflector) to improve string manipulation utility in Golang.

This project strives for stability and robustness so that it can be used at any scale for the manipulation of arbitrary text.

Regular expressions are used throughout the project. Users are encouraged to check the unit test cases and the word list test to understand function behaviour.

![Image of XKCD #208](https://imgs.xkcd.com/comics/regular_expressions.png)

### Features
* Pluralize individual words and nth words within strings
    * Accurate and tested on the ~2500 most common English language nouns (contained in `word_list.csv` and viewable via `word_list_test.go`)
* Convert between camelCase and snake_case
    * Simple implementations; easy to read
* Runs in linear time relative to size of input (see [here](https://golang.org/pkg/regexp/) and [here](https://swtch.com/~rsc/regexp/regexp1.html))

### Function Behaviour
* `Pluralize()` – Pluralize a single word (in place)
    * All-caps input is assumed to be an abbreviation (e.g. FBI, USA) and is left unchanged
    * Scoped to manipulate words containing alphanumeric characters only
    * Leading and trailing non-alphanumeric characters are preserved (e.g. `^American` becomes `^Americans`)
    * Capitalization is preserved (e.g. `CalTrain` becomes `CalTrains`)
        * Exception: The overwritten portion of a word is always lowercase (e.g. `citY` becomes `cities`)
        * Exception: For words on the [irregulars list](/irregulars.go), mid-word capitalization is not preserved (e.g. `PainTing` becomes `Paintings`)
* `PluralizeNth()` – Pluralize the nth word of a string (in place)
    * Trims whitespace at ends of string, then runs `Pluralize` on the nth word of the input
    * Expects words to only be separated by one space only; other splits can cause unintended results
        * You can work around this; underlying split is `strings.Split(s, " ")` where `s` is input
* `Irregulars()` – Return a `map[string]string` of irregulars, with singulars as keys and plurals as values
    * Takes a string argument to match to the correct map; these keys match the related function in lowercase, e.g. `pluralize`

### Testing
* Standard `go test` suite contained in `timbre_test.go`
    * Tests can be examined to observe results of numerous cases
* Use `go test word_list/word_list_test.go` to print more exhaustive list of test words

### Improvements
* If you believe a certain input is not handled correctly, please search the Issues for related discussion. If it has not been previously discussed, feel free to submit a new issue or PR.