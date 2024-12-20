package term

type Renderable interface {
	X() int
	Y() int
	Width() int
	Height() int
	Render() *FrameBuffer
}

func (t *Term) Label(s string, x int, y int) {
	// t.draw_list = append(t.draw_list, CreateLabel(s))
	l := CreateLabel(s)
	l_buff := l.Render()
	t.front.Overlay(l_buff, x, y)
}
