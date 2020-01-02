package main

import (
	"encoding/json"
	"errors"
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

	bs, err := toJSON(price)
	if err != nil {
		log.Panicln("Cannot Marshal Price ->", err)
	}

	fmt.Println(string(bs))
}

// toJSON need to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return nil, errors.New("error while marshaling JSON: " + err.Error())
	}

	return bs, nil
}
