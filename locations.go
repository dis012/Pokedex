package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokeMap struct {
	Count    int        `json:"count"`
	NextPage string     `json:"next"`
	PrevPage string     `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Config struct {
	PokeClient client
	NextPage   *string
	PrevPage   *string
}

func (c *client) ListLocations(api *string) (PokeMap, error) {
	url := API_URL
	if api != nil {
		url = *api
	}

	if val, ok := c.cache.Get(url); ok {
		locations := PokeMap{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return PokeMap{}, fmt.Errorf("error unmarshalling cached response: %v", err)
		}
		return locations, nil
	}

	req, err := http.Get(url)
	if err != nil {
		return PokeMap{}, fmt.Errorf("error getting locations: %v", err)
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if req.StatusCode > 299 {
		return PokeMap{}, fmt.Errorf("error getting locations: %s", body)
	}

	if err != nil {
		return PokeMap{}, fmt.Errorf("error reading response body: %v", err)
	}
	location := PokeMap{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return PokeMap{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	c.cache.Add(url, body)
	return location, nil
}
