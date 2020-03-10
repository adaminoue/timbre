package test

import (
	"github.com/adaminoue/timbre"
	"testing"
)

func TestIrregularsFirstLettersMatch(t *testing.T) {
	irr := timbre.Irregulars("pluralize")

	for k, v := range irr {
		if k[0] != v[0] {
			t.Log("Irregulars", k, "and", v, "have different first letters which will break stem swap")
			t.Fail()
		}
	}
}