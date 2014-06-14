package main

import (
	"github.com/csaunders/windeau"
	"github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	termbox.Clear(termbox.ColorGreen, termbox.ColorBlack)
	defer termbox.Close()

	border := windeau.MakeSimpleBorder('+', '|', '-')
	window := &windeau.Window{X: 0, Y: 0, Width: 30, Height: 25, Fg: termbox.ColorGreen, Bg: termbox.ColorDefault, Border: border}
	window.Title = "Hello World"
	window.Draw()

	fancyBorder := FancyBorder{TopLeft: '┏', TopRight: '┓', BottomLeft: '┗', BottomRight: '┛', Vertical: '┃', Horizontal: '━'}
	fancyWindow := &windeau.Window{X: 40, Y: 20, Width: 40, Height: 30, Fg: termbox.ColorBlue, Bg: termbox.ColorDefault, Border: fancyBorder}
	fancyWindow.Title = "Fancy Window"
	fancyWindow.Draw()

	termbox.Flush()

	termbox.PollEvent()
}

type FancyBorder struct {
	TopLeft, TopRight, BottomLeft, BottomRight rune
	Vertical, Horizontal                       rune
}

func (f FancyBorder) Edge(x, y int, w windeau.Window) rune {
	switch {
	case x == w.X && y == w.Y:
		return f.TopLeft
	case x == w.X+w.Width-1 && y == w.Y:
		return f.TopRight
	case x == w.X && y == w.Y+w.Height-1:
		return f.BottomLeft
	case x == w.X+w.Width-1 && y == w.Y+w.Height-1:
		return f.BottomRight
	default:
		return ' '
	}
}

func (f FancyBorder) VerticalBorder() rune {
	return f.Vertical
}

func (f FancyBorder) HorizontalBorder() rune {
	return f.Horizontal
}
