package tetrominos

import (
	"reflect"
	"testing"
)

func TestParts(t *testing.T) {
	out := []Coords{Coords{4, 5}, Coords{5, 6}}
	shape := Tetromino{parts: out}
	parts := shape.Parts()
	if reflect.DeepEqual(parts, out) == false {
		t.Errorf("parts == %v, expected %v", parts, out)
	}
}

func TestT(t *testing.T) {
	data := []struct {
		in  Coords
		out []Coords
	}{
		{
			Coords{4, 10},
			[]Coords{Coords{4, 10}, Coords{5, 10}, Coords{6, 10}, Coords{5, 11}},
		},
	}
	for _, d := range data {
		if tetro := T(d.in); reflect.DeepEqual(tetro.parts, d.out) == false {
			t.Errorf("tetro.parts == %v, expected %v", tetro.parts, d.out)
		}
	}
}
