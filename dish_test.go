package sandsmad_parser

import (
	"testing"
	"time"
)

func TestDishOfTheDay_IsDishOnTheMenuAt(t *testing.T) {
	dotd := givenDishOfTheDay(Wednesday)

	err, isOn := dotd.IsDishOnTheMenuAt(time.Date(2023, 11, 15, 11, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("got error while checking if dish is on the menu at: %v", err)
	}

	if !isOn {
		t.Fatalf("dish should be on the menu, was not")
	}
}

func TestDishOfTheDay_IsDishOnTheMenuAtWeekdayNotDishDay(t *testing.T) {
	dotd := givenDishOfTheDay(Wednesday)

	err, _ := dotd.IsDishOnTheMenuAt(time.Date(2023, 11, 12, 11, 0, 0, 0, time.UTC))
	if err == nil {
		t.Fatalf("expected to fail, but did not")
	}
}
