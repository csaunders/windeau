package windeau

import (
	"github.com/nsf/termbox-go"
)

type SimpleBorder struct {
	edge, verticalBorder, horizontalBorder rune
}

func (sb SimpleBorder) Edge(x, y int, window Window) rune {
	return sb.edge
}

func (sb SimpleBorder) VerticalBorder() rune {
	return sb.verticalBorder
}

func (sb SimpleBorder) HorizontalBorder() rune {
	return sb.horizontalBorder
}

func MakeSimpleBorder(edge, vertical, horizontal rune) SimpleBorder {
	return SimpleBorder{edge: edge, verticalBorder: vertical, horizontalBorder: horizontal}
}

type PositionHandler func(x, y int)

func WalkRegion(x, y, w, h int, handler PositionHandler) {
	for px := x; px < x+w; px++ {
		for py := y; py < y+h; py++ {
			handler(px, py)
		}
	}
}

type Window struct {
	X, Y          int
	Width, Height int
	Border        WindowBorder
	Title         string
	Fg, Bg        termbox.Attribute
	View          *Drawable
}

func MakeBasicWindow(x, y, w, h int) *Window {
	border := MakeSimpleBorder('+', '|', '-')
	return &Window{X: x, Y: y, Width: w, Height: h, Border: border}
}

func (w *Window) Draw() {
	w.drawBorder()
	w.drawTitle()
	if w.View != nil {
		(*w.View).Draw()
	}
}

func (w *Window) SetParent(parent Drawable) {}

func (w *Window) IsRoot() bool {
	return true
}

func (w *Window) GetRect() Rect {
	return Rect{X: w.X, Y: w.Y, Width: w.Width, Height: w.Height}
}

func (w *Window) WithinBox(x, y int) bool {
	return w.GetRect().WithinRect(x, y)
}

func (w *Window) drawBorder() {
	WalkRegion(w.X, w.Y, w.Width, w.Height, func(x, y int) {
		var element rune
		if w.isEdge(x, y) {
			element = w.Border.Edge(x, y, *w)
		} else if w.isTopOrBottom(x, y) {
			element = w.Border.HorizontalBorder()
		} else if w.isLeftOrRight(x, y) {
			element = w.Border.VerticalBorder()
		} else {
			element = ' '
		}
		termbox.SetCell(x, y, element, w.Fg, w.Bg)
	})
}

func (w *Window) drawTitle() {
	title := w.Title
	if len(title) > 0 {
		if len(title) > w.titleWidth() {
			title = title[0:w.titleWidth()]
		}
		for c := 0; c < len(title); c++ {
			termbox.SetCell(c+w.X+w.titleStart(), w.Y, rune(title[c]), w.Fg, w.Bg)
		}
	}
}

func (w *Window) titleWidth() int {
	return w.Width - w.titleStart() - w.titlePadding()
}

func (w *Window) titleStart() int {
	return w.titlePadding()
}

func (w *Window) titlePadding() int {
	return 2
}

func (w *Window) isEdge(x, y int) bool {
	return w.isTopOrBottom(x, y) && w.isLeftOrRight(x, y)
}

func (w *Window) isTopOrBottom(x, y int) bool {
	return y == w.Y || y == w.Y+w.Height-1
}

func (w *Window) isLeftOrRight(x, y int) bool {
	return x == w.X || x == w.X+w.Width-1
}
