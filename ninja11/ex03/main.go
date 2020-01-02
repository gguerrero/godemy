package main

import (
	"fmt"
	"log"
	"time"
)

// Price struct
type Price struct {
	Value    float64   `json:"value"`
	Datetime time.Time `json:"datetime"`
	Source   string    `json:"source"`
}

func (p Price) Error() string {
	return fmt.Sprintf("Something wrong with the price! - price: %.2f - datetime: %s", p.Value, p.Datetime)
}

func main() {
	price := Price{
		Value:    8.453,
		Datetime: time.Now(),
		Source:   "ICE",
	}

	thoughError(price)
}

func thoughError(err error) {
	fmt.Println("---------------------")
	fmt.Println(err)
	fmt.Println(err.(Price).Value)
	fmt.Println(err.(Price).Datetime)
	fmt.Println(err.(Price).Source)
	fmt.Println("---------------------")

	log.Panicln(err)
}
