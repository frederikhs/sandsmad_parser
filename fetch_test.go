package sandsmad_parser

import (
	"net/http"
	"testing"
	"time"
)

func TestCanFetchBody(t *testing.T) {
	_, err := fetchBody()
	if err != nil {
		t.Fatalf("got error while matching menu: %v", err)
	}
}

func TestCanParseFetchedBody(t *testing.T) {
	body, err := fetchBody()
	if err != nil {
		t.Fatalf("got error while matching menu: %v", err)
	}

	page, err := Parse(body)
	if err != nil {
		t.Fatalf("got error while parsing menu: %v", err)
	}

	givenIRangeOverDayRegexesForPage(t, page)
}

func TestFetchThisWeeksDishes(t *testing.T) {
	givenIGetDishesBack(t, FetchThisWeeksDishes)
}

func TestFetchThisWeeksDishesWithClient(t *testing.T) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	givenIGetDishesBack(t, func() ([]DishOfTheDay, error) {
		return FetchThisWeeksDishesWithClient(client)
	})
}
