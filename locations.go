package main

type config struct {
	Count    int        `json:"count"`
	NextPage string     `json:"next"`
	PrevPage string     `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
