package sandsmad_parser

import (
	"regexp"
	"testing"
)

func TestDayMatcher_Match(t *testing.T) {
	dm := DayMatcher{
		NameOfDay: "Onsdag",
		Regex:     regexp.MustCompile(`(?i)Onsdag\s([\S\s]*)\sTorsdag`),
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
		Regex:     regexp.MustCompile(`(?i)IngenDag\s([\S\s]*)\sEnAndenIngenDag`),
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

func TestDayMatcherTestCaseInsensitiveMatch(t *testing.T) {
	dm := DayMatcher{
		NameOfDay: "Torsdag",
		Regex:     regexp.MustCompile(`(?i)Torsdag\s([\S\s]*)\sFredag`),
	}

	sp := &SandsPage{Menu: `
Onsdag
Kalkun deller med Bacon og frisk basilikum løg hvidløg (3 stk pr person) indeholder gluten
Kartoffelsalat med karry og forårsløg samt drænet yoghurt
Tærte med porre og kartoffel peber mælk og æg

torsdag
Cremet kokos tomat kyllingegryde med pære og citron græs samt løg og urter mild karry
Ris med Basilikum olie
Vegetar Cremet kokos tomat gryde med limabønne pære og citron græs samt løg og urter mild karry

fredag
Mexicansk spinat tortilla
Chili con carne med tomat og mager hakket oksekød chili bønner løg hvidløg
Vegetar vegansk chili sin carne med tomat og chili bønner løg hvidløg
	`}

	match, err := dm.Match(sp)
	if err != nil {
		t.Fatalf("got error while matching menu: %v", err)
	}

	if match.NameOfDay != dm.NameOfDay {
		t.Fatalf("expected \"%s\" got \"%s\"", dm.NameOfDay, match.NameOfDay)
	}
}
