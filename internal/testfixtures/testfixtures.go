package testfixtures

func Find[T any](data []T, fn func(T) bool) (T, bool) {
	for _, item := range data {
		if fn(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}
