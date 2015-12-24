package game

import (
	"reflect"
	"testing"
)

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
