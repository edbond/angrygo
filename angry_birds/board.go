package angry_birds

import (
	"fmt"
	"unicode/utf8"
)

type Pig rune

type Board [][]Pig

func pigs(s string) []Pig {
	var p = make([]Pig, utf8.RuneCountInString(s))
	for r := range []rune(s) {
		p = append(p, Pig(r))
	}
	return p
}

var board = Board{
	pigs("ABC"),
	pigs("BCA"),
	pigs("ABA"),
}

type Figure []string

func (figure Figure) String() string {
	output := "\n"
	for _, s := range figure {
		output += fmt.Sprintln(s)
	}
	return output
}

// RotateCW rotates figure clockwise and returns new figure
func (figure Figure) RotateCW() Figure {
	m := len(figure)
	n := len(figure[0])

	//fmt.Println("Figure\n", figure)

	rows := []string{}
	for i := 0; i < n; i++ {
		var row string
		for j := 0; j < m; j++ {
			c := string(figure[m-j-1][i])
			//fmt.Println("Char=", c)
			//fmt.Println("i=", i, "j=", j, "row=", row)
			row += string(c)
		}
		rows = append(rows, row)
	}

	return Figure(rows)
}

func (figure Figure) equal(other Figure) bool {
	if len(figure) != len(other) {
		return false
	}

	for i, s := range figure {
		if s != other[i] {
			return false
		}
	}

	return true
}

func (figure Figure) notEqual(other Figure) bool {
	return !figure.equal(other)
}
