package game

// Coords represents a pair of cartesian coordinates
type Coords struct {
	X int
	Y int
}

// Tetromino represents parts that form a piece of the game
type Tetromino struct {
	parts        [4]Coords
	translations [][]Coords
}

// Parts returns parts that compose the tetromino
func (t Tetromino) Parts() [4]Coords {
	return t.parts
}

// T returns a specific implementation of a Tetromino
func T(fp Coords) Tetromino {
	return Tetromino{
		parts: [4]Coords{fp, {fp.X + 1, fp.Y}, {fp.X + 2, fp.Y}, {fp.X + 1, fp.Y + 1}},
		translations: [][]Coords{
			{{1, 1}, {0, 0}, {-1, -1}, {1, -1}},
			{{-1, -1}, {0, 0}, {1, 1}, {-1, -1}},
			{{1, 1}, {0, 0}, {-1, -1}, {-1, 1}},
			{{-1, -1}, {0, 0}, {1, 1}, {1, 1}},
		},
	}
}
