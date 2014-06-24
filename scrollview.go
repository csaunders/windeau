package windeau

import (
	"github.com/nsf/termbox-go"
)

const (
	SVMoreUp   string = "/\\more/\\"
	SVMoreDown string = "\\/more\\/"
)

type Scrollview struct {
	Parent          Drawable
	Entries         []string
	Position        int
	Handler         EventHandler
	visibleRowCount int
}

func MakeScrollview(parent *FocusableWindow, entries []string, handler EventHandler) *Scrollview {
	scrollview := &Scrollview{Entries: entries, Position: -1, Handler: handler}
	scrollview.SetParent(parent)
	return scrollview
}

func (s *Scrollview) SetParent(parent Drawable) {
	s.Parent = parent
	s.determineVisibleRowCount()
}

func (s *Scrollview) IsRoot() bool {
	return false
}

func (s *Scrollview) IsFocused() bool {
	return s.Parent.IsFocused()
}

func (s *Scrollview) GetRect() Rect {
	return s.Parent.GetRect().ShrinkBy(1)
}

func (s *Scrollview) WithinBox(x, y int) bool {
	return s.Parent.WithinBox(x, y)
}

func (s *Scrollview) GetEntries() []string {
	return s.Entries
}

func (s *Scrollview) GetPosition() int {
	return s.Position
}

func (s *Scrollview) SetPosition(position int) int {
	if position >= 0 && position < len(s.GetEntries()) {
		s.Position = position
	}
	return s.Position
}

func (s *Scrollview) SetHandler(handler EventHandler) {
	s.Handler = handler
}

func (s *Scrollview) Draw() {
	if s.Parent == nil {
		panic("Parent cannot be nil for a Scrollview")
	}
	s.Parent.Draw()
	rect := s.GetRect()
	for i := 0; i < s.visibleRowCount; i++ {
		entry := i + s.startingPosition()
		if i >= len(s.GetEntries()) {
			break
		}
		fg, bg := termbox.ColorWhite, termbox.ColorDefault
		if s.Position == entry && s.IsFocused() {
			bg = termbox.ColorBlue
		}

		x := rect.X
		y := rect.Y + i
		message := s.GetEntries()[entry]
		DrawStringWithinSize(message, x, y, rect.Width, fg, bg)
	}
	s.hintMoreDataExists()
}

func (s *Scrollview) startingPosition() int {
	halfway := s.visibleRowCount / 2
	numberOfEntries := len(s.GetEntries())
	switch {
	case numberOfEntries < s.visibleRowCount:
		return 0
	case s.Position < s.visibleRowCount:
		return 0
	case s.Position >= numberOfEntries-s.visibleRowCount:
		return numberOfEntries - s.visibleRowCount
	default:
		return s.Position - halfway
	}
}

func (s *Scrollview) determineVisibleRowCount() {
	bindingBox := s.GetRect()
	s.visibleRowCount = bindingBox.Height - 1
}

func (s *Scrollview) hintMoreDataExists() {
	if len(s.GetEntries()) < s.visibleRowCount {
		return
	}
	parentRect := s.Parent.GetRect()
	fg, bg := s.Parent.GetColors()
	leftEdge := parentRect.X + parentRect.Width - 2
	if s.Position >= s.visibleRowCount {
		DrawStringWithinSize(SVMoreUp, leftEdge-len(SVMoreUp), parentRect.Y, parentRect.Width-2, fg, bg)
	}
	if s.Position < len(s.GetEntries())-s.visibleRowCount {
		DrawStringWithinSize(SVMoreDown, leftEdge-len(SVMoreDown), parentRect.Y+parentRect.Height-1, parentRect.Width-2, fg, bg)
	}
}
