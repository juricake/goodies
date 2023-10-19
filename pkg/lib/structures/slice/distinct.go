package slice

import (
	"github.com/juricake/goodies/pkg/lib/structures/set"
)

// Distinct slice data structure is a wrapper around the slice, with a guarantee of uniqueness.
// This means there will never be duplicates in a slice.
type Distinct[V comparable] struct {
	store  []V
	filter *set.Set[V]
}

// NewDistinct slice constructor.
func NewDistinct[V comparable](values ...V) *Distinct[V] {
	d := &Distinct[V]{store: make([]V, 0), filter: set.New[V]()}
	d.Append(values...)
	return d
}

// Append a single or multiple values inside the slice.
func (s *Distinct[V]) Append(values ...V) {
	for _, value := range values {
		if s.filter.Contains(value) {
			continue
		}
		s.store = append(s.store, value)
		s.filter.Put(value)
	}
}

// GetAll values from the slice.
func (s *Distinct[V]) GetAll() []V {
	return s.store
}

// Size of the slice.
func (s *Distinct[V]) Size() int {
	return len(s.store)
}
