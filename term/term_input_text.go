package term

import (
	"github.com/sulicat/goboi/container"
)

type InputText struct {
	val          *string
	width        int
	height       int
	store        *container.AnyStore
	should_flash bool
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

	// render text
	render_text := *b.val // TODO: suli change me to pointer in future
	cursor_pos := 0
	if is_editing {
		render_text = container.AnyStoreGetAs[string](b.store, "temp_val")
		cursor_pos = b.cursor_pos()
	}

	if CheckInside(
		state.MouseX, state.MouseY,
		offset_x, offset_y,
		b.Width(), b.Height(),
	) {
		border_color = RGB{255, 0, 255}

		if state.MouseClicked {
			b.start_editing()
		}

	} else if state.MouseClicked && is_editing { // clicked outside of the box
		b.stop_editing()
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

	b.key_input(state)

	// cursor flash
	b.update_cursor_flash()

	// render the text. Start at the top left position, start walking down
	ri := 0
	for y := 1; y < b.Height()-1; y++ {
		for x := 1; x < b.Width()-1; x++ {

			if ri >= len(render_text) {
				goto done
			}

			char := render_text[ri]
			out[y][x].Char = string(char)
			out[y][x].FGColor = RGB{255, 0, 255}

			if ri == cursor_pos && b.should_flash && is_editing {
				out[y][x].Char = CharBlock

			}

			ri += 1
		}
	}

done:

	return &out
}

func (b *InputText) cursor_pos() int {
	return container.AnyStoreGetAs[int](b.store, "cursor_pos")

}

func (b *InputText) update_cursor_flash() {
	flash := container.AnyStoreGetAs[int](b.store, "cursor_flash")
	flash += 1

	if flash > 20000 {
		flash = 0
	}

	b.should_flash = false
	if flash > 12000 {
		b.should_flash = true
	}

	b.store.Store("cursor_flash", int(flash))
}
func (b *InputText) key_input(state *TermState) {
	is_editing := container.AnyStoreGetAs[bool](b.store, "is_editing")
	current_val := container.AnyStoreGetAs[string](b.store, "temp_val")

	if is_editing {
		for _, key := range state.KeysDown {

			if IsAlphaNumeric(key) {
				current_val = current_val[:b.cursor_pos()] + string(rune(key)) + current_val[b.cursor_pos():]
				b.move_cursor(+1)
			}

			if key == KeyCodeEnter {
				b.stop_editing()
			}

			if key == KeyCodeArrowLeft {
				b.move_cursor(-1)
			}

			if key == KeyCodeArrowRight {
				b.move_cursor(+1)
			}

			if key == KeyCodeArrowUp {
				b.move_cursor(-1 * (b.Width() - 2))
			}

			if key == KeyCodeArrowDown {
				b.move_cursor(+1 * (b.Width() - 2))
			}

			if key == KeyCodeDelete {
				if b.cursor_pos() >= 1 {
					current_val = current_val[:b.cursor_pos()-1] + current_val[b.cursor_pos():]
					b.move_cursor(-1)
				}

			}

			// TODO: suli ctr+left/right

			b.store.Store("temp_val", current_val)
			*b.val = current_val
		}
	}
}

func (b *InputText) move_cursor(move_dist int) {
	cursor_pos := b.cursor_pos()
	cursor_pos += move_dist
	if cursor_pos >= 0 && cursor_pos < len(*b.val) {
		b.store.Store("cursor_pos", cursor_pos)
	}
	b.should_flash = true
}

func (b *InputText) start_editing() {
	b.store.Store("is_editing", true)
	// the value is the current val
	b.store.Store("temp_val", *b.val)
	b.store.Store("cursor_pos", 0)

}

func (b *InputText) stop_editing() {
	b.store.Store("is_editing", false)
}

func CreateInputText(s *string, width int, height int, store *container.AnyStore) InputText {
	out := InputText{val: s, width: width, height: height, store: store}
	return out
}
