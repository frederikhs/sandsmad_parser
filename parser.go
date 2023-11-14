package sandsmad_parser

import (
	"astuart.co/goq"
	"bytes"
	"errors"
	"regexp"
	"strings"
)

// SandsPage is a struct containing the menu goquery for extracting this weeks menu
type SandsPage struct {
	Menu string `goquery:"#main > div > section > div > div > div.this-week > div > div"`
}

// DayRegexes is a slice of DayMatcher for each weekday
var DayRegexes = []DayMatcher{
	{
		NameOfDay: "Mandag",
		Regex:     regexp.MustCompile(`(?i)Mandag\s([\S\s]*)\sTirsdag`),
	},
	{
		NameOfDay: "Tirsdag",
		Regex:     regexp.MustCompile(`(?i)Tirsdag\s([\S\s]*)\sOnsdag`),
	},
	{
		NameOfDay: "Onsdag",
		Regex:     regexp.MustCompile(`(?i)Onsdag\s([\S\s]*)\sTorsdag`),
	},
	{
		NameOfDay: "Torsdag",
		Regex:     regexp.MustCompile(`(?i)Torsdag\s([\S\s]*)\sFredag`),
	},
	{
		NameOfDay: "Fredag",
		Regex:     regexp.MustCompile(`(?i)Fredag\s([\S\s]*)\sMonday`),
	},
}

// Parse takes some content and parses it into a SandsPage extracting the menu using the SandsPage goquery
func Parse(content []byte) (*SandsPage, error) {
	var sp SandsPage

	err := goq.NewDecoder(bytes.NewReader(content)).Decode(&sp)
	if err != nil {
		return nil, err
	}

	// remove html spaces NBSP
	sp.Menu = strings.ReplaceAll(sp.Menu, "\u00a0", " ")

	if len(sp.Menu) == 0 {
		return nil, errors.New("menu was empty after decoding")
	}

	return &sp, nil
}
