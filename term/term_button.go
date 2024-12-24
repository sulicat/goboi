package term

type Button struct {
}

func (b *Button) Width() int {
	return 10
}

func (b *Button) Height() int {
	return 3
}

func (b *Button) Render(state *TermState) *FrameBuffer {
	out := FrameBuffer{}
	out.Make(b.Width(), b.Height())

	for x := range b.Width() {
		for y := range b.Height() {
			out[y][x].has_changed = true
			out[y][x].BGColor = RGB{0, 255, 255}
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
