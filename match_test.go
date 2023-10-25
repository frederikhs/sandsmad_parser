package sandsmad_parser

import (
	"regexp"
	"testing"
)

func TestDayMatcher_Match(t *testing.T) {
	dm := DayMatcher{
		NameOfDay: "Onsdag",
		Regex:     regexp.MustCompile(`Onsdag\s([\S\s]*)\sTorsdag`),
	}

	sp, err := Parse(ExampleFile)
	if err != nil {
		t.Fatalf("got error while parsing menu: %v", err)
	}

	match, err := dm.Match(sp)
	if err != nil {
		t.Fatalf("got error while matching menu: %v", err)
	}

	if match.NameOfDay != dm.NameOfDay {
		t.Fatalf("expected \"%s\" got \"%s\"", dm.NameOfDay, match.NameOfDay)
	}
}

func TestDayMatcher_Match2(t *testing.T) {
	dm := DayMatcher{
		NameOfDay: "IngenDag",
		Regex:     regexp.MustCompile(`IngenDag\s([\S\s]*)\sEnAndenIngenDag`),
	}

	sp, err := Parse(ExampleFile)
	if err != nil {
		t.Fatalf("got error while parsing menu: %v", err)
	}

	_, err = dm.Match(sp)
	if err == nil {
		t.Fatalf("expected match of \"%s\" to fail, but it did not", dm.NameOfDay)
	}
}
