package game

import "github.com/aubm/gotris/ui"

const (
	// OutOfBounds is the value returned by At if the given Coords are out of the playfield
	OutOfBounds = -1
	// Empty is the value returned by At if the given Coords don't match any part
	Empty = 0
)

// Playfield represents the current state of the game
type Playfield struct {
	Piece  Tetromino
	Width  int
	Height int
}

// Parts returns all parts that make up the playfield
func (p Playfield) Parts() [4]Coords {
	return p.Piece.Parts()
}

// At is used to know if a piece exists at a given coordinate
func (p Playfield) At(c Coords) int {
	if c.X < 0 || c.X > (p.Width-1) || c.Y < 0 {
		return OutOfBounds
	}

	for _, part := range p.Piece.Parts() {
		if part.X == c.X && part.Y == c.Y {
			return ui.MAGENTA
		}
	}
	return Empty
}

// Fits checks if a given Tetromino can fit the playfield
func (p Playfield) Fits(piece Tetromino) bool {
	fakePiece(&p)
	for _, part := range piece.Parts() {
		if p.At(part) != Empty {
			return false
		}
	}
	return true
}

func fakePiece(p *Playfield) {
	part := Coords{p.Width, p.Height}
	p.Piece = Tetromino{parts: [4]Coords{part, part, part, part}}
}

// NewStdPlayfield builds a new playfield of 10x20
func NewStdPlayfield() Playfield {
	return Playfield{Width: 10, Height: 20}
}
