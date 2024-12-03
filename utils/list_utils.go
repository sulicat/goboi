package utils

func First[T any](first T, rest ...any) T {
	return first
}

func Second[T any](first any, second T, rest ...any) T {
	return second
}

func MapArray[T any, U any](in []T, cb func(T) U) []U {
	out := make([]U, len(in))
	for i, val := range in {

		out[i] = cb(val)
	}
	return out
}

func ForEach[T any](in []T, cb func(T)) {
	for _, val := range in {
		cb(val)
	}
}
