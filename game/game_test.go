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

func createPlayfieldWithBlocs() Playfield {
	return Playfield{
		Width:  10,
		Height: 20,
		blocs: []Bloc{
			{X: 3, Y: 9}, {X: 4, Y: 9}, {X: 5, Y: 9}, {X: 6, Y: 9}, {X: 9, Y: 9},
			{X: 0, Y: 8}, {X: 1, Y: 8}, {X: 8, Y: 8}, {X: 9, Y: 8},
			{X: 0, Y: 7}, {X: 1, Y: 7}, {X: 2, Y: 7}, {X: 3, Y: 7}, {X: 4, Y: 7}, {X: 5, Y: 7}, {X: 6, Y: 7}, {X: 7, Y: 7},
			{X: 0, Y: 6}, {X: 1, Y: 6}, {X: 2, Y: 6}, {X: 3, Y: 6}, {X: 4, Y: 6}, {X: 5, Y: 6}, {X: 6, Y: 6}, {X: 7, Y: 6},
			{X: 0, Y: 5}, {X: 1, Y: 5}, {X: 2, Y: 5}, {X: 3, Y: 5}, {X: 4, Y: 5}, {X: 5, Y: 5}, {X: 6, Y: 5}, {X: 7, Y: 5}, {X: 8, Y: 5}, {X: 9, Y: 5},
			{X: 0, Y: 4}, {X: 1, Y: 4}, {X: 2, Y: 4}, {X: 3, Y: 4}, {X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 7, Y: 4}, {X: 8, Y: 4}, {X: 9, Y: 4},
			{X: 0, Y: 3}, {X: 1, Y: 3}, {X: 2, Y: 3}, {X: 3, Y: 3}, {X: 4, Y: 3}, {X: 5, Y: 3}, {X: 6, Y: 3}, {X: 7, Y: 3}, {X: 8, Y: 3}, {X: 9, Y: 3},
			{X: 8, Y: 2},
			{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}, {X: 4, Y: 1}, {X: 5, Y: 1}, {X: 6, Y: 1}, {X: 7, Y: 1}, {X: 8, Y: 1}, {X: 9, Y: 1},
			{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}, {X: 6, Y: 0}, {X: 7, Y: 0}, {X: 8, Y: 0}, {X: 9, Y: 0},
		},
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

func TestFindBlocWithY(t *testing.T) {
	pf := createPlayfieldWithBlocs()
	expected := []Bloc{{X: 3, Y: 9}, {X: 4, Y: 9}, {X: 5, Y: 9}, {X: 6, Y: 9}, {X: 9, Y: 9}}

	blocs := pf.findBlocsWithY(9)

	if reflect.DeepEqual(blocs, expected) == false {
		t.Errorf("blocs == %v, expected %v", blocs, expected)
	}
}

func TestRemoveBlocsWithY(t *testing.T) {
	pf := createPlayfieldWithBlocs()
	expected := []Bloc{
		{X: 3, Y: 9}, {X: 4, Y: 9}, {X: 5, Y: 9}, {X: 6, Y: 9}, {X: 9, Y: 9},
		{X: 0, Y: 8}, {X: 1, Y: 8}, {X: 8, Y: 8}, {X: 9, Y: 8},
		{X: 0, Y: 7}, {X: 1, Y: 7}, {X: 2, Y: 7}, {X: 3, Y: 7}, {X: 4, Y: 7}, {X: 5, Y: 7}, {X: 6, Y: 7}, {X: 7, Y: 7},
		{X: 0, Y: 5}, {X: 1, Y: 5}, {X: 2, Y: 5}, {X: 3, Y: 5}, {X: 4, Y: 5}, {X: 5, Y: 5}, {X: 6, Y: 5}, {X: 7, Y: 5}, {X: 8, Y: 5}, {X: 9, Y: 5},
		{X: 0, Y: 4}, {X: 1, Y: 4}, {X: 2, Y: 4}, {X: 3, Y: 4}, {X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 7, Y: 4}, {X: 8, Y: 4}, {X: 9, Y: 4},
		{X: 0, Y: 3}, {X: 1, Y: 3}, {X: 2, Y: 3}, {X: 3, Y: 3}, {X: 4, Y: 3}, {X: 5, Y: 3}, {X: 6, Y: 3}, {X: 7, Y: 3}, {X: 8, Y: 3}, {X: 9, Y: 3},
		{X: 8, Y: 2},
		{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}, {X: 4, Y: 1}, {X: 5, Y: 1}, {X: 6, Y: 1}, {X: 7, Y: 1}, {X: 8, Y: 1}, {X: 9, Y: 1},
		{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}, {X: 6, Y: 0}, {X: 7, Y: 0}, {X: 8, Y: 0}, {X: 9, Y: 0},
	}

	pf.removeBlocsWithY(6)

	if reflect.DeepEqual(pf.blocs, expected) == false {
		t.Errorf("pf.blocs == %v, expected %v", pf.blocs, expected)
	}
}

func TestStepDownBlocsOverY(t *testing.T) {
	pf := createPlayfieldWithBlocs()
	expected := []Bloc{
		{X: 3, Y: 8}, {X: 4, Y: 8}, {X: 5, Y: 8}, {X: 6, Y: 8}, {X: 9, Y: 8},
		{X: 0, Y: 7}, {X: 1, Y: 7}, {X: 8, Y: 7}, {X: 9, Y: 7},
		{X: 0, Y: 6}, {X: 1, Y: 6}, {X: 2, Y: 6}, {X: 3, Y: 6}, {X: 4, Y: 6}, {X: 5, Y: 6}, {X: 6, Y: 6}, {X: 7, Y: 6},
		{X: 0, Y: 5}, {X: 1, Y: 5}, {X: 2, Y: 5}, {X: 3, Y: 5}, {X: 4, Y: 5}, {X: 5, Y: 5}, {X: 6, Y: 5}, {X: 7, Y: 5},
		{X: 0, Y: 5}, {X: 1, Y: 5}, {X: 2, Y: 5}, {X: 3, Y: 5}, {X: 4, Y: 5}, {X: 5, Y: 5}, {X: 6, Y: 5}, {X: 7, Y: 5}, {X: 8, Y: 5}, {X: 9, Y: 5},
		{X: 0, Y: 4}, {X: 1, Y: 4}, {X: 2, Y: 4}, {X: 3, Y: 4}, {X: 4, Y: 4}, {X: 5, Y: 4}, {X: 6, Y: 4}, {X: 7, Y: 4}, {X: 8, Y: 4}, {X: 9, Y: 4},
		{X: 0, Y: 3}, {X: 1, Y: 3}, {X: 2, Y: 3}, {X: 3, Y: 3}, {X: 4, Y: 3}, {X: 5, Y: 3}, {X: 6, Y: 3}, {X: 7, Y: 3}, {X: 8, Y: 3}, {X: 9, Y: 3},
		{X: 8, Y: 2},
		{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}, {X: 4, Y: 1}, {X: 5, Y: 1}, {X: 6, Y: 1}, {X: 7, Y: 1}, {X: 8, Y: 1}, {X: 9, Y: 1},
		{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}, {X: 6, Y: 0}, {X: 7, Y: 0}, {X: 8, Y: 0}, {X: 9, Y: 0},
	}

	pf.stepDownBlocsOverY(5)

	if reflect.DeepEqual(pf.blocs, expected) == false {
		t.Errorf("pf.blocs == %v, expected %v", pf.blocs, expected)
	}
}

func TestRemoveLines(t *testing.T) {
	pf := createPlayfieldWithBlocs()
	expectedBlocs := []Bloc{
		{X: 3, Y: 5}, {X: 4, Y: 5}, {X: 5, Y: 5}, {X: 6, Y: 5}, {X: 9, Y: 5},
		{X: 0, Y: 4}, {X: 1, Y: 4}, {X: 8, Y: 4}, {X: 9, Y: 4},
		{X: 0, Y: 3}, {X: 1, Y: 3}, {X: 2, Y: 3}, {X: 3, Y: 3}, {X: 4, Y: 3}, {X: 5, Y: 3}, {X: 6, Y: 3}, {X: 7, Y: 3},
		{X: 0, Y: 2}, {X: 1, Y: 2}, {X: 2, Y: 2}, {X: 3, Y: 2}, {X: 4, Y: 2}, {X: 5, Y: 2}, {X: 6, Y: 2}, {X: 7, Y: 2},
		{X: 8, Y: 1},
		{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}, {X: 6, Y: 0}, {X: 7, Y: 0}, {X: 8, Y: 0}, {X: 9, Y: 0},
	}
	lines := pf.RemoveLines()

	if reflect.DeepEqual(pf.blocs, expectedBlocs) == false {
		t.Errorf("pf.blocs == %v, expected %v", pf.blocs, expectedBlocs)
	}

	if lines != 4 {
		t.Errorf("lines == %v, expected %v", lines, 4)
	}
}
