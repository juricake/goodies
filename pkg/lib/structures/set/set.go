package set

// Set data structure is a wrapper around the idiomatic way to emulate sets in Go.
type Set[V comparable] struct {
	store map[V]struct{}
}

// New constructor.
func New[V comparable](values ...V) *Set[V] {
	s := &Set[V]{store: make(map[V]struct{}, 0)}
	s.Put(values...)
	return s
}

// Put value inside the set.
func (s *Set[V]) Put(values ...V) {
	for _, value := range values {
		s.store[value] = struct{}{}
	}
}

// Del value from the set.
func (s *Set[V]) Del(value V) {
	delete(s.store, value)
}

// Contains returns true if the set contains the value, false otherwise.
func (s *Set[V]) Contains(value V) bool {
	_, exists := s.store[value]
	return exists
}

// Size of the set.
func (s *Set[V]) Size() int {
	return len(s.store)
}
