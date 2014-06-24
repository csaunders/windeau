package windeau

type Rect struct {
	X, Y, Width, Height int
}

func (r Rect) WithinRect(x, y int) bool {
	return x >= r.X && x < r.X+r.Width && y >= r.Y && y < r.Y+r.Height
}

func (r Rect) Contains(other Rect) bool {
	if r.X <= other.X && r.X+r.Width >= other.X+other.Width {
		if r.Y <= other.Y && r.Y+r.Height >= other.Y+other.Height {
			return true
		}
	}
	return false
}

func (r Rect) DoesNotContain(other Rect) bool {
	return !r.Contains(other)
}

func (r Rect) ShrinkBy(amount int) Rect {
	return Rect{r.X + amount, r.Y + amount, r.Width - amount, r.Height - amount}
}
