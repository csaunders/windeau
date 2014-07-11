package main

import (
	"github.com/csaunders/windeau"
	"github.com/csaunders/windeau/demos"
	"github.com/nsf/termbox-go"
	"time"
)

var window *windeau.Window
var progress *windeau.Progressbar
var reset, done chan (bool)

func setup(ev *termbox.Event) {
	border := windeau.MakeSimpleBorder('+', '|', '-')
	window = &windeau.Window{
		0, 0, 80, 5,
		border,
		"Hello World",
		termbox.ColorGreen,
		termbox.ColorDefault,
		nil,
	}
	progress = windeau.NewSimpleProgressbar(1000)
	progress.Fg = termbox.ColorYellow
	progress.SetParent(window)
	go func() {
		for true {
			select {
			case <-reset:
				progress.Progress = 0
			case <-done:
				break
			default:
				progress.Tick(10)
			}
			time.Sleep(1 * time.Millisecond)
		}
	}()
}

func draw(ev *termbox.Event) {
	progress.Draw()
}

func update(ev *termbox.Event) {
	if ev.Type == termbox.EventMouse && window.WithinBox(ev.MouseX, ev.MouseY) {
		progress.Progress = 0
	}
}

func main() {
	callbacks := map[string]func(ev *termbox.Event){
		"setup":  setup,
		"draw":   draw,
		"update": update,
	}
	demos.RunLoop(callbacks)
}
