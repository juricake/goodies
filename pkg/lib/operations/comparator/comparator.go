package comparator

import "sort"

// First returns the first value after a given slice of values has been sorted.
// This func is often used for Min/Max extraction.
// If given an empty slice, it will return a nil.
func First[V any](values []V, compare func(first, second V) bool) V {
	if len(values) == 0 {
		return nil
	}

	sort.SliceStable(values, func(i, j int) bool {
		return compare(values[i], values[j])
	})

	return values[0]
}
