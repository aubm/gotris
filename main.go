package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/aubm/gotris/game"
	"github.com/aubm/gotris/ui/ncurses"
)

func main() {
	p := game.NewStdPlayfield()
	p.Piece = game.T(game.Coords{3, 19})

	defer initLoggerOutput().Close()
	defer ncurses.Init()()

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	b := make([]byte, 1)
	var transform game.Transform
	var newPiece game.Tetromino
main:
	for {
		ncurses.Render(p, p.Width, p.Height)
		os.Stdin.Read(b)
		switch string(b) {
		case "q":
			break main
		case "A", "k": // up
			transform = game.Rotate
		case "C", "l": // right
			transform = game.MoveRight
		case "B", "j": // bottom
			transform = game.MoveDown
		case "D", "h": // left
			transform = game.MoveLeft
		default:
			continue
		}
		if transform != nil {
			newPiece = transform(p.Piece)
			if p.Fits(newPiece) {
				p.Piece = newPiece
			}
		}
	}
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
