package windeau

import (
	"github.com/nsf/termbox-go"
)

type Cell struct {
	Char   rune
	Fg, Bg termbox.Attribute
}

var BlankCell Cell = Cell{' ', termbox.ColorDefault, termbox.ColorDefault}

type Canvas struct {
	X, Y, Width, Height int
	Parent              *Drawable
	cells               [][]Cell
}

func MakeCanvas(x, y, w, h int) *Canvas {
	canvas := &Canvas{X: x, Y: y, Width: w, Height: h}
	canvas.initialize()
	return canvas
}

func (c *Canvas) GetRect() Rect {
	return Rect{c.X, c.Y, c.Width, c.Height}
}

func (c *Canvas) Fill(char rune, fg, bg termbox.Attribute) {
	c.FilledRect(char, fg, bg, c.GetRect())
}

func (c *Canvas) FilledRect(char rune, fg, bg termbox.Attribute, rect Rect) {
	if c.GetRect().DoesNotContain(rect) {
		return
	}

	offsetX := rect.X - c.X
	offsetY := rect.Y - c.Y

	for i := 0; i < rect.Width; i++ {
		for j := 0; j < rect.Height; j++ {
			c.MarkCell(i+offsetX, j+offsetY, char, fg, bg)
		}
	}
}

func (c *Canvas) MarkCell(x, y int, char rune, fg, bg termbox.Attribute) {
	c.cells[x][y] = Cell{char, fg, bg}
}

func (c *Canvas) Draw() {
	for x, row := range c.cells {
		for y, cell := range row {
			termbox.SetCell(x+c.X, y+c.Y, cell.Char, cell.Fg, cell.Bg)
		}
	}
}

func (c *Canvas) Cells() [][]Cell {
	return c.cells
}

func (c *Canvas) initialize() {
	c.cells = make([][]Cell, c.Width)
	for i := range c.cells {
		c.cells[i] = make([]Cell, c.Height)
	}
}
