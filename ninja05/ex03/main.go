package main

import (
	"fmt"
)


type Vehicle struct{
	Doors int
	Color string
}

type Sedan struct{
	Vehicle
	Luxury bool
}

type Truck struct{
	Vehicle
	FourWheel bool
}

func main() {
	sedan := Sedan{
		Vehicle: Vehicle{
			Doors: 4,
			Color: "red",
		},
		Luxury: true,
	}

	truck := Truck{
		Vehicle: Vehicle{
			Doors: 2,
			Color: "gray",
		},
		FourWheel: false,
	}

	fmt.Printf("%+v\n", sedan)
	fmt.Printf("%+v\n", truck)

	fmt.Printf("%v\t%v\n", sedan.Doors, sedan.Luxury)
	fmt.Printf("%v\t%v\n", truck.Doors, truck.Color)
}
