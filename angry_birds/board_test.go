package angry_birds

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNotEqual(t *testing.T) {
	f1 := Figure{"X"}
	f2 := Figure{"Y"}

	if !f1.notEqual(f2) {
		t.Error("Should be not equal", f1, f2)
	}

	if f1.notEqual(f1) {
		t.Error("Same should be equal")
	}
}

func TestFigureRotation(t *testing.T) {
	f1 := Figure{
		".X.",
		"XXX",
	}
	f1Rotated := Figure{
		"X.",
		"XX",
		"X.",
	}

	rotated := f1.RotateCW()
	if f1Rotated.notEqual(rotated) {
		t.Error("Rotated != expected,", f1Rotated, rotated)
	}
}

func TestPositions(t *testing.T) {
	f := Figure{
		".X.",
		"XXX",
	}

	b := Board{
		[][]Pig{
			Pigs("ABC"),
			Pigs("BCD"),
			Pigs("ABC"),
		},
		[]FigureOnBoard{},
	}

	//fmt.Println(b)

	p := Positions(f, b)
	if len(p) != 2 {
		//fmt.Println(p)
		t.Error("Positions != 2", len(p))
	}
}

func TestBoardPrinting(t *testing.T) {
	f := Figure{
		"..X",
		"XXX",
	}

	b := Board{
		[][]Pig{
			Pigs("ABCD"),
			Pigs("BCDA"),
			Pigs("DACB"),
		},
		[]FigureOnBoard{
			{f, 1, 1},
		},
	}

	result := fmt.Sprintf("%v", b)
	expected := "ABCD\nBCDX\nDXXX"

	if result != expected {
		t.Errorf("%s\n%q\n%s\n%q", "Board printed incorrectly", result, "Expected:", expected)
	}

}

func TestRotations(t *testing.T) {
	f := Figure{
		"X",
		"X",
	}

	expected := []Figure{
		{"X", "X"},
		{"XX"},
	}

	results := rotations(f)
	if len(results) != 2 {
		t.Error("Invalid rotations", f, results)
	}

	if !reflect.DeepEqual(results, expected) {
		t.Error("Rotations invalid", results, "Expected", expected)
	}
}

func TestPermutations(t *testing.T) {
	input := [][]int{
		{1, 2},
		{3},
		{6, 7},
	}

	expected := [][]int{
		{1, 3, 6},
		{2, 3, 6},
		{1, 3, 7},
		{2, 3, 7},
	}

	result := permutations(input)
	if len(result) != len(expected) {
		t.Error("Invalid number of permutation results", len(result), "expected", len(expected))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error("Result != Expected", result, " != ", expected)
	}

}

func TestPermutations2(t *testing.T) {
	input := [][]int{
		{0, 1},
		{0},
	}

	expected := [][]int{
		{0, 0},
		{1, 0},
	}

	result := permutations(input)
	if len(result) != len(expected) {
		t.Error("Invalid number of permutation results", len(result), "expected", len(expected))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error("Result != Expected", result, " != ", expected)
	}
}

func TestPermutations3(t *testing.T) {
	input := [][]int{
		{0, 1, 2},
		{3, 4, 5},
	}

	expected := [][]int{
		{0, 3},
		{1, 3},
		{2, 3},
		{0, 4},
		{1, 4},
		{2, 4},
		{0, 5},
		{1, 5},
		{2, 5},
	}

	result := permutations(input)
	if len(result) != len(expected) {
		t.Error("Invalid number of permutation results", len(result), "expected", len(expected))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error("Result != Expected", result, " != ", expected)
	}
}

func TestFigureCombinations(t *testing.T) {
	input := [][]Figure{
		{Figure{"X", "X"}, Figure{"XX"}},
		{Figure{
			"X",
			"XX",
			"XX",
		}},
	}

	expected := [][]Figure{
		{Figure{"X", "X"}, Figure{"X", "XX", "XX"}},
		{Figure{"XX"}, Figure{"X", "XX", "XX"}},
	}

	result := FigureCombinations(input)

	if len(result) != len(expected) {
		t.Error("Figure combinations count not correct", len(result), " != ", len(expected))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error("Figure combinations are not correct", " Result ", result, " Expected ", expected)
	}
}

func TestFull(t *testing.T) {
	figures := []Figure{
		{
			"X",
			"X",
		},
		{
			"Y",
		},
	}

	board := Board{
		Board: [][]Pig{
			Pigs(" A"),
			Pigs("BC"),
		},
		figures: []FigureOnBoard{},
	}

	// A list of board pieces should left uncovered
	target := map[string]int{
		"C": 1,
	}

	// For all combinations of figure rotations AND
	// all figure positions on board
	// DO following checks:
	// - figures don't overlap
	// - board has uncovered requested positions
	results := Solutions(board, figures, target)
	if len(results) < 1 {
		t.Error("No results found")
	}
}

func TestBoardValid(t *testing.T) {
	board := Board{
		Board: [][]Pig{
			Pigs(" ABCD"),
			Pigs("BCFEA"),
			Pigs(" DEFA"),
			Pigs("ABFED"),
			Pigs("DBEA "),
		},
		figures: []FigureOnBoard{
			{
				figure: Figure{"XX", "XX"},
				x:      0,
				y:      0,
			},
			{
				figure: Figure{"YY", "YY"},
				x:      1,
				y:      1,
			},
		},
	}

	if board.valid() {
		t.Errorf("Invalid board returns true for valid()\n%s\n", board)
	}
}

func TestBoardValid2(t *testing.T) {
	board := Board{
		Board: [][]Pig{
			Pigs(" ABCD"),
			Pigs("BCFEA"),
			Pigs(" DEFA"),
			Pigs("ABFED"),
			Pigs("DBEA "),
		},
		figures: []FigureOnBoard{
			{
				figure: Figure{"XX", "X", "X"},
				x:      0,
				y:      0,
			},
			{
				figure: Figure{".Y", ".Y", "YY"},
				x:      0,
				y:      1,
			},
		},
	}

	if !board.valid() {
		t.Errorf("Invalid board returns false for valid()\n%s\n", board)
	}
}

func TestFull2(t *testing.T) {
	board := Board{
		Board: [][]Pig{
			Pigs("HSP  "),
			Pigs(" AHBP"),
			Pigs(" SRPS"),
			Pigs("HSAPP"),
			Pigs("B HAS"),
		},
	}

	figures := []Figure{
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
		"R": 1,
		//"S": 2,
		//"P": 1,
	}

	results := Solutions(board, figures, target)
	if len(results) < 1 {
		t.Error("No results found")
	}
}
