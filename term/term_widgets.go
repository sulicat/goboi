package term

import (
	"runtime"

	"github.com/sulicat/goboi/container"
)

// a map of line number -> iteration count.
// reset every step
// increment every GET_ID()
var iteration_state = map[int]int{}

type Renderable interface {
	Width() int
	Height() int
	Render() *FrameBuffer
}

func GET_LINE_NUM() int {
	pc, _, _, _ := runtime.Caller(3)
	return int(pc)
}

func StepWidgets() {
	iteration_state = map[int]int{}
}

func GetUniqueStore(t *Term) *container.AnyStore {
	line_num := GET_LINE_NUM()

	id := iteration_state[line_num] + (line_num * 1000)
	iteration_state[line_num] += 1

	store, has_store := t.WidgetStores[id]
	if !has_store {
		store = container.CreateAnyStore()
		t.WidgetStores[id] = store
	}
	return store
}

// SomeText
func (t *Term) Label(s string) {

	// create a label and overlay its framebuffer into the main one
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()

	l := CreateLabel(s)
	l_buff := l.Render(&t.term_state)

	t.front.Overlay(
		l_buff,
		draw_pos_x, draw_pos_y)

	t.term_state.update_cursor_pos(l.Width(), l.Height())
}

// ┌─────────┐
// │ ClickMe │
// └─────────┘
func (t *Term) Button(s string) bool {

	// get the state for this button, whether we are hovering or something of the like
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()

	b := CreateButton(s)
	b_buff := b.Render(
		&t.term_state,
		draw_pos_x, draw_pos_y,
	)

	t.front.Overlay(
		b_buff,
		draw_pos_x, draw_pos_y)
	t.term_state.update_cursor_pos(b.Width(), b.Height())

	return b.IsClicked(&t.term_state, draw_pos_x, draw_pos_y)
}

// ○ -----
func (t *Term) CheckBox(s string, checked *bool) {
	// get the state for this button, whether we are hovering or something of the like
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()

	b := CreateCheckBox(s, *checked)
	b_buff := b.Render(
		&t.term_state,
		draw_pos_x, draw_pos_y,
	)

	t.front.Overlay(
		b_buff,
		draw_pos_x, draw_pos_y)
	t.term_state.update_cursor_pos(b.Width(), b.Height())

	if b.IsClicked(&t.term_state, draw_pos_x, draw_pos_y) {
		*checked = !*checked
	}
}

// ┌───────↑
// │500.123▕
// └───────↓
func (t *Term) InputFloat(val *float64) {

	// get the state for this button, whether we are hovering or something of the like
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()
	store := GetUniqueStore(t)

	b := CreateInputFloat(val, store)
	b_buff := b.Render(
		&t.term_state,
		draw_pos_x, draw_pos_y,
	)

	t.front.Overlay(
		b_buff,
		draw_pos_x, draw_pos_y)
	t.term_state.update_cursor_pos(b.Width(), b.Height())

}

// ────────█────────── 100
func (t *Term) Slider(val *float64, min float64, max float64) {
	store := GetUniqueStore(t)

	// get the state for this button, whether we are hovering or something of the like
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()

	b := CreateSlider(val, min, max, store)
	b_buff := b.Render(
		&t.term_state,
		draw_pos_x, draw_pos_y,
	)

	t.front.Overlay(
		b_buff,
		draw_pos_x, draw_pos_y)
	t.term_state.update_cursor_pos(b.Width(), b.Height())
}

// ┌────────────────────────────────────────────────┐
// │hello world                                     │
// │                                                │
// └────────────────────────────────────────────────┘
func (t *Term) InputText(in *string, width int, height int) {
	store := GetUniqueStore(t)

	// get the state for this button, whether we are hovering or something of the like
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()

	b := CreateInputText(in, width, height, store)
	b_buff := b.Render(
		&t.term_state,
		draw_pos_x, draw_pos_y,
	)

	t.front.Overlay(
		b_buff,
		draw_pos_x, draw_pos_y)
	t.term_state.update_cursor_pos(b.Width(), b.Height())

}

// CANVAS
func (t *Term) CreatePixels(width int, height int) [][]RGBA {
	out := make([][]RGBA, height)
	for i := range len(out) {
		out[i] = make([]RGBA, width)
	}
	return out
}

// CANVAS
func (t *Term) Canvas(pixels *[][]RGBA) {
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()
	b := CreateCanvasType(pixels)
	b_buff := b.Render(
		&t.term_state,
		draw_pos_x, draw_pos_y,
	)

	t.front.Overlay(
		b_buff,
		draw_pos_x, draw_pos_y)
	t.term_state.update_cursor_pos(b.Width(), b.Height())

}

func (t *Term) STDOUT(width int, height int) {
	store := GetUniqueStore(t)

	// get the state for this button, whether we are hovering or something of the like
	draw_pos_x, draw_pos_y := t.term_state.get_cursor_pos()

	b := CreateSTDOUT(width, height, store)
	b_buff := b.Render(
		&t.term_state,
		draw_pos_x, draw_pos_y,
	)

	t.front.Overlay(
		b_buff,
		draw_pos_x, draw_pos_y)
	t.term_state.update_cursor_pos(b.Width(), b.Height())
}

// when

// tab group
func (t *Term) BeginTabGroup() {
	// store := GetUniqueStore(t)
}

func (t *Term) EndTabGroup() {

}

func (t *Term) BeginTab(name string) bool {
	return true
}

func (t *Term) EndTab() {

}

func (t *Term) Space(width int, height int) {
	// t.term_state.get_cursor_pos()
	t.term_state.update_cursor_pos(width, height)

}
