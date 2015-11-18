package tetrominos

// Tetromino represents four parts that form a
// piece of the game
type Tetromino interface {
	Parts() []Coords
	SetParts(parts []Coords)
}

// Coords represents a pair of cartesian coordinates
type Coords struct {
	X int
	Y int
}

type shapeT struct {
	parts []Coords
}

func (s shapeT) Parts() []Coords {
	return s.parts
}

func (s *shapeT) SetParts(parts []Coords) {
	s.parts = parts
}

// T returns a specific implementation of a Tetromino
func T(center Coords) Tetromino {
	return &shapeT{[]Coords{
		center,
		Coords{center.X + 1, center.Y},
		Coords{center.X, center.Y - 1},
		Coords{center.X + 1, center.Y - 1},
		Coords{center.X - 1, center.Y - 1},
		Coords{center.X - 1, center.Y},
		Coords{center.X, center.Y + 1},
		Coords{center.X + 1, center.Y + 1},
		Coords{center.X + 2, center.Y - 1},
		Coords{center.X + 2, center.Y},
	}}
}
