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
	Parent              Drawable
	cells               [][]Cell
}

func MakeCanvas(x, y, w, h int) *Canvas {
	canvas := &Canvas{X: x, Y: y, Width: w, Height: h}
	canvas.initialize()
	return canvas
}

func (c *Canvas) GetRect() Rect {
	if c.Parent != nil {
		return c.Parent.GetRect().ShrinkBy(1)
	} else {
		return Rect{c.X, c.Y, c.Width, c.Height}
	}
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

func (c *Canvas) DrawString(x, y int, s string, fg, bg termbox.Attribute) {
	runes := ConvertToRuneArray(s)
	width := c.Width - x
	if y > c.Height || y < 0 {
		return
	}

	for i := 0; i < width && i < len(runes); i++ {
		c.MarkCell(i+x, y, runes[i], fg, bg)
	}
}

func (c *Canvas) Draw() {
	rect := c.GetRect()
	for x, row := range c.cells {
		for y, cell := range row {
			px := x + rect.X
			py := y + rect.Y

			if rect.WithinRect(px, py) && px < rect.X+rect.Width-1 && py < rect.Y+rect.Height-1 {
				termbox.SetCell(px, py, cell.Char, cell.Fg, cell.Bg)
			}
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
