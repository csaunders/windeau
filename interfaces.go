package windeau

type WindowBorder interface {
	Edge(x, y int, window Window) rune
	VerticalBorder() rune
	HorizontalBorder() rune
}

type Drawable interface {
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
