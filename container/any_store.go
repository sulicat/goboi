package container

import "reflect"

type AnyStore struct {
	data map[string]any
}

func CreateAnyStore() *AnyStore {
	out := &AnyStore{}
	out.data = map[string]any{}
	return out
}

func (as *AnyStore) Store(key string, val any) {
	as.data[key] = val
}

func (as *AnyStore) Get(key string) any {
	v, exists := as.data[key]
	if exists {
		return v
	}

	var empty any
	return reflect.Zero(reflect.TypeOf(empty)).Interface()
}

func AnyStoreGetAs[T any](as *AnyStore, key string) T {
	v, exists := as.data[key]
	if exists {
		return v.(T)
	}

	var out T
	return out
}
