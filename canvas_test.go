package windeau

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func assert_canvas_cells(t *testing.T, c rune, fg, bg termbox.Attribute, cells [][]Cell) {
	for _, row := range cells {
		for _, cell := range row {
			assert_canvas_cell(t, c, fg, bg, cell)
		}
	}
}

func assert_canvas_cell(t *testing.T, c rune, fg, bg termbox.Attribute, cell Cell) {
	assert.Equal(t, c, cell.Char, fmt.Sprintf("Cell should've been %v but was %v instead", c, cell.Char))
	assert.Equal(t, fg, cell.Fg, "Cell foreground colour should be Red")
	assert.Equal(t, bg, cell.Bg, "Cell background colour should be Default")
}

func TestFillingCanvas(t *testing.T) {
	canvas := MakeCanvas(0, 0, 2, 2)
	canvas.Fill('x', termbox.ColorRed, termbox.ColorDefault)
	cells := canvas.Cells()
	assert_canvas_cells(t, 'x', termbox.ColorRed, termbox.ColorDefault, cells)
}

func TestFillingARegionInsideTheCanvas(t *testing.T) {
	canvas := MakeCanvas(0, 0, 5, 5)
	canvas.Fill('x', termbox.ColorRed, termbox.ColorDefault)
	canvas.FilledRect('y', termbox.ColorGreen, termbox.ColorDefault, Rect{2, 2, 1, 1})
	cells := canvas.Cells()
	assert_canvas_cell(t, 'y', termbox.ColorGreen, termbox.ColorDefault, cells[2][2])
}
