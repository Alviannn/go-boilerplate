package helpers

type sliceHelper[T comparable] struct{}

func (h sliceHelper[T]) IsInWithFunc(slice []T, itemToFind T, equalFunc func(T, T) bool) bool {
	for _, item := range slice {
		if equalFunc(itemToFind, item) {
			return true
		}
	}
	return false
}

func (h sliceHelper[T]) IsIn(slice []T, itemToFind T) bool {
	return h.IsInWithFunc(
		slice,
		itemToFind,
		func(a T, b T) bool {
			return a == b
		},
	)
}
