package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Price struct {
	Value    float64   `json:"value"`
	Datetime time.Time `json:"datetime"`
	Source   string    `json:"source"`
}

func main() {
	price := Price{
		Value:    8.453,
		Datetime: time.Now(),
		Source:   "ICE",
	}

	bs, err := json.Marshal(price)
	if err != nil {
		log.Panicln("Cannot Marshal Price ->", err)
	}

	fmt.Println(string(bs))
}
