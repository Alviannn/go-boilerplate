package helpers

var (
	Http   = httpHelper{}
	Random = randomHelper{}
)

func Slice[T comparable]() sliceHelper[T] {
	return sliceHelper[T]{}
}
