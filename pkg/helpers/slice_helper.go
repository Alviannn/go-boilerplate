package helpers

func SliceIsInWithFunc[T comparable](slice []T, itemToFind T, matchFunc func(expected T, actual T) bool) bool {
	for _, item := range slice {
		if matchFunc(itemToFind, item) {
			return true
		}
	}
	return false
}

func SliceIsIn[T comparable](slice []T, itemToFind T) bool {
	return SliceIsInWithFunc(slice, itemToFind, func(expected T, actual T) bool {
		return expected == actual
	})
}
