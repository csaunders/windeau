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

func MakeFocusableWindow(x, y, w, h int, on, off WindowState, border WindowBorder) *FocusableWindow {
	parent := &Window{X: x, Y: y, Width: w, Height: h, Border: border}
	return &FocusableWindow{parent, on, off, false}
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

func (c *FocusableWindow) IsFocused() bool {
	return c.Focused
}

func (c *FocusableWindow) GetRect() Rect {
	return c.Parent.GetRect()
}

func (c *FocusableWindow) WithinBox(x, y int) bool {
	c.Focused = c.Parent.WithinBox(x, y)
	return c.Focused
}

func (c *FocusableWindow) ActiveColors() WindowState {
	if c.Focused {
		return c.FocusOn
	} else {
		return c.FocusOff
	}
}

func (c *FocusableWindow) setColors() {
	c.Parent.Fg = c.ActiveColors().FgColor
	c.Parent.Bg = c.ActiveColors().BgColor
}
