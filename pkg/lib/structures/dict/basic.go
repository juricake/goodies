package dict

import "sort"

// Map data structure.
type Map[K comparable, V any] struct {
	store map[K]V
}

// New constructor.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		store: make(map[K]V, 0),
	}
}

// Put value inside the key.
func (m *Map[K, V]) Put(key K, value V) {
	m.store[key] = value
}

// Get values belonging to a specified key. Returns nil if none exists.
func (m *Map[K, V]) Get(key K) V {
	return m.store[key]
}

// Contains returns true if a key exists, false otherwise.
func (m *Map[K, V]) Contains(key K) bool {
	_, exists := m.store[key]
	return exists
}

// KeysSorted returns a sorted slice of keys, conforming to the key comparator func.
func (m *Map[K, V]) KeysSorted(comparator func(one, two K) bool) []K {
	keys := make([]K, 0, len(m.store))
	for k := range m.store {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return comparator(keys[i], keys[j])
	})

	return keys
}

// KeysSortedByValue returns a sorted slice of keys, conforming to the value comparator func.
func (m *Map[K, V]) KeysSortedByValue(comparator func(one, two V) bool) []K {
	keys := make([]K, 0, len(m.store))
	for k := range m.store {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return comparator(m.Get(keys[i]), m.Get(keys[j]))
	})

	return keys
}

// Size of the map.
func (m *Map[K, V]) Size() int {
	return len(m.store)
}
