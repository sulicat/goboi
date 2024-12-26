package term

import "fmt"

type InputFloat struct {
	val     float64
	val_str string
}

func (b *InputFloat) Width() int {
	l := len(b.val_str)
	return l + 2
}

func (b *InputFloat) Height() int {
	return 3
}

// offset is to handle input, it assumes it will be overlayed
// at offset_x, offset_y
func (b *InputFloat) Render(
	state *TermState,
	offset_x int, offset_y int,
) *FrameBuffer {

	out := FrameBuffer{}
	out.Make(b.Width(), b.Height())

	bg := RGB{0, 0, 0}
	fg := RGB{255, 255, 255}

	// if CheckInside(
	// 	state.MouseX, state.MouseY,
	// 	offset_x, offset_y,
	// 	b.Width(), b.Height(),
	// ) {
	// 	bg = RGB{0, 0, 0}
	// 	fg = RGB{255, 0, 0}
	// }

	for x := range b.Width() {
		for y := range b.Height() {
			out[y][x].has_changed = true
			out[y][x].BGColor = bg
			out[y][x].FGColor = fg
		}
	}

	out[0][0].Char = CharTL
	out[1][0].Char = CharV
	out[2][0].Char = CharBL

	out[0][b.Width()-1].Char = CharUp
	out[1][b.Width()-1].Char = CharBlockR
	out[2][b.Width()-1].Char = CharDown

	for i, c := range b.val_str {
		out[0][i+1].Char = CharH
		out[1][i+1].Char = string(c)
		out[2][i+1].Char = CharH
	}

	return &out

}

func CreateInputFloat(val float64) InputFloat {
	out := InputFloat{val: val}
	out.val_str = fmt.Sprintf("%.3f", val)

	return out
}
