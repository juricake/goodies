package transform

import "github.com/juricake/goodies/pkg/lib/structures/dict"

// ToMap converts a given slice of values into a map where keys are specified according to the getKey func,
// while the values are the initial values from the slice. This func is often used for indexing of complex structures.
// In a case of duplicate keys, only the last processed value will be retained. If given an empty slice,
// this func will return an empty map.
func ToMap[T any, t comparable](values []T, getKey func(value T) t) *dict.Map[t, T] {
	result := dict.New[t, T]()

	if len(values) == 0 {
		return result
	}

	for _, value := range values {
		result.Put(getKey(value), value)
	}
	return result
}

// ToSliceMap converts a given slice of values into a map where keys are specified according to the getKey func,
// while the values are slices containing the initial values from the input slice.
// This func is often used for indexing of complex structures where multiple values for the same key are expected.
// In a case of duplicate keys, all according values will be retained. If given an empty slice, this func will return an empty map.
func ToSliceMap[T comparable, t comparable](values []T, getKey func(value T) t) *dict.SliceMap[t, T] {
	result := dict.NewSliceMap[t, T]()
	if len(values) == 0 {
		return result
	}

	for _, value := range values {
		result.Append(getKey(value), value)
	}
	return result
}
