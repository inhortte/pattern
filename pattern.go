package pattern

import (
	"fmt"
	"math/rand"
)

type ref struct {
	Notes  []string
	Length int
	Odds   float64
}

var fg1 = ref{
	Notes:  []string{"Ab", "c", "Eb", "g"},
	Length: 12,
	Odds:   0.3,
}

var fg2 = ref{
	Notes:  []string{"Eb", "g", "d", "f"},
	Length: 12,
	Odds:   0.2,
}
var fg3 = ref{
	Notes:  []string{"c", "g", "d", "Eb"},
	Length: 12,
	Odds:   0.15,
}

func Pattern() {
	three := [3]ref{fg1, fg2, fg3}
	for i := 0; i < three[0].Length; i++ {
		for _, r := range three {
			var c string
			if rand.Float64() < r.Odds {
				c = r.Notes[rand.Intn(len(r.Notes))]
			} else {
				c = "-"
			}
			fmt.Printf("%s ", c)
		}
		fmt.Printf("\n")
	}
}
