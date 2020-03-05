package test

import (
	"encoding/csv"
	"github.com/adaminoue/timbre"
	"io"
	"os"
	"strings"
	"testing"
)

// --------------------------------------------
// -------------- WORD LIST TEST --------------
// - (for manual inspection; not a unit test) -
// --------------------------------------------

func TestWordListPluralize(t *testing.T) {
	csvFile, err := os.Open("./word_list.csv")
	if err != nil {
		t.Error("word list could not be opened")
	}

	r := csv.NewReader(csvFile)

	var bigList []string

	for {
		// read records individually from CSV
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err.Error())
		}

		bigList = append(bigList, strings.TrimSpace(rec[0]))
	}

	for _, s := range bigList {
		p := timbre.Pluralize(s)
		println(s, "->", p)
	}
}