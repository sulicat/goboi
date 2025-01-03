package term

type ColorScheme struct {
	TextColor      RGB
	HoverColor     RGB
	SelectedColor  RGB
	PrimaryColor   RGB
	SecondaryColor RGB
}

func CreateColorScheme() ColorScheme {
	out := ColorScheme{
		TextColor:      RGB{255, 255, 255},
		HoverColor:     RGB{234, 152, 40},
		SelectedColor:  RGB{245, 103, 2},
		PrimaryColor:   RGB{234, 152, 40},
		SecondaryColor: RGB{26, 186, 175},
	}
	return out
}
