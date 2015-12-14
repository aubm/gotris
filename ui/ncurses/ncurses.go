package ncurses

// http://mrcook.uk/how-to-install-go-ncurses-on-mac-osx

import (
	"log"
	"math"

	"github.com/aubm/gotris/game"
	"github.com/aubm/gotris/tetrominos"
	"github.com/aubm/gotris/ui"
	gc "github.com/rthornton128/goncurses"
)

var stdscr *gc.Window

// Render takes a playfield and reprensents it in a terminal
// using ncurses libraires
func Render(p game.Playfield) {
	for _, part := range p.Piece.Parts() {
		defer drawPart(part, p.Width, p.Height)()
	}

	gc.Update()
}

func drawPart(c tetrominos.Coords, fieldWidth int, fieldHeight int) func() {
	var size int
	maxY, maxX := stdscr.MaxYX()
	maxY, maxX = maxY/fieldHeight, maxX/fieldWidth/2
	size = int(math.Min(float64(maxY), float64(maxX)))

	h := size
	w := h * 2
	y := (fieldHeight - c.Y) * h
	x := c.X * w

	win, err := gc.NewWindow(h, w, y, x)
	if err != nil {
		log.Fatal(err)
	}

	win.Box(gc.ACS_VLINE, gc.ACS_HLINE)
	win.SetBackground(gc.Char(' ') | gc.ColorPair(ui.MAGENTA))
	win.NoutRefresh()

	return func() {
		win.Erase()
		win.SetBackground(gc.Char(' '))
		win.NoutRefresh()
		win.Delete()
	}
}

// Init configures ncurses and initialize a standard screen.
// It returns a function that must be called before the program exits.
func Init() func() {
	initStdscr()
	initConfiguration()
	initColors()
	return gc.End
}

func initConfiguration() {
	gc.CBreak(true)
	gc.Echo(false)
	gc.Cursor(0)
}

func initStdscr() {
	var err error
	stdscr, err = gc.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	stdscr.NoutRefresh()
}

func initColors() {
	gc.StartColor()
	gc.InitPair(ui.MAGENTA, gc.C_MAGENTA, gc.C_MAGENTA)
}
