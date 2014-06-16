package windeau

import "github.com/nsf/termbox-go"

type WindowState struct {
	FgColor termbox.Attribute
	BgColor termbox.Attribute
}

type FocusableWindow struct {
	Parent            *Window
	FocusOn, FocusOff WindowState
	Focused           bool
}

func (c *FocusableWindow) ToggleFocus() {
	c.Focused = !c.Focused
}

func (c *FocusableWindow) Draw() {
	c.setColors()
	c.Parent.Draw()
}

func (c *FocusableWindow) SetParent(parent Drawable) {
	window, ok := parent.(*Window)
	if ok {
		c.Parent = window
	}
}

func (c *FocusableWindow) IsRoot() bool {
	return false
}

func (c *FocusableWindow) GetRect() Rect {
	return c.Parent.GetRect()
}

func (c *FocusableWindow) WithinBox(x, y int) bool {
	c.Focused = c.Parent.WithinBox(x, y)
	return c.Focused
}

func (c *FocusableWindow) setColors() {
	if c.Focused {
		c.Parent.Fg = c.FocusOn.FgColor
		c.Parent.Bg = c.FocusOn.BgColor
	} else {
		c.Parent.Fg = c.FocusOff.FgColor
		c.Parent.Bg = c.FocusOff.BgColor
	}
}
