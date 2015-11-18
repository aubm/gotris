package tetrominos

// Rotate applies a rotation on a tetromino and returns
// a new transformed one
func Rotate(t Tetromino) {
	parts := times10(t.Parts())
	center := parts[0]
	var newCenterIndex int

	newParts := []Coords{}
	for i, part := range parts {
		transformedPart := Coords{
			(center.X + 5) - part.Y + (center.Y - 5),
			(center.Y - 5) + part.X - (center.X + 5),
		}
		if transformedPart.X == center.X && transformedPart.Y == center.Y {
			newCenterIndex = i
		}
		newParts = append(newParts, transformedPart)
	}

	// replace tetromino's center at index 0
	newParts[0], newParts[newCenterIndex] = newParts[newCenterIndex], newParts[0]
	t.SetParts(dividePer10(newParts))
}

// MoveDown applies a translation of 1 point bottom
// on a tetromino and returns a new transformed one
func MoveDown(t Tetromino) Tetromino {
	return &shapeT{}
}

func times10(coords []Coords) []Coords {
	newCoords := []Coords{}
	for _, c := range coords {
		newCoords = append(newCoords, Coords{c.X * 10, c.Y * 10})
	}
	return newCoords
}

func dividePer10(coords []Coords) []Coords {
	newCoords := []Coords{}
	for _, c := range coords {
		newCoords = append(newCoords, Coords{c.X / 10, c.Y / 10})
	}
	return newCoords
}
