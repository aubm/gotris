package game

// Transform is a shortcut for all transform func signature
type Transform func(t Tetromino) Tetromino

// Rotate applies a rotation on a tetromino
func Rotate(t Tetromino) Tetromino {
	if len(t.translations) > 0 {
		for i, trans := range t.translations[0] {
			t.parts[i].X += trans.X
			t.parts[i].Y += trans.Y
		}
		t.translations = append(t.translations[1:], t.translations[0])
	}
	return t
}

// MoveDown applies a translation of 1 point bottom on a tetromino
func MoveDown(t Tetromino) Tetromino {
	for i := range t.parts {
		t.parts[i].Y--
	}
	return t
}

// MoveLeft applies a translation of 1 point left on a tetromino
func MoveLeft(t Tetromino) Tetromino {
	for i := range t.parts {
		t.parts[i].X--
	}
	return t
}

// MoveRight applies a translation of 1 point right on a tetromino
func MoveRight(t Tetromino) Tetromino {
	for i := range t.parts {
		t.parts[i].X++
	}
	return t
}
