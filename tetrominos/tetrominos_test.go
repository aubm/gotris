package tetrominos

import (
	"reflect"
	"testing"
)

func TestT(t *testing.T) {
	data := []struct {
		in  Coords
		out *shapeT
	}{
		{
			Coords{5, 14},
			&shapeT{[]Coords{{5, 14}, {6, 14}, {5, 13}, {6, 13}, {4, 13}, {4, 14}, {5, 15}, {6, 15}, {7, 13}, {7, 14}}},
		},
	}

	for _, d := range data {
		if out := T(d.in); reflect.DeepEqual(out, d.out) == false {
			t.Errorf("out == %v, expected %v", out, d.out)
		}
	}
}
