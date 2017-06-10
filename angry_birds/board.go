package angry_birds

import (
	"fmt"
	"strings"
)

type Pig rune

type Figure []string

// FigureOnBoard Figure positioned on board
// x - horizontal left to right
// y - vertical top down
type FigureOnBoard struct {
	figure Figure
	x      int
	y      int
}

type Board struct {
	board   [][]Pig
	figures []FigureOnBoard
}

// String function prints Board with figures
func (b Board) String() string {
	output := ""

	// Print background
	for _, row := range b.board {
		rs := make([]string, 0)
		for _, x := range row {
			rs = append(rs, string(x))
		}
		s := strings.Join(rs, "")
		output += s + "\n"
	}

	// TODO: Print figures

	return output
}

func pigs(s string) []Pig {
	var p = make([]Pig, 0)
	for _, r := range []rune(s) {
		p = append(p, Pig(r))
	}
	return p
}

var exampleBoard = Board{
	[][]Pig{
		pigs("ABC"),
		pigs("BCA"),
		pigs("ABA"),
	},
	[]FigureOnBoard{},
}

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

	rows := []string{}
	for i := 0; i < n; i++ {
		var row string
		for j := 0; j < m; j++ {
			c := string(figure[m-j-1][i])
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

// positions returns slice of possible positions of figure on board
func Positions(figure Figure, board Board) []FigureOnBoard {
	p := make([]FigureOnBoard, 0)
	bw, bh := len(board.board[0]), len(board.board)
	fw, fh := len(figure[0]), len(figure)

	fmt.Println("Board ", bw, " x ", bh)
	fmt.Println("Figure ", fw, " x ", fh)

	for x := 0; x <= bw-fw; x++ {
		for y := 0; y <= bh-fh; y++ {
			p = append(p, FigureOnBoard{figure, x, y})
		}
	}

	return p
}
