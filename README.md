# Sandsmad Parser

[![Release](https://img.shields.io/github/v/release/frederikhs/sandsmad_parser.svg)](https://github.com/frederikhs/sandsmad_parser/releases/latest)
[![Quality](https://goreportcard.com/badge/github.com/frederikhs/sandsmad_parser)](https://goreportcard.com/report/github.com/frederikhs/sandsmad_parser)
[![Test](https://github.com/frederikhs/sandsmad_parser/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/frederikhs/sandsmad_parser/actions/workflows/test.yml)
[![Go Coverage](https://github.com/frederikhs/sandsmad_parser/wiki/coverage.svg)](https://raw.githack.com/wiki/frederikhs/sandsmad_parser/coverage.html)
[![License](https://img.shields.io/github/license/frederikhs/sandsmad_parser)](LICENSE)
[![GoDoc](https://godoc.org/github.com/frederikhs/sandsmad_parser?status.svg)](https://godoc.org/github.com/frederikhs/sandsmad_parser)

This is a golang package for accessing this weeks dishes at [sandsmad.dk](https://sandsmad.dk) in a structured format.

### Example usage

[examples/cli-print/main.go](examples/cli-print/main.go)
```go
package main

import (
	"fmt"
	"github.com/frederikhs/sandsmad_parser"
	"log"
)

func main() {
	dishes, err := sandsmad_parser.FetchThisWeeksDishes()
	if err != nil {
		log.Fatal(err)
	}

	for i, dish := range dishes {
		fmt.Println(dish.NameOfDay)
		for _, line := range dish.Lines {
			fmt.Println(line)
		}

		if i < len(dishes)-1 {
			fmt.Println("---")
		}
	}
}
```

Results in the following ouput
```text
$ go run examples/cli-print/main.go
Mandag
Inderfilet med flødeost og tomater samt basilikums olie (3 stk. pr. person indeholder laktose )
Pasta salat
Vegetar Svampebøf med halumi og rødbeder
---
Tirdsag
Bouf Stroganof med cocktail pølser og rodfrugter svampe samt perleløg indeholder laktose og svinekød
Ris
Vegetar Gryderet med kidney bønner rodfrugter perleløg og svampe
---
Onsdag
Kalkun deller med Bacon og frisk basilikum løg hvidløg (3 stk pr person) indeholder gluten
Kartoffelsalat med karry og forårsløg samt drænet yoghurt
Tærte med porre og kartoffel peber mælk og æg
---
Torsdag
Cremet kokos tomat kyllingegryde med pære og citron græs samt løg og urter mild karry
Ris med Basilikum olie
Vegetar Cremet kokos tomat gryde med limabønne pære og citron græs samt løg og urter mild karry
---
Fredag
Mexicansk spinat tortilla
Chili con carne med tomat og mager hakket oksekød chili bønner løg hvidløg
Vegetar vegansk chili sin carne med tomat og chili bønner løg hvidløg
Ingefærshot med presset æble og appelsin saft
Så i kan gå sunde på weekend
```

See also [examples/json-api-webserver/main.go](examples/json-api-webserver/main.go)
