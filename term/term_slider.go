package term

import (
	"fmt"
	"math"

	"github.com/sulicat/goboi/container"
)

type Slider struct {
	store *container.AnyStore
	val   *float64
	min   float64
	max   float64
}

func (b *Slider) Width() int {
	return 12 + 8
}

func (b *Slider) Height() int {
	return 1
}

// offset is to handle input, it assumes it will be overlayed
// at offset_x, offset_y
func (b *Slider) Render(
	state *TermState,
	offset_x int, offset_y int,
) *FrameBuffer {

	bg := state.color_scheme.BackgroundColor
	fg := state.color_scheme.TextColor
	text_color := fg

	out := FrameBuffer{}
	out.Make(b.Width(), b.Height())
	for x := range b.Width() {
		for y := range b.Height() {
			out[y][x].has_changed = true
			out[y][x].BGColor = bg
			out[y][x].FGColor = fg
		}
	}

	slider_width := b.Width() - 8
	slider_pos := (*b.val - b.min) / (b.max - b.min)
	slider_pos = math.Max(0, slider_pos)
	slider_pos = math.Min(1, slider_pos)
	slider_index := int(slider_pos * float64(slider_width))

	if b.CheckInsideChar(
		0, slider_index,
		state,
		offset_x, offset_y,
	) {

		if state.MouseDown {
			out[0][slider_index].FGColor = state.color_scheme.SelectedColor
			text_color = state.color_scheme.SelectedColor
		} else {
			out[0][slider_index].FGColor = state.color_scheme.HoverColor
			text_color = state.color_scheme.HoverColor
		}

	}

	// if the mouse is inside the slider anywhere
	// and the mouse is down, set the position of the slider to that
	// and the float value to that
	if CheckInside(
		state.MouseX, state.MouseY,
		offset_x, offset_y,
		slider_width, 1,
	) {
		if state.MouseDown {
			slider_index = state.MouseX - offset_x
			percent := float64(slider_index) / float64(slider_width)
			new_val := (percent * (b.max - b.min)) + b.min
			*b.val = new_val
		}
	}

	for x := range slider_width {
		if x == slider_index {
			out[0][x].Char = CharBlock

		} else {

			out[0][x].Char = CharH
		}
	}

	val_string := fmt.Sprintf("%f", *b.val)
	for i, c := range val_string {
		if i < 5 {
			out[0][slider_width+i+1].Char = string(c)
			out[0][slider_width+i+1].FGColor = text_color
		}
	}

	return &out
}

func (b *Slider) CheckInsideChar(
	char_r int, char_c int,
	state *TermState,
	offset_x int, offset_y int,
) bool {
	if state.MouseX == char_c+offset_x && state.MouseY == char_r+offset_y {
		return true
	}
	return false
}

func CreateSlider(
	val *float64,
	min float64,
	max float64,
	store *container.AnyStore) Slider {

	out := Slider{
		store: store,
		min:   min,
		max:   max,
		val:   val}
	return out
}
