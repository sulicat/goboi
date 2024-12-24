package term

type TermState struct {
	// render state
	bg_color  RGB
	fg_color  RGB
	same_line bool

	MouseX      int
	MouseY      int
	MouseButton int
}
