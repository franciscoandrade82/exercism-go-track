package strain

func Keep[T any](s []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

func Discard[T any](s []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range s {
		if !predicate(v) {
			result = append(result, v)
		}
	}

	return result
}
