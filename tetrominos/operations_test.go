package tetrominos

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	shape := T(Coords{4, 10})
	expectedCoords := [][]Coords{
		[]Coords{Coords{5, 11}, Coords{5, 10}, Coords{5, 9}, Coords{6, 10}},
		[]Coords{Coords{4, 10}, Coords{5, 10}, Coords{6, 10}, Coords{5, 9}},
		[]Coords{Coords{5, 11}, Coords{5, 10}, Coords{5, 9}, Coords{4, 10}},
		[]Coords{Coords{4, 10}, Coords{5, 10}, Coords{6, 10}, Coords{5, 11}},
	}

	for i := 0; i < len(expectedCoords); i++ {
		Rotate(shape)
		if reflect.DeepEqual(shape.parts, expectedCoords[i]) == false {
			t.Errorf("rotation nb %v : shape.parts == %v, expected %v", i+1, shape.parts, expectedCoords[i])
		}
	}
}

func TestMoveDown(t *testing.T) {
	shape := T(Coords{4, 10})
	out := []Coords{Coords{4, 9}, Coords{5, 9}, Coords{6, 9}, Coords{5, 10}}
	MoveDown(shape)
	if reflect.DeepEqual(shape.parts, out) == false {
		t.Errorf("shape.parts == %v, expected %v", shape.parts, out)
	}
}
