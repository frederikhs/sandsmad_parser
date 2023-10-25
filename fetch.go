package sandsmad_parser

import (
	"io"
	"net/http"
)

const URL = "https://sandsmad.dk/"

// fetchBody fetches the remote content using the default http client http.DefaultClient
func fetchBody() ([]byte, error) {
	return fetchBodyWithClient(http.DefaultClient)
}

// fetchBodyWithClient fetches the remote content using the provided http.Client
func fetchBodyWithClient(client *http.Client) ([]byte, error) {
	resp, err := client.Get(URL)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

// FetchThisWeeksDishesWithFetcher fetching, parses and maps this weeks dishes using the provided fetcher
func FetchThisWeeksDishesWithFetcher(fetcher func() ([]byte, error)) ([]DishOfTheDay, error) {
	b, err := fetcher()
	if err != nil {
		return nil, err
	}

	page, err := Parse(b)
	if err != nil {
		return nil, err
	}

	var dishes []DishOfTheDay
	for _, dm := range DayRegexes {
		dish, err := dm.Match(page)
		if err != nil {
			return nil, err
		}

		dishes = append(dishes, *dish)
	}

	return dishes, err
}

// FetchThisWeeksDishes fetches this weeks dishes using the default fetcher
func FetchThisWeeksDishes() ([]DishOfTheDay, error) {
	return FetchThisWeeksDishesWithFetcher(fetchBody)
}

// FetchThisWeeksDishesWithClient fetches this weeks dishes using the provided client in the fetcher
func FetchThisWeeksDishesWithClient(client *http.Client) ([]DishOfTheDay, error) {
	return FetchThisWeeksDishesWithFetcher(func() ([]byte, error) {
		return fetchBodyWithClient(client)
	})
}
