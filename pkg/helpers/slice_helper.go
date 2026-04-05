package helpers

func SliceMap[Before any, After any](slice []Before, mapFunc func(Before) After) (newList []After) {
	newList = make([]After, len(slice))
	for i, v := range slice {
		newList[i] = mapFunc(v)
	}
	return
}

func SliceDeduplicateFunc[T any, K comparable](oldList []T, keyFunc func(T) K) []T {
	var (
		seen   = make(map[K]struct{}, len(oldList))
		result = make([]T, 0)
	)

	for _, v := range oldList {
		key := keyFunc(v)

		if _, exists := seen[key]; exists {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, v)
	}

	return result
}

func SliceDeduplicate[T comparable](oldList []T) []T {
	return SliceDeduplicateFunc(oldList, func(key T) T {
		return key
	})
}
