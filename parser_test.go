package sandsmad_parser

import (
	"testing"
)

func TestCanParse(t *testing.T) {
	_, err := Parse(ExampleFile)
	if err != nil {
		t.Fatalf("got error while parsing menu: %v", err)
	}
}

func TestCannotParseBadContent(t *testing.T) {
	_, err := Parse([]byte(""))
	if err == nil {
		t.Fatalf("expected error while parsing empty content")
	}
}

func TestCanParseAllDays(t *testing.T) {
	page, err := Parse(ExampleFile)
	if err != nil {
		t.Fatalf("got error while parsing menu: %v", err)
	}

	givenIRangeOverDayRegexesForPage(t, page)
}
