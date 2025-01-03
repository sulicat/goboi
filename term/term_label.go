package term

type Label struct {
	val string
}

func CreateLabel(s string) *Label {
	return &Label{val: s}
}

func (l *Label) Width() int {
	return len(l.val)
}

func (l *Label) Height() int {
	return 1
}

func (l *Label) Render(t *TermState) *FrameBuffer {
	out := FrameBuffer{}
	out.Make(l.Width(), l.Height())

	for i, s := range l.val {
		out[0][i].has_changed = true
		out[0][i].Char = string(s)
		out[0][i].BGColor = t.color_scheme.BackgroundColor
		out[0][i].FGColor = t.color_scheme.TextColor
	}

	return &out
}
