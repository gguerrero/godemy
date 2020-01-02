package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

type Price struct {
	Value    float64   `json:"value"`
	Datetime time.Time `json:"datetime"`
	Source   string    `json:"source"`
	Index    int       `json:"index"`
	Channel  int       `json:"channel"`
}

const totalGoRoutines = 10
const pricesPerGoRoutine = 25

func main() {
	channels := generateRealTimePrices(totalGoRoutines, pricesPerGoRoutine)
	processRealTimePrices(channels...)
}

func generateRealTimePrices(totalGoRoutines int, pricesPerGoRoutine int) []<-chan Price {
	channels := [](<-chan Price){}

	for i := 1; i <= totalGoRoutines; i++ {
		priceChan := make(chan Price)

		go func(channelNum int) {
			for j := 1; j <= pricesPerGoRoutine; j++ {
				price := Price{
					Value:    roundFloat(rand.Float64() * 20),
					Datetime: time.Now(),
					Source:   "ICE",
					Index:    j,
					Channel:  channelNum,
				}

				priceChan <- price
				fmt.Printf("[PRICE SEND TO CHANNEL #%d]\n", channelNum)
				randSleep()
			}
		}(i)

		channels = append(channels, priceChan)
	}

	return channels
}

func processRealTimePrices(pricesChannels ...<-chan Price) {
	for _, pricesChan := range pricesChannels {
		for price := range pricesChan {
			b, err := json.Marshal(price)
			if err != nil {
				fmt.Fprint(os.Stderr, "Cannot Marshal Price!")
				fmt.Fprint(os.Stdout, "{}\n")
			} else {
				fmt.Fprint(os.Stdout, string(b), "\n")
			}
		}
	}
}

func randSleep() {
	r := rand.Intn(3000)
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func roundFloat(v float64) float64 { return math.Round(v*100) / 100 }
