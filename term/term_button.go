package term

type Button struct {
}

func (b *Button) Width() int {
	return 5
}

func (b *Button) Height() int {
	return 3
}

func (b *Button) Render(state *TermState) *FrameBuffer {
	out := FrameBuffer{}
	out.Make(b.Width(), b.Height())

	bg := RGB{255, 0, 0}

	// fmt.Println(state.MouseX, state.MouseY,
	// 	state.CursorX, state.CursorY)

	if CheckInside(
		state.MouseX, state.MouseY,
		state.CursorX, state.CursorY,
		b.Width(), b.Height(),
	) {
		bg = RGB{255, 255, 0}
	}

	for x := range b.Width() {
		for y := range b.Height() {
			out[y][x].has_changed = true
			out[y][x].BGColor = bg
			out[y][x].Char = " "
			out[y][x].FGColor = RGB{255, 255, 255}
		}
	}

	return &out

}

func CreateButton(s string) Button {
	out := Button{}

	return out
}
