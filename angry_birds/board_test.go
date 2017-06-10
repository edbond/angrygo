package angry_birds

import (
	"fmt"
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

	fmt.Println(b)

	p := Positions(f, b)
	if len(p) != 2 {
		fmt.Println(p)
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
	expected := "\nABCD\nBCDX\nDXXX"

	if result != expected {
		t.Errorf("%s\n%s\n%s\n%s", "Board printed incorrectly", result, "Expected:", expected)
	}

}
