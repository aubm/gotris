package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/aubm/gotris/game"
	"github.com/aubm/gotris/tetrominos"
	"github.com/aubm/gotris/ui/ncurses"
)

func main() {
	p := game.NewStdPlayfield()
	p.Piece = tetrominos.T(tetrominos.Coords{3, 19})

	defer initLoggerOutput().Close()
	endFunc := ncurses.Init()
	defer endFunc()

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	b := make([]byte, 1)
main:
	for {
		ncurses.Render(p)
		os.Stdin.Read(b)
		switch string(b) {
		case "q":
			break main
		case "A": // up
			tetrominos.Rotate(p.Piece)
		case "C": // right
			break
		case "B": // bottom
			tetrominos.MoveDown(p.Piece)
		case "D": // left
			break
		}
		log.Printf("User input : %v", string(b))
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
