package windeau

import "github.com/nsf/termbox-go"

type WindowState struct {
	FgColor termbox.Attribute
	BgColor termbox.Attribute
}

type FocusableWindow struct {
	WindowImpl        *Window
	FocusOn, FocusOff WindowState
	Focused           bool
}

func (c *FocusableWindow) ToggleFocus() {
	c.Focused = !c.Focused
}

func (c *FocusableWindow) Draw() {
	c.setColors()
	c.WindowImpl.Draw()
}

func (c *FocusableWindow) WithinBox(x, y int) bool {
	c.Focused = c.WindowImpl.WithinBox(x, y)
	return c.Focused
}

func (c *FocusableWindow) setColors() {
	if c.Focused {
		c.WindowImpl.Fg = c.FocusOn.FgColor
		c.WindowImpl.Bg = c.FocusOn.BgColor
	} else {
		c.WindowImpl.Fg = c.FocusOff.FgColor
		c.WindowImpl.Bg = c.FocusOff.BgColor
	}
}
