package mat

func New2DMat[T any](width int, height int) [][]T {
	out := make([][]T, height)
	for r := 0; r < height; r++ {
		out[r] = make([]T, width)
	}
	return out
}

func Subset[T any](start_x, end_x, start_y, end_y int) [][]T {
	out := make([][]T, end_y-start_y)

	return out
}
