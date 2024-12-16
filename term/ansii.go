package term

import (
	"fmt"

	"github.com/sulicat/goboi/colors"
)

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
