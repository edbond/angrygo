package angry_birds

import "testing"

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
