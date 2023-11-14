package sandsmad_parser

import (
	"fmt"
	"regexp"
	"strings"
)

// DayMatcher is a struct containing a regular expression for matching the lines that describes a dish on a particular day
type DayMatcher struct {
	NameOfDay string
	Regex     *regexp.Regexp
}

// Match takes a menu and returns a DishOfTheDay parsing the menu using the DayMatcher
func (dm DayMatcher) Match(page *SandsPage) (*DishOfTheDay, error) {
	matches := dm.Regex.FindStringSubmatch(page.Menu)
	if len(matches) != 2 {
		return nil, fmt.Errorf("could not find match for day: \"%s\" with menu: %s", dm.NameOfDay, page.Menu)
	}

	lines := trimLines(strings.Split(matches[1], "\n"))
	mainDish := lines[0]

	return &DishOfTheDay{
		NameOfDay:        dm.NameOfDay,
		Lines:            lines,
		ShortDescription: shortDescription(mainDish),
	}, nil
}
