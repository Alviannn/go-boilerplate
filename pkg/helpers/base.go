package helpers

var (
	Echo   = echoHelper{}
	Random = randomHelper{}
)

func Slice[T comparable]() sliceHelper[T] {
	return sliceHelper[T]{}
}

func Mapper[Before, After any]() mapperHelper[Before, After] {
	return mapperHelper[Before, After]{}
}
