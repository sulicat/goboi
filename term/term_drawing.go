package term

import "fmt"

type Color [3]int

func (t *Term) FillRect(x_in int, y_in int, width_in int, height_in int, color Color) {

	t.sb.Reset()
	// move the cursor
	t.sb.WriteString(MoveCursor(t.start_x+x_in, t.start_y+y_in))

	for y := y_in; y < y_in+height_in; y++ {
		for x := x_in; x < x_in+width_in; x++ {
			t.sb.WriteString(DrawBlock(color[0], color[1], color[2]))
		}
		t.sb.WriteString(MoveCursor(t.start_x+x_in, t.start_y+y+1))
	}

	t.writer.Write([]byte(t.sb.String()))
	fmt.Fprint(t.writer, t.sb.String())
	t.writer.Flush()

}
