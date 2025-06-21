package helpers

import "slices"

type SliceMatchFunc[T comparable] func(current T) bool

func SliceFindFunc[T comparable](slice []T, matchFunc SliceMatchFunc[T]) *T {
	idx := slices.IndexFunc(slice, matchFunc)
	if idx == -1 {
		return nil
	}
	return &slice[idx]
}

func SliceFind[T comparable](slice []T, itemToFind T) *T {
	return SliceFindFunc(slice, func(actual T) bool {
		return itemToFind == actual
	})
}

func SliceMap[A comparable, B any](slice []A, mapFunc func(A) B) (newList []B) {
	newList = make([]B, len(slice))
	for i, v := range slice {
		newList[i] = mapFunc(v)
	}
	return
}
