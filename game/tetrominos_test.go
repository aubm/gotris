package game

import (
	"reflect"
	"testing"
)

func TestParts(t *testing.T) {
	out := [4]Coords{{4, 5}, {5, 6}, {6, 7}, {7, 8}}
	shape := Tetromino{parts: out}
	parts := shape.Parts()
	if reflect.DeepEqual(parts, out) == false {
		t.Errorf("parts == %v, expected %v", parts, out)
	}
}

func TestT(t *testing.T) {
	data := []struct {
		in  Coords
		out [4]Coords
	}{
		{
			Coords{4, 10},
			[4]Coords{{4, 10}, {5, 10}, {6, 10}, {5, 11}},
		},
	}
	for _, d := range data {
		if tetro := T(d.in); reflect.DeepEqual(tetro.parts, d.out) == false {
			t.Errorf("tetro.parts == %v, expected %v", tetro.parts, d.out)
		}
	}
}
