package term

import (
	"fmt"

	"github.com/sulicat/goboi/colors"
)

const (
	// Arrow Keys
	KeyArrowUp    = "\x1b[A"
	KeyArrowDown  = "\x1b[B"
	KeyArrowRight = "\x1b[C"
	KeyArrowLeft  = "\x1b[D"

	// Function Keys
	KeyF1  = "\x1bOP"
	KeyF2  = "\x1bOQ"
	KeyF3  = "\x1bOR"
	KeyF4  = "\x1bOS"
	KeyF5  = "\x1b[15~"
	KeyF6  = "\x1b[17~"
	KeyF7  = "\x1b[18~"
	KeyF8  = "\x1b[19~"
	KeyF9  = "\x1b[20~"
	KeyF10 = "\x1b[21~"
	KeyF11 = "\x1b[23~"
	KeyF12 = "\x1b[24~"

	// Home, End, Insert, Delete, Page Up, Page Down
	KeyHome     = "\x1b[H"
	KeyEnd      = "\x1b[F"
	KeyInsert   = "\x1b[2~"
	KeyDelete   = "\x1b[3~"
	KeyPageUp   = "\x1b[5~"
	KeyPageDown = "\x1b[6~"

	// Control Keys
	KeyCtrlA  = "\x01"
	KeyCtrlB  = "\x02"
	KeyCtrlC  = "\x03"
	KeyCtrlD  = "\x04"
	KeyCtrlE  = "\x05"
	KeyCtrlF  = "\x06"
	KeyCtrlG  = "\x07"
	KeyCtrlH  = "\x08" // Backspace
	KeyCtrlI  = "\x09" // Tab
	KeyCtrlJ  = "\x0A" // Line Feed
	KeyCtrlK  = "\x0B"
	KeyCtrlL  = "\x0C"
	KeyCtrlM  = "\x0D" // Carriage Return (Enter)
	KeyCtrlN  = "\x0E"
	KeyCtrlO  = "\x0F"
	KeyCtrlP  = "\x10"
	KeyCtrlQ  = "\x11"
	KeyCtrlR  = "\x12"
	KeyCtrlS  = "\x13"
	KeyCtrlT  = "\x14"
	KeyCtrlU  = "\x15"
	KeyCtrlV  = "\x16"
	KeyCtrlW  = "\x17"
	KeyCtrlX  = "\x18"
	KeyCtrlY  = "\x19"
	KeyCtrlZ  = "\x1A"
	KeyEscape = "\x1b" // ESC Key

	// Mouse Input Prefix
	MouseInputPrefix = "\x1b[M"
)

func IsAlphaNumeric(keycode int) bool {
	return (keycode >= 'A' && keycode <= 'Z') ||
		(keycode >= 'a' && keycode <= 'z') ||
		(keycode >= '0' && keycode <= '9') ||
		keycode == ',' ||
		keycode == '.' ||
		keycode == ' ' ||
		keycode == '!' ||
		keycode == '?'
}

func MoveCursor(x int, y int) string {
	return fmt.Sprintf("\033[%d;%dH", y, x)
}

func Clear() string {
	return "\033[2J"
}

func DrawBlock(r int, g int, b int) string {
	return fmt.Sprintf(colors.Color(r, g, b) + "█" + colors.Reset)
}

func DrawChar(char string, r int, g int, b int) string {
	return fmt.Sprintf(colors.Color(r, g, b) + char + colors.Reset)
}

func DrawBlank() string {
	return " "
}

func EnableMouseTracking() string {
	return "\x1b[?1003h"
}

func DisableMouseTracking() string {
	return "\x1b[?1003l"
}
