package demos

import (
	"github.com/nsf/termbox-go"
)

func RunLoop(callbacks map[string]func(ev *termbox.Event)) {
	termbox.Init()
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputMouse)

	draw := func() {
		callbacks["draw"](nil)
		termbox.Flush()
	}

	if setup := callbacks["setup"]; setup != nil {
		setup(nil)
		draw()
	}

	done := make(chan bool)
	handleEvent := func(event termbox.Event) {
		switch {
		case event.Key == termbox.KeyEsc:
			done <- true
		default:
			callbacks["update"](&event)
		}
	}

	events := make(chan termbox.Event)
	go func() {
		events <- termbox.PollEvent()
	}()

	for true {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		var event termbox.Event
		select {
		case event = <-events:
			termbox.SetCell(0, 10, 'X', termbox.ColorRed, termbox.ColorRed)
			handleEvent(event)
		case <-done:
			break
		default:
			// Just keep going
		}
		draw()
	}

	if teardown := callbacks["teardown"]; teardown != nil {
		teardown(nil)
	}
}
