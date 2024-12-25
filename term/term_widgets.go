package term

type Renderable interface {
	Width() int
	Height() int
	Render() *FrameBuffer
}

// TODO: suli figure out the SameLine alignment properly

func (t *Term) Label(s string) {
	// create a label and overlay its framebuffer into the main one
	l := CreateLabel(s)
	l_buff := l.Render(&t.term_state)
	t.front.Overlay(
		l_buff,
		t.term_state.CursorX, t.term_state.CursorY)

	// if t.term_state.same_line {
	// 	t.term_state.same_line = false
	// 	t.term_state.CursorX += l.Width()
	// } else {
	t.term_state.CursorY += l.Height()
	t.term_state.CursorX = 0
	// }
}

func (t *Term) Button(s string) bool {

	// get the state for this button, whether we are hovering or something of the like
	b := CreateButton(s)
	b_buff := b.Render(&t.term_state)
	t.front.Overlay(
		b_buff,
		t.term_state.CursorX, t.term_state.CursorY)

	// if t.term_state.same_line {
	// 	t.term_state.same_line = false
	// 	t.term_state.CursorX += b.Width()
	// } else {
	t.term_state.CursorY += b.Height()
	t.term_state.CursorX = 0
	// }

	return true
}
