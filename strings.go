package windeau

import "github.com/nsf/termbox-go"

func DrawString(s string, x, y int, fg, bg termbox.Attribute) {
	DrawStringWithinSize(s, x, y, len(s), fg, bg)
}

func DrawStringWithinSize(s string, x, y, width int, fg, bg termbox.Attribute) {
	for px, char := range ConvertToRuneArray(s) {
		if px >= width {
			break
		}
		termbox.SetCell(x+px, y, char, fg, bg)
	}
}

func ConvertToRuneArray(s string) []rune {
	runes := make([]rune, len(s))
	for px := 0; px < len(s) && px < len(s); px++ {
		runes[px] = rune(s[px])
	}
	return runes
}
