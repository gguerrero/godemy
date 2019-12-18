package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`James`: []string{`red`, `blue`, `orange`},
		`Will`:  []string{`red`, `green`, `yellow`},
	}

	m[`Sam`] = []string{`blue`, `black`}

	delete(m, "James")
	delete(m, "Other")

	if v, ok := m["James"]; ok {
		fmt.Println(`Here is the value for "James" ->`, v)
	}

	if v, ok := m["Will"]; ok {
		fmt.Println(`Here is the value for "Will" ->`, v)
	}

	if v, ok := m["Sam"]; ok {
		fmt.Println(`Here is the value for "Sam" ->`, v)
	}
}
