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
	defer initLoggerOutput().Close()
	defer ncurses.Init()()

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	p := game.NewStdPlayfield()
	game.ChangeOrInitPiece(&p)

	b := make([]byte, 1)
	var transform game.Transform
	var newPiece game.Tetromino
main:
	for {
		ncurses.Render(p, p.Width, p.Height)
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

		for {
			newPiece = transform(p.Piece)
			if p.Fits(newPiece) {
				p.Piece = newPiece
			} else {
				break
			}

			if c != "L" && c != "$" && c != " " && c != "J" && c != "H" && c != "0" {
				break
			}
		}

		if c == " " || c == "J" {
			game.ChangeOrInitPiece(&p)
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
