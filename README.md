# Timbre Inflector

This library is a project inspired by the [Doctrine PHP inflector](https://github.com/doctrine/inflector) to improve string manipulation utility in Golang.

This project strives for stability and robustness so that it can be used at any scale, including large business operations, for the manipulation of arbitrary text.

![Image of XKCD #208](https://imgs.xkcd.com/comics/regular_expressions.png)

### Features
* Convert between camelCase and snake_case
* Pluralize and singularize individual words and strings

### Limitations
* Manipulates alphanumeric characters only; other characters are typically ignored (see unit test cases for examples)

### Testing

* Standard `go test` suite offers complete coverage.
* Test suite can be used to observe results of numerous cases