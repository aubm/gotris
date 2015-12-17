package tetrominos

// Rotate applies a rotation on a tetromino
func Rotate(t *Tetromino) {
	for i, trans := range t.translations[0] {
		t.parts[i].X += trans.X
		t.parts[i].Y += trans.Y
	}
	t.translations = append(t.translations[1:], t.translations[0])
}

// MoveDown applies a translation of 1 point bottom on a tetromino
func MoveDown(t *Tetromino) {
	for i := range t.parts {
		t.parts[i].Y--
	}
}

// MoveLeft applies a translation of 1 point left on a tetromino
func MoveLeft(t *Tetromino) {
	for i := range t.parts {
		t.parts[i].X--
	}
}

// MoveRight applies a translation of 1 point right on a tetromino
func MoveRight(t *Tetromino) {
	for i := range t.parts {
		t.parts[i].X++
	}
}
