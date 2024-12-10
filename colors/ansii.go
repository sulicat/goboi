package colors

import "fmt"

func MoveCursor(x int, y int) string {
	return fmt.Sprintf("\033[%d;%dH", y, x)
}

func Clear() string {
	return "\033[2J"
}

func DrawBlock(r int, g int, b int) string {
	return fmt.Sprintf(Color(r, g, b) + "█" + Reset)
}

func DrawChar(char string, r int, g int, b int) string {
	return fmt.Sprintf(Color(r, g, b) + char + Reset)
}

func DrawBlank() string {
	return " "
}
