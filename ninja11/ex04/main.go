package main

import (
	"errors"
	"fmt"
	"math"
	"log"
)

type sqrt struct {
	lat		float64
	long	float64
}

type sqrtError struct {
	sqrt
	err	error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}


func main() {
	result, err := doSqrt(-10.23)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("The square root is", result)
}

func doSqrt(f float64) (float64, error) {
	if f < 0 {
		err := sqrtError{
			sqrt: sqrt{
				lat: 10.2,
				long: 20.4,
			},
			err: errors.New("no sense to have lat-long"),
		}
		return 0.0, err
	}
	return math.Sqrt(f), nil
}
