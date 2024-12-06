package utils

func MapContains[T comparable, U any](in map[T]U, key T) bool {
	_, contains := in[key]
	return contains
}
