package helpers

func Ptr[T any](value T) *T {
	return &value
}

func PtrSafeDeref[T any](ptr *T) T {
	if ptr == nil {
		return *new(T)
	}
	return *ptr
}
