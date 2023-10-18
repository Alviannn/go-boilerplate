package helpers

var (
	Http = httpHelper{}
)

func Slice[T comparable]() sliceHelper[T] {
	return sliceHelper[T]{}
}
