package utils

import "cmp"

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

func MaxIndex[T cmp.Ordered](in []T) int {
	out_i := 0
	for i, v := range in {
		if v > in[out_i] {
			out_i = i
		}
	}

	return out_i
}

func ArrayInsert[T any](slice []T, index int, value T) []T {
	new_slice := make([]T, len(slice)+1)
	copy(new_slice[:index], slice[:index])
	new_slice[index] = value
	copy(new_slice[index+1:], slice[index:])
	return new_slice
}

func ArrayRemove[T any](slice []T, index int) []T {
	new_slice := make([]T, len(slice)-1)
	copy(new_slice[:index], slice[:index])
	copy(new_slice[index:], slice[index+1:])
	return new_slice
}
