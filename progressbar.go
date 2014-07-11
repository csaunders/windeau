package windeau

import (
	"github.com/nsf/termbox-go"
)

type Progressbar struct {
	Drawable
	Progress, Total int
	Complete, Empty rune
	Fg, Bg          termbox.Attribute
	Parent          Drawable
}

func NewSimpleProgressbar(total int) *Progressbar {
	return NewProgressbar(total, '=', '_')
}

func NewProgressbar(total int, complete, empty rune) *Progressbar {
	return &Progressbar{
		Progress: 0,
		Total:    total,
		Complete: complete,
		Empty:    empty,
		Fg:       termbox.ColorGreen,
		Bg:       termbox.ColorDefault,
	}
}

func (p *Progressbar) SetParent(d Drawable) {
	p.Parent = d
}

func (p *Progressbar) Tick(amount int) {
	p.Progress += amount
}

func (p *Progressbar) Draw() {
	if p.Parent != nil {
		p.Parent.Draw()
		r := p.Parent.GetRect()
		width := r.Width
		tickSize := float64(width) / float64(p.Total)
		ticks := int(tickSize * float64(p.Progress))
		for i := 0; i < width; i++ {
			symbol := p.Empty
			if i < ticks {
				symbol = p.Complete
			}
			termbox.SetCell(r.X+i, r.Y, symbol, p.Fg, p.Bg)
		}
	}
}
