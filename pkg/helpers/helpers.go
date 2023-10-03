package helpers

var (
	Echo = echoHelper{}
	Http = httpHelper{}
)

func Slice[T comparable]() sliceHelper[T] {
	return sliceHelper[T]{}
}
