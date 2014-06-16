package windeau

type Rect struct {
	X, Y, Width, Height int
}

func (r Rect) WithinRect(x, y int) bool {
	if x >= r.X && x < r.X+r.Width {
		if y >= r.Y && y < r.Y+r.Height {
			return true
		}
	}
	return false
}
