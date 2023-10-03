package helpers

type sliceHelper[T comparable] struct{}

func (sliceHelper[T]) IsIn(slice []T, itemToFind T) bool {
	for _, item := range slice {
		if itemToFind == item {
			return true
		}
	}

	return false
}
