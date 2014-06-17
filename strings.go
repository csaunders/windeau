package windeau

import "github.com/nsf/termbox-go"

func DrawString(s string, x, y int, fg, bg termbox.Attribute) {
	DrawStringWithinSize(s, x, y, len(s), fg, bg)
}

func DrawStringWithinSize(s string, x, y, width int, fg, bg termbox.Attribute) {
	for px := 0; px < width && px < len(s); px++ {
		termbox.SetCell(x+px, y, rune(s[px]), fg, bg)
	}
}
