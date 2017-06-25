package angry_birds

import (
	"fmt"
	"reflect"
	"strings"
)

type Pig rune

type Figure []string

type Coord [2]int

var Skip string = "."
var Empty string = " "

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
		rs := []string{}
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

// valid returns true if figures on Board doesn't overlap
func (board Board) valid() bool {
	for i, f := range board.figures {
		for n, g := range board.figures {
			if n == i {
				continue
			}

			vboard := map[Coord]bool{}

			for _, fig := range []FigureOnBoard{f, g} {
				for j, row := range fig.figure {
					for i, v := range row {
						if string(v) != Skip {
							coord := Coord{i + fig.x, j + fig.y}
							if vboard[coord] {
								return false
							} else {
								vboard[coord] = true
							}
						}
					}
				}
			}
		}
	}

	return true
}

func (board Board) uncovered() map[string]int {
	m := map[string]int{}

	vboard := map[Coord]string{}

	// Convert board to map of coords to string
	for j, row := range board.board {
		for i, v := range row {
			coord := Coord{i, j}
			vboard[coord] = string(v)
		}
	}

	// Replace vboard position with Covered
	for _, fig := range board.figures {
		for j, row := range fig.figure {
			for i, v := range row {
				if string(v) != Skip {
					coord := Coord{fig.x + i, fig.y + j}
					vboard[coord] = Empty
				}
			}
		}
	}

	for _, v := range vboard {
		if string(v) != Empty {
			m[v] += 1
		}
	}

	return m
}

// pigs is a helper function to convert string to array of Pigs
func pigs(s string) []Pig {
	var p = make([]Pig, 0)
	for _, r := range []rune(s) {
		p = append(p, Pig(r))
	}
	return p
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

	fmt.Println("Board ", bw, " x ", bh)
	fmt.Println("Figure ", fw, " x ", fh)

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
		for j := 0; j < current_i; j++ {
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

func solutions(board Board, figures []Figure, left map[string]int) []Board {
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
		// say we have a set of 3 figures
		// for each figure we generate slice of possible positions
		// [][2]int{ {0,0}, {1,0}, {2,0}, ... }
		// AND list of indexes for each slice {0, 1, 2, 3, ...} to
		// get all possible combinations
		// Pass slice of indexes to permutations function
		// { {0,1,2,3}, {0,1,2}, {0,1,2,3,4,5} }
		// Permutation will return slice of combinations of positions
		// for all figure in set (indexes)
		// Map from indexes to actual positions of figures
		// Place figures on board and verify figures don't overlap
		// Get fields that left on board after figures positioned
		// and compare with target (left variable)

		coords := make([][][2]int, len(set))
		indexes := make([][]int, len(set))

		for fi, f := range set {
			ii := 0

			// get possible position coordinates
			for j := 0; j <= board.height()-f.height(); j++ {
				for i := 0; i <= board.width()-f.width(); i++ {
					coords[fi] = append(coords[fi], [2]int{i, j})

					indexes[fi] = append(indexes[fi], ii)
					ii += 1
				}
			}
		}

		//fmt.Println("Board,", board)
		//fmt.Println("Figures", set)
		//fmt.Println("Coords", coords)
		//fmt.Println("Indexes", indexes)

		for _, p := range permutations(indexes) {
			//fmt.Println("Positions on board", p)

			board.figures = make([]FigureOnBoard, len(set))

			//var figures = make([][2]int, len(set))

			for i, off := range p {
				//figures[i] = coords[i][off]
				c := coords[i][off]
				f := FigureOnBoard{
					figure: set[i],
					x:      c[0],
					y:      c[1],
				}
				board.figures = append(board.figures, f)
			}

			// Next we put figures on board and check if position valid
			// and verify what left
			if board.valid() {
				leftOnBoard := board.uncovered()
				fmt.Println("Left on board", leftOnBoard)
				if reflect.DeepEqual(leftOnBoard, left) {
					fmt.Printf("Found solution!\n%s\n", board)
					results = append(results, board)
					return results
				}
			}
		}
	}

	return results
}
