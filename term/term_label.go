package term

type Label struct {
	val string
	x   int
	y   int
}

func CreateLabel(s string) *Label {
	return &Label{val: s}
}

func (l *Label) X() int {
	return 0
}

func (l *Label) Y() int {
	return 0
}

func (l *Label) Width() int {
	return len(l.val)
}

func (l *Label) Height() int {
	return 1
}

func (l *Label) Render() *FrameBuffer {
	out := FrameBuffer{}
	out.Make(l.Width(), l.Height())

	for i, s := range l.val {
		out[0][i].has_changed = true
		out[0][i].Char = string(s)
		out[0][i].BGColor = RGB{255, 0, 0}
		out[0][i].FGColor = RGB{255, 0, 255}
	}

	return &out
}
