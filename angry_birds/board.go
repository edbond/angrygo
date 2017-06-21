package angry_birds

import (
	"fmt"
	"reflect"
	"strings"
)

type Pig rune

type Figure []string

var Skip string = "."

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
	output := []string{}

	// Print background
	for _, row := range b.board {
		rs := make([]string, 0)
		for _, x := range row {
			rs = append(rs, string(x))
		}
		s := strings.Join(rs, "")
		output = append(output, s)
	}

	for _, f := range b.figures {
		for j, row := range f.figure {
			for i, x := range row {
				if string(x) != Skip {
					old := output[f.y+j]
					output[f.y+j] = old[:(f.x+i)] + string(x) + old[(f.x+i+1):]
				}
			}
		}
	}

	return strings.Join(output, "\n")
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

func (figure Figure) width() int {
	return len(figure[0])
}

func (figure Figure) height() int {
	return len(figure)
}

func (board Board) width() int {
	return len(board.board[0])
}

func (board Board) height() int {
	return len(board.board)
}

// Positions returns possible positions of figure on board
func Positions(figure Figure, board Board) []FigureOnBoard {
	p := make([]FigureOnBoard, 0)
	bw, bh := board.width(), board.height()
	fw, fh := figure.width(), figure.height()

	//fmt.Println("Board ", bw, " x ", bh)
	//fmt.Println("Figure ", fw, " x ", fh)

	for x := 0; x <= bw-fw; x++ {
		for y := 0; y <= bh-fh; y++ {
			p = append(p, FigureOnBoard{figure, x, y})
		}
	}

	return p
}

// rotations returns all valid rotations of figure f
// excluding duplicates
func rotations(f Figure) []Figure {
	results := []Figure{f}
	var r Figure = f

	for range []int{1, 2, 3} {
		r = r.RotateCW()

		var found bool = false

		// Find if there is already such rotation
		for _, x := range results {
			if x.equal(r) {
				found = true
			}
		}

		if !found {
			results = append(results, r)
		}
	}
	return results
}

func permutations(input [][]int) [][]int {
	m := [][]int{}

	l := len(input)
	i := make([]int, l)

	current_i := 0

	// last index of each row
	ends := make([]int, l)
	for j, x := range input {
		ends[j] = len(x) - 1
	}

	var found bool

	for {
		current := make([]int, l)
		// get current perm
		for j, v := range i {
			current[j] = input[j][v]
		}

		//fmt.Println("Current perm", current)
		//fmt.Println("Indexes", i)

		m = append(m, current)

		// exit if we get to the end
		if reflect.DeepEqual(i, ends) {
			break
		}

		// Make current I +1
		// if result > available in this row
		// get to next and try to increase
		// if success reset all left to 0
		// start over
		// if no success - break
		found = false
		for j := range i {
			if i[j] < ends[j] {
				current_i = j
				found = true
				break
			}
		}

		if !found {
			break
		}

		// increase current index
		// and reset all to the left
		i[current_i] += 1
		for j := 0; j < (current_i - 1); j++ {
			i[j] = 0
		}

		//fmt.Println("Indexes updated", i)
	}

	return m
}

func FigureCombinations(rows [][]Figure) [][]Figure {
	sets := [][]Figure{}

	// get indices
	indices := make([][]int, len(rows))

	for i, r := range rows {
		x := make([]int, 0)
		for j := 0; j < len(r); j++ {
			x = append(x, j)
		}
		indices[i] = x
	}

	for _, p := range permutations(indices) {
		f := []Figure{}
		for n, pi := range p {
			f = append(f, rows[n][pi])
		}
		sets = append(sets, f)
	}
	return sets
}

func solutions(board Board, figures []Figure, left []string) []Board {
	results := []Board{}

	// Slice of figures possible rotations
	rots := [][]Figure{}

	for _, f := range figures {
		rots = append(rots, rotations(f))
	}

	// Get possible combinations
	figureSets := FigureCombinations(rots)

	// For each set of figures position them on board
	for _, set := range figureSets {
		coords := [][]int{}
		for _, f := range set {
			// get possible position coordinates
			for j := 0; j < board.height()-f.height(); j++ {
				for i := 0; i < board.width()-f.width(); i++ {
					coords = append(coords, []int{i, j})
				}
			}
		}

		fmt.Println("Board,", board)
		fmt.Println("Figures", set)
		for i, p := range permutations(coords) {
			fmt.Println("Positions on board", p)
		}
	}

	return results
}
