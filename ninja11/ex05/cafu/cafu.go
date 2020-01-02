// package cafu have all the functions set for the CAFU serivice
package cafu

import (
	"math/rand"
	"time"
)

// GetPrice function returns the CAFU fuel price
func GetPrice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
