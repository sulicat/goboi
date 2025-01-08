package term

type CanvasType struct {
	pixels *[][]RGBA
}

func (b *CanvasType) Width() int {
	if b.Height() > 0 {
		return len((*b.pixels)[0])
	} else {
		return 0
	}
}

func (b *CanvasType) Height() int {
	return len(*b.pixels)

}

// offset is to handle input, it assumes it will be overlayed
// at offset_x, offset_y
func (b *CanvasType) Render(
	state *TermState,
	offset_x int, offset_y int,
) *FrameBuffer {

	out := FrameBuffer{}
	out.Make(b.Width(), b.Height())

	for y := range b.Height() {
		for x := range b.Width() {

			if (*b.pixels)[y][x][3] == 0 {
				out[y][x].has_changed = false

			} else {
				out[y][x].Char = CharSquare
				out[y][x].FGColor = (*b.pixels)[y][x].RGB()
				out[y][x].has_changed = true
			}
		}
	}

	return &out
}

func CreateCanvasType(pixels *[][]RGBA) CanvasType {
	out := CanvasType{pixels: pixels}

	return out
}
