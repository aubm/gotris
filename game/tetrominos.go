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
	code         int
}

// I returns the I-block tetromino
func I(fp Coords) Tetromino {
	return Tetromino{
		parts: [4]Coords{fp, {fp.X + 1, fp.Y}, {fp.X + 2, fp.Y}, {fp.X + 3, fp.Y}},
		translations: [][]Coords{
			{{1, 1}, {0, 0}, {-1, -1}, {-2, -2}},
			{{-1, -1}, {0, 0}, {1, 1}, {2, 2}},
		},
		code: CYAN,
	}
}

// J returns the J-block tetromino
func J(fp Coords) Tetromino {
	return Tetromino{
		parts: [4]Coords{fp, {fp.X, fp.Y - 1}, {fp.X + 1, fp.Y - 1}, {fp.X + 2, fp.Y - 1}},
		translations: [][]Coords{
			{{1, -1}, {1, 1}, {0, 2}, {0, 2}},
			{{-1, 1}, {0, 0}, {1, -1}, {0, -2}},
			{{0, -1}, {0, -1}, {-1, 0}, {-1, 2}},
			{{0, 1}, {-1, 0}, {0, -1}, {1, -2}},
		},
		code: BLUE,
	}
}

// L returns the L-block tetromino
func L(fp Coords) Tetromino {
	return Tetromino{
		parts: [4]Coords{fp, {fp.X + 1, fp.Y}, {fp.X + 2, fp.Y}, {fp.X + 2, fp.Y + 1}},
		translations: [][]Coords{
			{{2, 0}, {0, 0}, {-1, 1}, {-1, 1}},
			{{-2, 0}, {-1, 1}, {0, 0}, {1, -1}},
			{{1, 0}, {1, 0}, {0, 1}, {-2, 1}},
			{{-1, 0}, {0, -1}, {1, -2}, {2, -1}},
		},
		code: WHITE,
	}
}

// O returns the O-block tetromino
func O(fp Coords) Tetromino {
	return Tetromino{
		parts: [4]Coords{fp, {fp.X + 1, fp.Y}, {fp.X, fp.Y - 1}, {fp.X + 1, fp.Y - 1}},
		code:  YELLOW,
	}
}

// S returns the S-block tetromino
func S(fp Coords) Tetromino {
	return Tetromino{
		parts: [4]Coords{fp, {fp.X - 1, fp.Y}, {fp.X - 1, fp.Y - 1}, {fp.X - 2, fp.Y - 1}},
		translations: [][]Coords{
			{{-1, -1}, {0, 0}, {-1, 1}, {0, 2}},
			{{1, 1}, {0, 0}, {1, -1}, {0, -2}},
		},
		code: GREEN,
	}
}

// T returns the T-block tetromino
func T(fp Coords) Tetromino {
	return Tetromino{
		parts: [4]Coords{fp, {fp.X + 1, fp.Y}, {fp.X + 2, fp.Y}, {fp.X + 1, fp.Y + 1}},
		translations: [][]Coords{
			{{1, 1}, {0, 0}, {-1, -1}, {1, -1}},
			{{-1, -1}, {0, 0}, {1, 1}, {-1, -1}},
			{{1, 1}, {0, 0}, {-1, -1}, {-1, 1}},
			{{-1, -1}, {0, 0}, {1, 1}, {1, 1}},
		},
		code: MAGENTA,
	}
}

// Z returns the Z-block tetromino
func Z(fp Coords) Tetromino {
	return Tetromino{
		parts: [4]Coords{fp, {fp.X + 1, fp.Y}, {fp.X + 1, fp.Y - 1}, {fp.X + 2, fp.Y - 1}},
		translations: [][]Coords{
			{{1, -1}, {0, 0}, {1, 1}, {0, 2}},
			{{-1, 1}, {0, 0}, {-1, -1}, {0, -2}},
		},
		code: RED,
	}
}
