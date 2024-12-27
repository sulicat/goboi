package term

import (
	"runtime"
)

type Renderable interface {
	Width() int
	Height() int
	Render() *FrameBuffer
}

type State map[string]any // key val store for a state, 1 per widet

func GET_ID() int {
	pc, _, _, _ := runtime.Caller(2) // use the pprogram counter as the ID
	return int(pc)
}

// SomeText
func (t *Term) Label(s string) {

	// create a label and overlay its framebuffer into the main one
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()

	l := CreateLabel(s)
	l_buff := l.Render(&t.term_state)

	t.front.Overlay(
		l_buff,
		draw_pos_x, draw_pos_y)

	t.term_state.update_cursor_pos(l.Width(), l.Height())
}

// ┌─────────┐
// │ ClickMe │
// └─────────┘
func (t *Term) Button(s string) bool {

	// get the state for this button, whether we are hovering or something of the like
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()

	b := CreateButton(s)
	b_buff := b.Render(
		&t.term_state,
		draw_pos_x, draw_pos_y,
	)

	t.front.Overlay(
		b_buff,
		draw_pos_x, draw_pos_y)
	t.term_state.update_cursor_pos(b.Width(), b.Height())

	return b.IsClicked(&t.term_state, draw_pos_x, draw_pos_y)
}

// ○ -----
func (t *Term) CheckBox(s string, checked *bool) {
	// get the state for this button, whether we are hovering or something of the like
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()

	b := CreateCheckBox(s, *checked)
	b_buff := b.Render(
		&t.term_state,
		draw_pos_x, draw_pos_y,
	)

	t.front.Overlay(
		b_buff,
		draw_pos_x, draw_pos_y)
	t.term_state.update_cursor_pos(b.Width(), b.Height())

	if b.IsClicked(&t.term_state, draw_pos_x, draw_pos_y) {
		*checked = !*checked
	}
}

// ┌───────↑
// │500.123▕
// └───────↓
func (t *Term) InputFloat(val *float64) {

	id := GET_ID()
	state, has_state := t.WidgetStates[id]
	if !has_state {
		state = &State{}
		t.WidgetStates[id] = state
	}

	// get the state for this button, whether we are hovering or something of the like
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()

	b := CreateInputFloat(*val, state)
	b_buff := b.Render(
		&t.term_state,
		draw_pos_x, draw_pos_y,
	)

	t.front.Overlay(
		b_buff,
		draw_pos_x, draw_pos_y)
	t.term_state.update_cursor_pos(b.Width(), b.Height())

	if b.IsClickedUp() {
		*val += 0.1
	}

	if b.IsClickedDown() {
		*val -= 0.1
	}

	if b.IsClickedText() {
		// Set state to editing, for this element
		// state management happens here
	}
}
