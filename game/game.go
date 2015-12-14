package game

import "github.com/aubm/gotris/tetrominos"

// Playfield represents the current state of the game
// It stores the current piece and the board size
type Playfield struct {
	Piece  *tetrominos.Tetromino
	Width  int
	Height int
}

// At is used to know if a piece exists at a given coordinate
func (p Playfield) At(c tetrominos.Coords) int {
	for _, part := range p.Piece.Parts() {
		if part.X == c.X && part.Y == c.Y {
			return 1
		}
	}
	return -1
}

// NewStdPlayfield builds a new playfield of 10x20
func NewStdPlayfield() Playfield {
	return Playfield{Width: 10, Height: 20}
}
