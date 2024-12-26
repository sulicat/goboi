package term

type Renderable interface {
	Width() int
	Height() int
	Render() *FrameBuffer
}

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

	if t.term_state.MouseClicked {
		if CheckInside(
			t.term_state.MouseX, t.term_state.MouseY,
			draw_pos_x, draw_pos_y,
			b.Width(), b.Height(),
		) {
			return true
		}
	}

	return false
}

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

	if t.term_state.MouseClicked {
		if CheckInside(
			t.term_state.MouseX, t.term_state.MouseY,
			draw_pos_x, draw_pos_y,
			b.Width(), b.Height(),
		) {
			*checked = !*checked
		}
	}
}

func (t *Term) InputFloat(val *float64) {

	// get the state for this button, whether we are hovering or something of the like
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()

	b := CreateInputFloat(*val)
	b_buff := b.Render(
		&t.term_state,
		draw_pos_x, draw_pos_y,
	)

	t.front.Overlay(
		b_buff,
		draw_pos_x, draw_pos_y)
	t.term_state.update_cursor_pos(b.Width(), b.Height())
}
