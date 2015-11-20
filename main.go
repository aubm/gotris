package main

import (
	"fmt"
	"github.com/aubm/gotris/tetrominos"
	"os"
	"os/exec"
	"time"
)

type playfield struct {
	piece  *tetrominos.Tetromino
	width  int
	height int
}

func (p playfield) At(c tetrominos.Coords) int {
	for _, part := range p.piece.Parts() {
		if part.X == c.X && part.Y == c.Y {
			return 1
		}
	}
	return -1
}

func newStdPlayfield() playfield {
	return playfield{width: 10, height: 20}
}

func main() {
	p := newStdPlayfield()
	p.piece = tetrominos.T(tetrominos.Coords{3, 18})

	draw(p)
	sleep()
	tetrominos.MoveDown(p.piece)
	draw(p)
	sleep()
	tetrominos.Rotate(p.piece)
	draw(p)
	sleep()
	tetrominos.Rotate(p.piece)
	draw(p)
	sleep()
	tetrominos.Rotate(p.piece)
	draw(p)
	sleep()
	tetrominos.Rotate(p.piece)
	draw(p)
	sleep()
}

func sleep() {
	d, _ := time.ParseDuration("1s")
	time.Sleep(d)
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func draw(p playfield) {
	clear()
	for y := p.height - 1; y >= 0; y-- {
		for x := 0; x < p.width; x++ {
			if p.At(tetrominos.Coords{x, y}) == 1 {
				fmt.Print("o ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
}
