package windeau

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type WindowBorder interface {
	Edge(x, y int, window Window) rune
	VerticalBorder() rune
	HorizontalBorder() rune
}

type Drawable interface {
	SetParent(parent Drawable)
	IsRoot() bool
	IsFocused() bool
	GetRect() Rect
	WithinBox(x, y int) bool
	Draw()
}

type DataSource interface {
	GetEntries() []fmt.Stringer
	GetPosition() int
}

type Event struct {
	termbox.Event
}

type EventHandler interface {
	OnFocus(context Event)
	OnUnfocus(context Event)
	OnHighlight(context Event)
	OnSelect(context Event)
}
