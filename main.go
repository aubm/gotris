package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/aubm/gotris/game"
	"github.com/aubm/gotris/ui/ncurses"
	"github.com/aubm/interval"
)

var stopInterval = func() {}

func main() {
	defer initLoggerOutput().Close()
	defer ncurses.Init()()

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	p := game.NewStdPlayfield()
	changeOrInitPiece(&p)

	b := make([]byte, 1)
	var transform game.Transform
main:
	for {
		render(p)
		os.Stdin.Read(b)
		c := string(b)
		switch c {
		case "q":
			break main
		case "A", "k": // up
			transform = game.Rotate
		case "C", "l", "L", "$": // right
			transform = game.MoveRight
		case "B", "j", "J", " ": // bottom
			transform = game.MoveDown
		case "D", "h", "H", "0": // left
			transform = game.MoveLeft
		default:
			continue
		}

		applyTransform(transform, &p, c == "L" || c == "$" || c == " " || c == "J" || c == "H" || c == "0")

		if c == " " || c == "J" {
			changeOrInitPiece(&p)
		}
	}
}

func changeOrInitPiece(p *game.Playfield) {
	stopInterval()
	game.ChangeOrInitPiece(p)
	stopInterval = interval.Start(func() {
		if applyTransform(game.MoveDown, p, false) == false {
			changeOrInitPiece(p)
		}
		render(*p)
	}, time.Second)
}

func applyTransform(transform game.Transform, p *game.Playfield, repeat bool) bool {
	fits := true
	var newPiece game.Tetromino
	for {
		newPiece = transform(p.Piece)
		if p.Fits(newPiece) {
			p.Piece = newPiece
		} else {
			fits = false
			break
		}

		if !repeat {
			break
		}
	}
	return fits
}

func render(p game.Playfield) {
	ncurses.Render(p, p.Width, p.Height)
}

func initLoggerOutput() *os.File {
	logFileName := "gotris.log"
	out, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to open %v", logFileName))
	}
	log.SetOutput(out)
	return out
}
