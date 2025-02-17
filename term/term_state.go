package term

import (
	"slices"
)

// TODO: suli, things like 2 row alignment, grid, etc etc
const (
	AlignmentNormal int = iota
)

type TermState struct {
	// render state
	MouseX       int
	MouseY       int
	MouseDown    bool // if mouse is held this is held true
	MouseClicked bool // if mouse is held, this is triggered true for 1 step. Rising edge of mouse down
	IsScrollUp   bool
	IsScrollDown bool

	KeysDown    []int // keys held down on this step
	KeysClicked []int // keys rising edge this step

	alignment_mode  int
	cursor_x        int
	cursor_y        int
	cursor_x_prev   int
	cursor_y_prev   int
	last_drawn_w    int
	last_drawn_h    int
	absolute_next   bool
	absolute_x      int
	absolute_y      int
	last_mouse_down bool
	scroll          int

	color_scheme      ColorScheme
	color_scheme_orig ColorScheme // used for scheme reset

	tab_group_stack []int // ID of tab groups
}

func (ts *TermState) reset_cursor_pos() {
	ts.cursor_x = 0 // reset x to 0
	ts.cursor_y = 0
	ts.cursor_x_prev = ts.cursor_x
	ts.cursor_y_prev = ts.cursor_x
}

func (ts *TermState) get_cursor_pos() (int, int) {

	if ts.absolute_next {
		return ts.absolute_x, ts.absolute_y + ts.scroll
	}

	return ts.cursor_x, ts.cursor_y + ts.scroll
}

// this indicates the user wants the next draw to happen at
// the same start position as the prev draw + it's width
func (ts *TermState) SameLine() {
	ts.cursor_x = ts.cursor_x_prev + ts.last_drawn_w
	ts.cursor_y = ts.cursor_y_prev
}

// This is the user saying they want the next thing to draw at this position
// after drawing at that position, revert back to the old position
func (ts *TermState) AbsolutePosition(x int, y int) {
	ts.absolute_next = true
	ts.absolute_x = x
	ts.absolute_y = y
}

func (ts *TermState) update_cursor_pos(added_w int, added_h int) {

	if ts.absolute_next {
		ts.absolute_next = false
		return
	}

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

func (ts *TermState) NewKey(code int) {
	// if they key is not in the down buffer add it
	if !slices.Contains(ts.KeysDown, code) {
		ts.KeysDown = append(ts.KeysDown, code)
	}
}

func (ts *TermState) Step() {
	ts.KeysClicked = []int{}
	ts.KeysDown = []int{}

	StepWidgets()
}
