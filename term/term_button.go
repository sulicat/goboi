package term

type Button struct {
	val string
}

func (b *Button) Width() int {
	return len(b.val) + 2
}

func (b *Button) Height() int {
	return 3
}

// offset is to handle input, it assumes it will be overlayed
// at offset_x, offset_y
func (b *Button) Render(
	state *TermState,
	offset_x int, offset_y int,
) *FrameBuffer {

	out := FrameBuffer{}
	out.Make(b.Width(), b.Height())

	bg := RGB{0, 0, 0}
	fg := RGB{255, 255, 255}

	if CheckInside(
		state.MouseX, state.MouseY,
		offset_x, offset_y,
		b.Width(), b.Height(),
	) {
		bg = RGB{0, 0, 0}
		fg = RGB{255, 0, 0}
	}

	for x := range b.Width() {
		for y := range b.Height() {
			out[y][x].has_changed = true
			out[y][x].BGColor = bg
			out[y][x].FGColor = fg
		}
	}

	out[0][0].Char = CharTL
	out[0][b.Width()-1].Char = CharTR

	out[2][0].Char = CharBL
	out[2][b.Width()-1].Char = CharBR

	out[1][0].Char = CharV
	out[1][b.Width()-1].Char = CharV

	for x := 1; x < b.Width()-1; x++ {
		out[0][x].Char = CharH
		out[1][x].Char = string(b.val[x-1])
		out[2][x].Char = CharH
	}

	return &out

}

func (b *Button) IsClicked(
	state *TermState,
	offset_x int, offset_y int,
) bool {

	if state.MouseClicked {
		if CheckInside(
			state.MouseX, state.MouseY,
			offset_x, offset_y,
			b.Width(), b.Height(),
		) {
			return true
		}
	}
	return false
}

func CreateButton(s string) Button {
	out := Button{val: " " + s + " "}

	return out
}
