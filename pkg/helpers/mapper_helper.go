package helpers

type mapperHelper[Before any, After any] struct{}

func (mapperHelper[Before, After]) MapSlice(slice []Before, mapFunc func(Before) After) (result []After) {
	result = make([]After, 0, len(slice))
	for _, v := range slice {
		result = append(result, mapFunc(v))
	}
	return
}
