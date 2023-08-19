package util

// Returns the number of elements matching the given predicate.
func CountMatches[T any](data []T, predicate func(T) bool) int {
	count := 0
	for _, s := range data {
		if predicate(s) {
			count += 1
		}
	}
	return count
}
