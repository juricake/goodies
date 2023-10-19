package dict

// SliceMap data structure.
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

// Size of the map.
func (s *SliceMap[K, V]) Size() int {
	return len(s.store)
}
