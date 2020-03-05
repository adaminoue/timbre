# Timbre Inflector

__WORK IN PROGRESS: Not yet ready for use.__

This library is a project inspired by the [Doctrine PHP inflector](https://github.com/doctrine/inflector) to improve string manipulation utility in Golang.

This project strives for stability and robustness so that it can be used at any scale for the manipulation of arbitrary text.

Regular expressions are used throughout the project. Users are encouraged to check the unit test cases and the word list test to understand function behaviour.

![Image of XKCD #208](https://imgs.xkcd.com/comics/regular_expressions.png)

### Features
* Convert between camelCase and snake_case
* Pluralize and singularize individual words and nth words within strings
* Accurate and tested on the 2500 most common English language nouns (contained in `word_list.csv`)

### Quirks to Know
* Manipulates alphanumeric characters only; other characters are typically ignored (see unit test cases for examples)
* Output is always lowercase

### Testing
* Standard `go test` suite offers complete coverage.
* Test suite can be used to observe results of numerous cases
