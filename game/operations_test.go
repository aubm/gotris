package game

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	shape := T(Coords{4, 10})
	expectedCoords := [][4]Coords{
		{{5, 11}, {5, 10}, {5, 9}, {6, 10}},
		{{4, 10}, {5, 10}, {6, 10}, {5, 9}},
		{{5, 11}, {5, 10}, {5, 9}, {4, 10}},
		{{4, 10}, {5, 10}, {6, 10}, {5, 11}},
	}

	for i := 0; i < len(expectedCoords); i++ {
		shape = Rotate(shape)
		if reflect.DeepEqual(shape.parts, expectedCoords[i]) == false {
			t.Errorf("rotation nb %v : shape.parts == %v, expected %v", i+1, shape.parts, expectedCoords[i])
		}
	}
}

func TestMoveDown(t *testing.T) {
	shape := T(Coords{4, 10})
	out := [4]Coords{{4, 9}, {5, 9}, {6, 9}, {5, 10}}
	newShape := MoveDown(shape)
	if reflect.DeepEqual(newShape.parts, out) == false {
		t.Errorf("newShape.parts == %v, expected %v", newShape.parts, out)
	}
}

func TestMoveLeft(t *testing.T) {
	shape := T(Coords{4, 10})
	out := [4]Coords{{3, 10}, {4, 10}, {5, 10}, {4, 11}}
	newShape := MoveLeft(shape)
	if reflect.DeepEqual(newShape.parts, out) == false {
		t.Errorf("newShape.parts == %v, expected %v", newShape.parts, out)
	}
}

func TestMoveRight(t *testing.T) {
	shape := T(Coords{4, 10})
	out := [4]Coords{{5, 10}, {6, 10}, {7, 10}, {6, 11}}
	newShape := MoveRight(shape)
	if reflect.DeepEqual(newShape.parts, out) == false {
		t.Errorf("newShape.parts == %v, expected %v", newShape.parts, out)
	}
}
