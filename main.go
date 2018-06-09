package main

import (
	"fmt"

	"github.com/edbond/angrygo/angrybirds"
)

func run() error {
	board := angrybirds.Board{
		Board: [][]angrybirds.Pig{
			angrybirds.Pigs("HSP  "),
			angrybirds.Pigs(" AHBP"),
			angrybirds.Pigs(" SRPS"),
			angrybirds.Pigs("HSAPP"),
			angrybirds.Pigs("B HAS"),
		},
	}

	figures := []angrybirds.Figure{
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
		"A": 2,
		"P": 2,
	}

	results := angrybirds.Solutions(board, figures, target)
	if len(results) < 1 {
		return fmt.Errorf("No results found")
	}
	return nil
}

func main() {
	// start a simple CPU profile and register
	// a defer to Stop (flush) the profiling data.
	// defer profile.Start().Stop()

	err := run()
	if err != nil {
		fmt.Println(err)
	}
}
