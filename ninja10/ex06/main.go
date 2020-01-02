package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"math"
	"math/rand"
)

type Price struct{
	Value    float64		`json:"value"`
	Datetime time.Time	`json:"datetime"`
	Source   string			`json:"source"`
	Index 	 uint64	    `json:"index"`
}

func main() {
	c := make(chan Price)
	go generateRealTimePrices(c, 25)

	processRealTimePrices(c)
}

func generateRealTimePrices(pricesSenderChannel chan<- Price, totalPrices uint64){
	var i uint64
	for i = 1; i <= totalPrices; i++ {
		price := Price{
			Value: roundFloat(rand.Float64() * 20),
			Datetime: time.Now(),
			Source: "ICE",
			Index: i,
		}

		pricesSenderChannel <- price
		randSleep()
	}
	close(pricesSenderChannel)
}


func processRealTimePrices(pricesReceiverChannel <-chan Price) {
	for price := range pricesReceiverChannel {
		b, err := json.Marshal(price)
		if err != nil {
			fmt.Fprint(os.Stderr, "Cannot Marshal Price!")
			fmt.Fprint(os.Stdout, "{}\n")
		} else {
			fmt.Fprint(os.Stdout, string(b), "\n")
		}
	}
}

func randSleep() {
	r := rand.Intn(3000)
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func roundFloat(v float64) float64 { return math.Round(v*100)/100 }
