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
