package term

type Renderable interface {
	Width() int
	Height() int
	Render() *FrameBuffer
}

func (t *Term) Label(s string) {
	// create a label and overlay its framebuffer into the main one
	l := CreateLabel(s)
	l_buff := l.Render(&t.term_state)
	t.front.Overlay(l_buff, t.cursor_x, t.cursor_y)

	if t.term_state.same_line {
		t.term_state.same_line = false
		t.cursor_x += l.Width()
	} else {
		t.cursor_y += l.Height()
	}
}

func (t *Term) Button(s string) bool {

	// get the state for this button, whether we are hovering or something of the like
	b := CreateButton(s)
	b_buff := b.Render(&t.term_state)
	t.front.Overlay(b_buff, t.cursor_x, t.cursor_y)

	if t.term_state.same_line {
		t.term_state.same_line = false
		t.cursor_x += b.Width()
	} else {
		t.cursor_y += b.Height()
	}

	return true
}
