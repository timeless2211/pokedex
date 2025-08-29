package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type locationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []any  `json:"results"`
}

var offset = 0

func fetchLocationAreas(config *config) (locationAreas, error) {
	if config.Next {
		offset += 1
	}
	if config.Previous {
		offset -= 1
		if offset < 0 {
			offset = 0
		}
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?limit=20&offset=%d", offset)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var locationAreas locationAreas

	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		log.Fatal(err)
	}

	return locationAreas, nil
}
