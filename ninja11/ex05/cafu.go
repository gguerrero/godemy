package cafu

import (
	"math/rand"
	"time"
)

// Order function creates and order
func GetPrice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
