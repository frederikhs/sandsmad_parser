package main

import (
	"encoding/json"
	"github.com/frederikhs/sandsmad_parser"
	"log"
	"net/http"
	"time"
)

func main() {
	var dishes []sandsmad_parser.DishOfTheDay
	go func() {
		for {
			log.Println("updating menu")
			d, err := sandsmad_parser.FetchThisWeeksDishes()
			if err != nil {
				log.Fatal(err)
			}

			dishes = d
			log.Println("done, waiting 24 hours")
			time.Sleep(time.Hour * 24)
		}
	}()

	handler := func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(dishes)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(b)
		if err != nil {
			log.Fatal(err)
		}
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handler))
	log.Println("listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
