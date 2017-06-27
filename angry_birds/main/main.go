package main

import (
	"fmt"

	"github.com/edbond/angrygo/angry_birds"
)

func run() {
	board := angry_birds.Board{
		Board: [][]angry_birds.Pig{
			angry_birds.Pigs("HSP  "),
			angry_birds.Pigs(" AHBP"),
			angry_birds.Pigs(" SRPS"),
			angry_birds.Pigs("HSAPP"),
			angry_birds.Pigs("B HAS"),
		},
	}

	figures := []angry_birds.Figure{
		{
			"XXX",
			"X.X",
		},
		{
			"YY.",
			".YY",
			".Y.",
		},
		{
			".Z.",
			".Z.",
			"ZZZ",
		},
		{
			"VVV",
			".VV",
		},
	}

	// A list of board pieces should left uncovered
	target := map[string]int{
		"H": 1,
		//"A": 2,
		"P": 2,
	}

	results := angry_birds.Solutions(board, figures, target)
	if len(results) < 1 {
		fmt.Errorf("No results found")
	}
}

func main() {
	// start a simple CPU profile and register
	// a defer to Stop (flush) the profiling data.
	//defer profile.Start().Stop()

	run()
}
