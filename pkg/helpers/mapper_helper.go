package helpers

func MapperSlice[Before any, After any](slice []Before, mapFunc func(Before) After) (result []After) {
	result = make([]After, 0, len(slice))
	for _, v := range slice {
		result = append(result, mapFunc(v))
	}
	return
}
