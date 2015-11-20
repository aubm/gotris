package tetrominos

// Coords represents a pair of cartesian coordinates
type Coords struct {
	X int
	Y int
}

// Tetromino represents parts that form a piece of the game
type Tetromino struct {
	parts        []Coords
	translations [][]Coords
}

// Parts returns parts that compose the tetromino
func (t Tetromino) Parts() []Coords {
	return t.parts
}

// T returns a specific implementation of a Tetromino
func T(fp Coords) *Tetromino {
	return &Tetromino{
		parts: []Coords{fp, Coords{fp.X + 1, fp.Y}, Coords{fp.X + 2, fp.Y}, Coords{fp.X + 1, fp.Y + 1}},
		translations: [][]Coords{
			{Coords{1, 1}, Coords{0, 0}, Coords{-1, -1}, Coords{1, -1}},
			{Coords{-1, -1}, Coords{0, 0}, Coords{1, 1}, Coords{-1, -1}},
			{Coords{1, 1}, Coords{0, 0}, Coords{-1, -1}, Coords{-1, 1}},
			{Coords{-1, -1}, Coords{0, 0}, Coords{1, 1}, Coords{1, 1}},
		},
	}
}
