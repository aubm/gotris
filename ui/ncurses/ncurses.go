package ncurses

import (
	"log"
	"math"

	"github.com/aubm/gotris/game"
	gc "github.com/rthornton128/goncurses"
)

var stdscr *gc.Window

type blocsBag interface {
	Blocs() []game.Bloc
}

// Render takes a playfield and reprensents it in a terminal
// using ncurses libraires
func Render(p blocsBag, fieldWidth int, fieldHeight int) {
	metrics := calculateMetrics(fieldWidth, fieldHeight)
	defer createWindow(metrics.height*fieldHeight+2, metrics.width*fieldWidth+2,
		metrics.offsetHeight+metrics.height-1, metrics.offsetWidth-1,
		0, true)()

	for _, bloc := range p.Blocs() {
		defer drawBloc(bloc, fieldWidth, fieldHeight)()
	}

	gc.Update()
}

func drawBloc(bloc game.Bloc, fieldWidth int, fieldHeight int) func() {
	metrics := calculateMetrics(fieldWidth, fieldHeight)

	y := (fieldHeight-bloc.Y)*metrics.height + metrics.offsetHeight
	x := bloc.X*metrics.width + metrics.offsetWidth

	return createWindow(metrics.height, metrics.width, y, x, bloc.Code, false)
}

type metrics struct {
	width        int
	height       int
	offsetWidth  int
	offsetHeight int
}

func calculateMetrics(fieldWidth int, fieldHeight int) metrics {
	var size int
	maxY, maxX := stdscr.MaxYX()
	fieldMaxY, fieldMaxX := maxY/fieldHeight, maxX/fieldWidth/2
	size = int(math.Min(float64(fieldMaxY), float64(fieldMaxX)))
	h := size
	w := h * 2
	offsetWidth, offsetHeight := (maxX-fieldWidth*w)/2, 0
	if maxY-(fieldMaxY*h) >= h {
		offsetHeight = h
	}
	return metrics{
		width:        w,
		height:       h,
		offsetWidth:  offsetWidth,
		offsetHeight: offsetHeight,
	}
}

func createWindow(h, w, y, x, bgColor int, borders bool) func() {
	win, err := gc.NewWindow(h, w, y, x)
	if err != nil {
		log.Fatal(err)
	}
	if bgColor != 0 {
		win.SetBackground(gc.Char(' ') | gc.ColorPair(int16(bgColor)))
	}
	if borders {
		win.Box(gc.ACS_VLINE, gc.ACS_HLINE)
	}
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
	gc.InitPair(game.BLUE, gc.C_BLUE, gc.C_BLUE)
	gc.InitPair(game.CYAN, gc.C_CYAN, gc.C_CYAN)
	gc.InitPair(game.GREEN, gc.C_GREEN, gc.C_GREEN)
	gc.InitPair(game.MAGENTA, gc.C_MAGENTA, gc.C_MAGENTA)
	gc.InitPair(game.RED, gc.C_RED, gc.C_RED)
	gc.InitPair(game.WHITE, gc.C_WHITE, gc.C_WHITE)
	gc.InitPair(game.YELLOW, gc.C_YELLOW, gc.C_YELLOW)
}
