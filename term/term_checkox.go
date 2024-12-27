package term

type CheckBox struct {
	val     string
	checked bool
}

func (b *CheckBox) Width() int {
	return len(b.val) + 2
}

func (b *CheckBox) Height() int {
	return 1
}

// offset is to handle input, it assumes it will be overlayed
// at offset_x, offset_y
func (b *CheckBox) Render(
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

	for x := 2; x < b.Width(); x++ {
		out[0][x].Char = string(b.val[x-2])
	}

	if b.checked {
		out[0][0].Char = CharFilled
	} else {
		out[0][0].Char = CharCircle
	}

	return &out

}

func (b *CheckBox) IsClicked(
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

func CreateCheckBox(s string, checked bool) CheckBox {
	out := CheckBox{val: s, checked: checked}

	return out
}
