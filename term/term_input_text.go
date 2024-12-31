package term

import (
	"github.com/sulicat/goboi/container"
)

type InputText struct {
	val    *string
	width  int
	height int
	store  *container.AnyStore
}

func (b *InputText) Width() int {
	return b.width
}

func (b *InputText) Height() int {
	return b.height
}

// offset is to handle input, it assumes it will be overlayed
// at offset_x, offset_y
func (b *InputText) Render(
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

	border_color := RGB{255, 255, 255}
	is_editing := container.AnyStoreGetAs[bool](b.store, "is_editing")

	if CheckInside(
		state.MouseX, state.MouseY,
		offset_x, offset_y,
		b.Width(), b.Height(),
	) {
		border_color = RGB{255, 255, 0}

		if state.MouseClicked {
			b.start_editing()
		}

	}

	out[0][0].FGColor = border_color
	out[0][b.Width()-1].FGColor = border_color
	out[b.Height()-1][0].FGColor = border_color
	out[b.Height()-1][b.Width()-1].FGColor = border_color

	// draw the border
	if is_editing {
		out[0][0].Char = CharDTL
		out[0][b.Width()-1].Char = CharDTR
		out[b.Height()-1][0].Char = CharDBL
		out[b.Height()-1][b.Width()-1].Char = CharDBR
	} else {
		out[0][0].Char = CharTL
		out[0][b.Width()-1].Char = CharTR
		out[b.Height()-1][0].Char = CharBL
		out[b.Height()-1][b.Width()-1].Char = CharBR
	}

	border_v := CharV
	border_h := CharH

	if is_editing {
		border_h = CharDH
		border_v = CharDV
	}

	for x := 1; x < b.Width()-1; x++ {
		out[0][x].Char = border_h
		out[0][x].FGColor = border_color
		out[b.Height()-1][x].Char = border_h
		out[b.Height()-1][x].FGColor = border_color
	}

	for y := 1; y < b.Height()-1; y++ {
		out[y][0].Char = border_v
		out[y][0].FGColor = border_color
		out[y][b.Width()-1].Char = border_v
		out[y][b.Width()-1].FGColor = border_color
	}

	return &out

}

func (b *InputText) start_editing() {
	b.store.Store("is_editing", true)
}

func (b *InputText) stop_editing() {
	b.store.Store("is_editing", false)
}

func CreateInputText(s *string, width int, height int, store *container.AnyStore) InputText {
	out := InputText{val: s, width: width, height: height, store: store}

	return out
}
