package term

// TODO: suli, things like 2 row alignment, grid, etc etc
const (
	AlignmentNormal int = iota
)

type TermState struct {
	// render state
	bg_color       RGB
	fg_color       RGB
	alignment_mode int

	MouseX      int
	MouseY      int
	MouseButton int

	// eventually handle the other keys,
	// either [3]int or different vals
	MouseClicked bool

	cursor_x      int
	cursor_y      int
	cursor_x_prev int
	cursor_y_prev int
	last_drawn_w  int
	last_drawn_h  int
}

func (ts *TermState) reset_cursor_pos() {
	ts.cursor_x = 0 // reset x to 0
	ts.cursor_y = 0
	ts.cursor_x_prev = ts.cursor_x
	ts.cursor_y_prev = ts.cursor_x
}

func (ts *TermState) get_cursor_pos() (int, int) {
	return ts.cursor_x, ts.cursor_y
}

// this indicates the user wants the next draw to happen at
// the same start position as the prev draw + it's width
func (ts *TermState) SameLine() {
	ts.cursor_x = ts.cursor_x_prev + ts.last_drawn_w
	ts.cursor_y = ts.cursor_y_prev
}

func (ts *TermState) update_cursor_pos(added_w int, added_h int) {
	ts.cursor_x_prev = ts.cursor_x
	ts.cursor_y_prev = ts.cursor_y

	ts.last_drawn_w = added_w
	ts.last_drawn_h = added_h

	switch ts.alignment_mode {

	case AlignmentNormal:
		// in the case of 1 item per line
		ts.cursor_x = 0 // reset x to 0
		ts.cursor_y += added_h
	}

}
