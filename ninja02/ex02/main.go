package main

import (
	"fmt"
)

func main() {
	a := (1 == 1)
	b := (1 <= 2)
	c := (10 >= 2)
	d := (1 != 2)
	e := (10 < 100)
	f := (1000 > 100)

	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", a, b, c, d, e, f)
}
