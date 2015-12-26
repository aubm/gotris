package game

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
	blocs  []Bloc
}

// Bloc represent a playfield fragment to be rendered
type Bloc struct {
	X    int
	Y    int
	Code int
}

// Blocs returns all blocs that make up the playfield
func (p Playfield) Blocs() []Bloc {
	var blocs []Bloc
	blocs = p.blocs
	for _, part := range p.Piece.parts {
		blocs = append(blocs, Bloc{X: part.X, Y: part.Y, Code: p.Piece.code})
	}
	return blocs
}

// At is used to know if a piece exists at a given coordinate
func (p Playfield) At(c Coords) int {
	if c.X < 0 || c.X > (p.Width-1) || c.Y < 0 {
		return OutOfBounds
	}
	for _, part := range p.Piece.parts {
		if part.X == c.X && part.Y == c.Y {
			return p.Piece.code
		}
	}
	for _, bloc := range p.blocs {
		if bloc.X == c.X && bloc.Y == c.Y {
			return bloc.Code
		}
	}
	return Empty
}

// Fits checks if a given Tetromino can fit the playfield
func (p Playfield) Fits(piece Tetromino) bool {
	fakePiece(&p)
	for _, part := range piece.parts {
		if p.At(part) != Empty {
			return false
		}
	}
	return true
}

// RemoveLines removes lines that are complete and returns the number of removed lines
func (p *Playfield) RemoveLines() int {
	return p.scrollLines(func(blocs []Bloc, currentY *int, total *int) {
		if len(blocs) == p.Width {
			p.removeBlocsWithY(*currentY)
			p.stepDownBlocsOverY(*currentY)
			*total++
		} else {
			*currentY++
		}
	})
}

// NbCompleteLines returns the number of lines that are complete
func (p Playfield) NbCompleteLines() int {
	return p.scrollLines(func(blocs []Bloc, currentY *int, total *int) {
		if len(blocs) == p.Width {
			*total++
		}
		*currentY++
	})
}

func (p *Playfield) scrollLines(cb func(blocs []Bloc, currentY *int, total *int)) int {
	total := 0
	currentY := 0
	for currentY < p.Height {
		blocs := p.findBlocsWithY(currentY)
		cb(blocs, &currentY, &total)
	}
	return total
}

func (p Playfield) findBlocsWithY(y int) []Bloc {
	var blocs []Bloc
	for _, b := range p.blocs {
		if b.Y == y {
			blocs = append(blocs, b)
		}
	}
	return blocs
}

func (p *Playfield) removeBlocsWithY(y int) {
	var blocs []Bloc
	for _, b := range p.blocs {
		if b.Y != y {
			blocs = append(blocs, b)
		}
	}
	p.blocs = blocs
}

func (p *Playfield) stepDownBlocsOverY(y int) {
	for i, b := range p.blocs {
		if b.Y > y {
			p.blocs[i].Y--
		}
	}
}

func (p *Playfield) freezePiece() {
	for _, part := range p.Piece.parts {
		p.blocs = append(p.blocs, Bloc{X: part.X, Y: part.Y, Code: p.Piece.code})
	}
}

func fakePiece(p *Playfield) {
	part := Coords{p.Width, p.Height}
	p.Piece = Tetromino{parts: [4]Coords{part, part, part, part}}
}

// ChangeOrInitPiece creates a new piece and attachs it at the given playfield.
// If a piece is already attached to the playfield, it registers one new bloc
// for each part that compose the piece.
func ChangeOrInitPiece(p *Playfield) {
	if p.Piece.code != 0 {
		p.freezePiece()
	}
	p.Piece = getRandomPiece()
}

// NewStdPlayfield builds a new playfield of 10x20
func NewStdPlayfield() Playfield {
	return Playfield{Width: 10, Height: 20}
}
