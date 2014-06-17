package demos

import (
	"github.com/nsf/termbox-go"
)

func RunLoop(callbacks map[string]func(ev *termbox.Event)) {
	termbox.Init()
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputMouse)
	if setup := callbacks["setup"]; setup != nil {
		setup(nil)
	}

	for true {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		callbacks["draw"](nil)
		termbox.Flush()

		event := termbox.PollEvent()
		switch {
		case event.Key == termbox.KeyEsc:
			break
		default:
			callbacks["update"](&event)
		}
	}

	if teardown := callbacks["teardown"]; teardown != nil {
		teardown(nil)
	}
}
