// HLCALC -- calculate capacitor, resistor, cutoff frequency for high pass and low pass filters
//
// usage: hlcalc c=... r=... f=... (exactly 2 out of 3 required)
//
// svm 31-JUL-2026
//

package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"math"
	"os"
	"strings"
)

func main() {
	var c, r, f float64

	args := os.Args[1:]
	for ac, av := range args {
		switch ac {
		case 0, 1:
			// we might want to implement some sort of
			// sanity checking here some day.
			xx := strings.Split(av, "=")
			switch xx[0] {
			case "c", "C":
				c, _, _ = humanize.ParseSI(xx[1])
			case "r", "R":
				r, _, _ = humanize.ParseSI(xx[1])
			case "f", "F":
				f, _, _ = humanize.ParseSI(xx[1])
			}
		case 2:
			println("bad arg")
			os.Exit(1)
		}
	}

	if f == 0 { // calc f from r, c
		f = 1 / (2 * math.Pi * c * r)

	} else if r == 0 { // calc r from f, c
		r = 1 / (2 * math.Pi * c * f)

	} else if c == 0 { // calc c from f, r
		c = 1 / (2 * math.Pi * r * f)

	} else {
		println("bad arg")
		os.Exit(1)
	}

	prpr("resistance", "R", "Ω", r)
	prpr("capacitance", "C", "F", c)
	prpr("frequency", "F", "Hz", f)
}

func prpr(label string, abbrev string, unit string, val float64) {
	vval := humanize.SIWithDigits(val, 2, unit)
	fmt.Printf("%-12s %s = %s\n", label, abbrev, vval)
}
