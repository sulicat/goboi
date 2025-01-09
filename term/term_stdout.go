package term

import (
	"github.com/sulicat/goboi/container"
)

type STDOUT struct {
	width  int
	height int
	data   *string
	store  *container.AnyStore
}

func (b *STDOUT) Width() int {
	return b.width
}

func (b *STDOUT) Height() int {
	return b.height
}

// offset is to handle input, it assumes it will be overlayed
// at offset_x, offset_y
func (b *STDOUT) Render(
	state *TermState,
	offset_x int, offset_y int,
) *FrameBuffer {

	out := FrameBuffer{}
	out.Make(b.Width(), b.Height())

	// CACHE buster
	for x := range b.Width() {
		for y := range b.Height() {
			out[y][x].has_changed = true
		}
	}

	border_color := state.color_scheme.TextColor

	// border
	out[0][0].Char = CharTL
	out[0][b.Width()-1].Char = CharTR
	out[b.Height()-1][0].Char = CharBL
	out[b.Height()-1][b.Width()-1].Char = CharBR

	out[0][0].FGColor = border_color
	out[0][b.Width()-1].FGColor = border_color
	out[b.Height()-1][0].FGColor = border_color
	out[b.Height()-1][b.Width()-1].FGColor = border_color

	for x := 1; x < b.Width()-1; x++ {
		out[0][x].Char = "-"
		out[0][x].FGColor = border_color
		out[b.Height()-1][x].Char = "-"
		out[b.Height()-1][x].FGColor = border_color
	}

	// scroll bar on the right side

	return &out
}

func CreateSTDOUT(width int, height int, store *container.AnyStore) STDOUT {
	out := STDOUT{width: width, height: height, store: store}
	return out
}
