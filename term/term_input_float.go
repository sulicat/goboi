package term

import (
	"fmt"

	"github.com/sulicat/goboi/container"
)

type InputFloat struct {
	val             float64
	val_str         string
	is_clicked_up   bool
	is_clicked_down bool
	is_clicked_text bool
	store           *container.AnyStore
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

	for x := range b.Width() {
		for y := range b.Height() {
			out[y][x].has_changed = true
			out[y][x].BGColor = bg
			out[y][x].FGColor = fg
		}
	}

	text_color := RGB{255, 255, 255}

	// top arrow
	if b.CheckInsideChar(
		0, b.Width()-1,
		state,
		offset_x, offset_y,
	) {

		out[0][b.Width()-1].FGColor = RGB{255, 0, 0}
		b.is_clicked_up = state.MouseClicked

		// bot arrow
	} else if b.CheckInsideChar(
		2, b.Width()-1,
		state,
		offset_x, offset_y,
	) {

		out[2][b.Width()-1].FGColor = RGB{255, 0, 0}
		b.is_clicked_down = state.MouseClicked

		// text
	} else if CheckInside(
		state.MouseX, state.MouseY,
		offset_x, offset_y,
		b.Width()-1, b.Height(),
	) {
		text_color = RGB{255, 0, 0}
		b.is_clicked_text = state.MouseClicked

		if b.is_clicked_text {
			b.store.Store("is_editing", true)
		}

	} else {
		if state.MouseClicked {
			b.store.Store("is_editing", false)
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
		out[1][i+1].FGColor = text_color
		out[2][i+1].Char = CharH
	}

	return &out

}

func (b *InputFloat) IsClickedUp() bool {
	return b.is_clicked_up
}

func (b *InputFloat) IsClickedDown() bool {
	return b.is_clicked_down
}

func (b *InputFloat) IsClickedText() bool {
	return b.is_clicked_text
}

func (b *InputFloat) CheckInsideChar(
	char_r int, char_c int,
	state *TermState,
	offset_x int, offset_y int,
) bool {
	if state.MouseX == char_c+offset_x && state.MouseY == char_r+offset_y {
		return true
	}
	return false
}

func CreateInputFloat(val float64, store *container.AnyStore) InputFloat {
	// we create one based on the state
	// if the widget is being edited, use memory to initialize

	if container.AnyStoreGetAs[bool](store, "is_editing") {
		val = container.AnyStoreGetAs[float64](store, "temprory_value")
	}

	out := InputFloat{val: val, store: store}
	out.val_str = fmt.Sprintf("%.3f", val)

	return out
}
