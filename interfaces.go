package windeau

type WindowBorder interface {
	Edge(x, y int, window Window) rune
	VerticalBorder() rune
	HorizontalBorder() rune
}

type Drawable interface {
	SetParent(parent Drawable)
	IsRoot() bool
	GetRect() Rect
	WithinBox(x, y int) bool
	Draw()
}

type DataSource interface {
	Entries() []string
	Position() int
}

type EventHandler interface {
	OnHighlight(context interface{})
	OnSelect(context interface{})
}
