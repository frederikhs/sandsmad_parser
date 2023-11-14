package sandsmad_parser

import (
	"testing"
)

func TestTrimLines(t *testing.T) {
	linesToTrim := []string{
		" i have some spaces at my ends ",
		"",
		"i have no spaces at my ends",
	}

	linesToExpect := []string{
		"i have some spaces at my ends",
		"i have no spaces at my ends",
	}

	trimmedLines := trimLines(linesToTrim)
	if len(trimmedLines) != len(linesToExpect) {
		t.Fatalf("expected %d lines, got %d", len(linesToExpect), len(trimmedLines))
	}

	for i := 0; i < len(linesToExpect); i++ {
		if trimmedLines[i] != linesToExpect[i] {
			t.Fatalf("expected \"%s\" to equal \"%s\"", trimmedLines[i], linesToExpect[i])
		}
	}
}

func TestShortDescription(t *testing.T) {
	testCases := []struct {
		Before string
		After  string
	}{
		{
			Before: "Inderfilet med flødeost og tomater samt basilikums olie",
			After:  "Inderfilet med flødeost og tomater",
		},
		{
			Before: "Kalkun deller med Bacon og frisk basilikum løg hvidløg",
			After:  "Kalkun deller med Bacon",
		},
		{
			Before: "Kalkun deller",
			After:  "Kalkun deller",
		},
	}

	for _, testCase := range testCases {
		shorted := shortDescription(testCase.Before)
		if shorted != testCase.After {
			t.Fatalf("expected \"%s\", got \"%s\"", testCase.After, shorted)
		}
	}
}
