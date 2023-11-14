package sandsmad_parser

import (
	"strings"
	"unicode"
)

// trimLines trims spaces for each line and removes empty lines
func trimLines(lines []string) []string {
	out := []string{}
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			out = append(out, trimmedLine)
		}
	}

	return out
}

// shortDescription create a short description based on the line provided using conjunction words and breaks
func shortDescription(line string) string {
	var short string
	if strings.Contains(line, "og") && strings.Contains(line, "samt") {
		short = strings.Split(line, "samt")[0]
	} else if strings.Contains(line, "og") {
		short = strings.Split(line, "og")[0]
	} else if strings.Contains(line, "(") && strings.Contains(line, ")") {
		short = strings.Split(line, "(")[0]
	} else {
		short = line
	}

	short = strings.TrimSpace(short)

	// capitalize first letter
	shortRunes := []rune(short)
	firstLetter := shortRunes[0]
	shortRunes[0] = unicode.ToUpper(firstLetter)
	short = string(shortRunes)

	return short
}
