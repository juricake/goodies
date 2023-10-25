package dict

import "sort"

// SliceMap is a data structure designed to operate over maps with multiple values belonging to a single key.
// It simplifies operation like initialization, appending, or sorting by conditions.
type SliceMap[K comparable, V comparable] struct {
	store map[K][]V
}

// NewSliceMap constructor.
func NewSliceMap[K comparable, V comparable]() *SliceMap[K, V] {
	return &SliceMap[K, V]{
		store: make(map[K][]V, 0),
	}
}

// Put values inside the key.
// This operation will replace the underlying slice of values with the specified one.
func (s *SliceMap[K, V]) Put(key K, values []V) {
	s.store[key] = values
}

// Get values belonging to a specified key. Returns nil if none exists.
func (s *SliceMap[K, V]) Get(key K) []V {
	return s.store[key]
}

// Del value belonging to a specified key.
func (s *SliceMap[K, V]) Del(key K) {
	delete(s.store, key)
}

// Append a new value to a list of values belonging to a specified key.
// This operation will auto-create the underlying slice if necessary.
func (s *SliceMap[K, V]) Append(key K, value V) {
	if !s.Contains(key) {
		s.store[key] = make([]V, 0)
	}
	values := s.store[key]
	values = append(values, value)
	s.store[key] = values
}

// Contains returns true if a key exists, false otherwise.
func (s *SliceMap[K, V]) Contains(key K) bool {
	_, exists := s.store[key]
	return exists
}

// ContainsWithin returns true if both Contains is true, and a specified value is contained within the
// underlying slice assigned to a key.
func (s *SliceMap[K, V]) ContainsWithin(key K, value V) bool {
	if !s.Contains(key) {
		return false
	}

	for _, val := range s.Get(key) {
		if val == value {
			return true
		}
	}

	return false
}

// KeysSorted returns a sorted slice of keys, conforming to the key comparator func.
func (s *SliceMap[K, V]) KeysSorted(comparator func(one, two K) bool) []K {
	keys := make([]K, 0, len(s.store))
	for k := range s.store {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return comparator(keys[i], keys[j])
	})

	return keys
}

// KeysSortedByValue returns a sorted slice of keys, conforming to the value comparator func.
func (s *SliceMap[K, V]) KeysSortedByValue(comparator func(one, two []V) bool) []K {
	keys := make([]K, 0, len(s.store))
	for k := range s.store {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return comparator(s.Get(keys[i]), s.Get(keys[j]))
	})

	return keys
}

// Size of the map.
func (s *SliceMap[K, V]) Size() int {
	return len(s.store)
}
