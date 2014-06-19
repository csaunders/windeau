package main

import (
	"github.com/csaunders/windeau"
	"github.com/nsf/termbox-go"
	"os"
)

var window, fancyWindow *windeau.Window
var focusableWindow *windeau.FocusableWindow
var canvas *windeau.Canvas

func main() {
	termbox.Init()
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputMouse)

	prepareWindows()

	for true {
		draw()
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventMouse:
			handleMouse(&event)
		case termbox.EventKey:
			if event.Key == termbox.KeyEsc {
				os.Exit(0)
			}
		}
	}
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

func draw() {
	termbox.Clear(termbox.ColorGreen, termbox.ColorBlack)
	window.Draw()
	fancyWindow.Draw()
	focusableWindow.Draw()
	canvas.Draw()
	termbox.Flush()
}

func handleMouse(ev *termbox.Event) {
	focusableWindow.WithinBox(ev.MouseX, ev.MouseY)
}

func prepareWindows() {
	border := windeau.MakeSimpleBorder('+', '|', '-')
	window = &windeau.Window{X: 0, Y: 0, Width: 30, Height: 25, Fg: termbox.ColorGreen, Bg: termbox.ColorDefault, Border: border}
	window.Title = "Hello World"
	window.Draw()

	fancyBorder := FancyBorder{TopLeft: '┏', TopRight: '┓', BottomLeft: '┗', BottomRight: '┛', Vertical: '┃', Horizontal: '━'}
	fancyWindow = &windeau.Window{X: 40, Y: 20, Width: 40, Height: 30, Fg: termbox.ColorBlue, Bg: termbox.ColorDefault, Border: fancyBorder}
	fancyWindow.Title = "Fancy Window"

	underlyingWindow := &windeau.Window{X: 40, Y: 0, Width: 15, Height: 15, Border: border}
	focusColor := windeau.WindowState{FgColor: termbox.ColorGreen, BgColor: termbox.ColorDefault}
	unfocusColor := windeau.WindowState{FgColor: termbox.ColorWhite, BgColor: termbox.ColorBlack}
	focusableWindow = &windeau.FocusableWindow{FocusOn: focusColor, FocusOff: unfocusColor, Focused: false}
	focusableWindow.SetParent(underlyingWindow)

	canvas = windeau.MakeCanvas(70, 5, 10, 10)
	canvas.Fill('x', termbox.ColorYellow, termbox.ColorYellow)
	canvas.FilledRect('y', termbox.ColorBlack, termbox.ColorYellow, windeau.Rect{75, 7, 3, 3})
	canvas.FilledRect('z', termbox.ColorBlue, termbox.ColorDefault, windeau.Rect{72, 12, 2, 3})
}
