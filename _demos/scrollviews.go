package main

import (
	"github.com/csaunders/windeau"
	"github.com/csaunders/windeau/demos"
	"github.com/nsf/termbox-go"
)

var focusWindow *windeau.FocusableWindow
var scrollview *windeau.Scrollview
var handler *ScrollviewHandler

type ScrollviewHandler struct {
	windeau.EventHandler
	actualPosition int
}

func (h *ScrollviewHandler) OnFocus(context windeau.Event) {
	scrollview.Position = h.actualPosition
}

func (h *ScrollviewHandler) OnUnfocus(context windeau.Event) {
	scrollview.Position = -1
}

func setup(ev *termbox.Event) {
	on := windeau.WindowState{termbox.ColorGreen, termbox.ColorDefault}
	off := windeau.WindowState{termbox.ColorWhite, termbox.ColorDefault}
	border := windeau.MakeSimpleBorder('+', '|', '-')
	focusWindow = windeau.MakeFocusableWindow(20, 5, 20, 5, on, off, border)
	focusWindow.Parent.Title = "Pokemon"
	entries := []string{"Pikachu", "Charmander", "Bulbasaur", "Squirtle", "Meowth", "Pidgey", "Vulpix", "Golem"}
	handler = &ScrollviewHandler{actualPosition: 0}
	scrollview = windeau.MakeScrollview(focusWindow, entries, handler)
}

func draw(ev *termbox.Event) {
	windeau.DrawString("Click on the window labeled 'Pokemon' to interact with the scrollview", 5, 1, termbox.ColorRed, termbox.ColorDefault)
	scrollview.Draw()
}

func update(ev *termbox.Event) {
	if scrollview.IsFocused() {
		switch ev.Key {
		case termbox.KeyArrowDown:
			handler.actualPosition++
		case termbox.KeyArrowUp:
			handler.actualPosition--
		}
	}
	if ev.Type == termbox.EventMouse {
		scrollview.WithinBox(ev.MouseX, ev.MouseY)
	}
	handler.actualPosition = scrollview.SetPosition(handler.actualPosition)
}

func main() {
	callbacks := map[string]func(ev *termbox.Event){
		"setup":  setup,
		"draw":   draw,
		"update": update,
	}
	demos.RunLoop(callbacks)
}
