package game

import (
	"reflect"
	"testing"
)

func createPlayfield() Playfield {
	return Playfield{
		Width:  10,
		Height: 20,
		Piece:  Tetromino{parts: [4]Coords{{4, 11}, {5, 11}, {6, 11}, {7, 11}}, code: MAGENTA},
	}
}

func TestAt(t *testing.T) {
	in := []Coords{{4, 5}, {7, 11}, {0, 0}, {10, 15}, {5, 20}, {10, 20}, {9, 19}, {-5, 1}}
	out := []int{Empty, MAGENTA, Empty, OutOfBounds, Empty, OutOfBounds, Empty, OutOfBounds}

	pf := createPlayfield()

	var result int
	for i, c := range in {
		result = pf.At(c)
		if result != out[i] {
			t.Errorf("iteration nb %v : result == %v, expected %v", i+1, result, out[i])
		}
	}
}

func TestFits(t *testing.T) {
	in := [][4]Coords{
		{{4, 5}, {5, 5}, {6, 5}, {7, 5}},
		{{6, 11}, {7, 11}, {8, 11}, {9, 11}},
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		{{10, 15}, {11, 15}, {12, 15}, {13, 15}},
		{{5, 20}, {6, 20}, {7, 20}, {8, 20}},
		{{10, 20}, {11, 20}, {12, 20}, {13, 20}},
		{{6, 19}, {7, 19}, {8, 19}, {9, 19}},
		{{-5, 1}, {-4, 1}, {-3, 1}, {-2, 1}},
	}
	out := []bool{true, true, true, false, true, false, true, false}

	pf := createPlayfield()

	var result bool
	for i, c := range in {
		result = pf.Fits(Tetromino{parts: c})
		if result != out[i] {
			t.Errorf("iteration nb %v : result == %v, expected %v", i+1, result, out[i])
		}
	}
}

func TestNewStdPlayfield(t *testing.T) {
	pf := NewStdPlayfield()

	if pf.Width != 10 {
		t.Errorf("pf.Width == %v, expected == 10", pf.Width)
	}

	if pf.Height != 20 {
		t.Errorf("pf.Height == %v, expected == 20", pf.Height)
	}
}

func TestBlocs(t *testing.T) {
	pf := createPlayfield()
	expectedBlocs := []Bloc{
		{X: 4, Y: 11, Code: MAGENTA},
		{X: 5, Y: 11, Code: MAGENTA},
		{X: 6, Y: 11, Code: MAGENTA},
		{X: 7, Y: 11, Code: MAGENTA},
	}

	blocs := pf.Blocs()

	if reflect.DeepEqual(blocs, expectedBlocs) == false {
		t.Errorf("blocs == %v, expected %v", blocs, expectedBlocs)
	}
}
