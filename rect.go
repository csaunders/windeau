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
