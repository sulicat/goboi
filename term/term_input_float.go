package term

import (
	"fmt"
	"strconv"

	"github.com/sulicat/goboi/container"
	"github.com/sulicat/goboi/regex_utils"
)

type InputFloat struct {
	val             *float64
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

	bg := RGB{0, 0, 0}
	fg := RGB{255, 255, 255}

	// if we are in edit mode and the user presses keys,
	// add the keys to the float
	is_editing := container.AnyStoreGetAs[bool](b.store, "is_editing")

	if is_editing {
		temp_val := container.AnyStoreGetAs[string](b.store, "temp_val")
		for _, key := range state.KeysDown {
			new_val := temp_val
			if (key >= 48 && key <= 57) ||
				key == KeyCodePeriod ||
				key == KeyCodeMinus {

				new_val += string(rune(key))
			}

			if key == KeyCodeDelete {
				if len(new_val) >= 1 {
					new_val = new_val[:len(new_val)-1]
				}
			}

			if key == KeyCodeEnter {
				b.store.Store("temp_val", temp_val)
				b.val_str = temp_val
				b.stop_editing()
			}

			// regex match for float
			// if pass set the temp val
			// TODO: suli, careful might be slow, check later

			matches := regex_utils.FloatingPointRE.Find([]byte(new_val))
			if len(matches) == len(new_val) || new_val == "-" {
				temp_val = new_val
			}
		}

		b.store.Store("temp_val", temp_val)
		b.val_str = temp_val

	} else {
		b.val_str = fmt.Sprintf("%.3f", *b.val)
	}

	out := FrameBuffer{}
	out.Make(b.Width(), b.Height())
	for x := range b.Width() {
		for y := range b.Height() {
			out[y][x].has_changed = true
			out[y][x].BGColor = bg
			out[y][x].FGColor = fg
		}
	}

	text_color := RGB{255, 255, 255}
	if is_editing {
		text_color = RGB{255, 0, 255}
	}

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
			b.start_editing()
		}

	} else {
		if state.MouseClicked && is_editing {
			b.stop_editing()
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

func (b *InputFloat) start_editing() {
	b.store.Store("is_editing", true)
	b.store.Store("temp_val", "")
}

func (b *InputFloat) stop_editing() {
	b.store.Store("is_editing", false)

	*b.val, _ = strconv.ParseFloat(b.val_str, 64)
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

func CreateInputFloat(val *float64, store *container.AnyStore) InputFloat {
	// we create one based on the state
	// if the widget is being edited, use memory to initialize

	out := InputFloat{val: val, store: store}
	out.val_str = fmt.Sprintf("%.5f", *val)

	return out
}
