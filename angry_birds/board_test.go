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
			pigs("ABC"),
			pigs("BCD"),
			pigs("ABC"),
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
			pigs("ABCD"),
			pigs("BCDA"),
			pigs("DACB"),
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

func TestFull(t *testing.T) {
	figures := []Figure{
		{
			".X.",
			"XXX",
		},
		{
			"YY",
			"YY",
			"Y.",
		},
	}

	board := Board{
		board: [][]Pig{
			pigs(" ABCD"),
			pigs("BCFEA"),
			pigs(" DEFA"),
			pigs("ABFED"),
			pigs("DBEA "),
		},
		figures: []FigureOnBoard{},
	}

	// A list of board pieces should left uncovered
	target := []string{"E", "C"}

	// For all combinations of figure rotations AND
	// all figure positions on board
	// DO following checks:
	// - figures don't overlap
	// - board has uncovered requested positions

	results := solutions(board, figures, target)
	if len(results) != 1 {
		t.Error("No results found")
	}
}
