package helpers

func SliceMap[Before any, After any](slice []Before, mapFunc func(Before) After) (newList []After) {
	newList = make([]After, len(slice))
	for i, v := range slice {
		newList[i] = mapFunc(v)
	}
	return
}
