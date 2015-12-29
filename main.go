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
var reload = make(chan int)
var quit = make(chan int)
var inputs = make(chan string)

func main() {
	defer initLoggerOutput().Close()
	defer ncurses.Init()()

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	go func() {
		b := make([]byte, 1)
		for {
			os.Stdin.Read(b)
			inputs <- string(b)
		}
	}()

	go play()

	for {
		select {
		case <-quit:
			return
		case <-reload:
			go play()
		}
	}
}

func play() {
	p := game.NewStdPlayfield()
	changeOrInitPiece(&p)

	var c string
	var transform game.Transform

	for {
		render(p)
		c = <-inputs

		switch c {
		case "q":
			quit <- 1
			return
		case "r":
			reload <- 1
			return
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

		if p.IsGameOver() {
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

	if p.IsGameOver() {
		return
	}

	if p.NbCompleteLines() > 0 {
		render(*p)
		time.Sleep(time.Millisecond * 200)
		p.RemoveLines()
		render(*p)
	}
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
