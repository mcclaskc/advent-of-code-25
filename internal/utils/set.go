package utils

// Set is a generic, unordered collection of unique elements.
type Set[T comparable] map[T]struct{}

// exists is a zero-size value used as the map value for a set.
var exists = struct{}{}

// NewSet creates and returns a new empty Set.
func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	for _, item := range items {
		s.Add(item)
	}
	return s
}

// Add inserts a new element into the Set.
func (s Set[T]) Add(item T) {
	s[item] = exists
}

// Remove deletes an element from the Set.
func (s Set[T]) Remove(item T) {
	delete(s, item)
}

// Contains checks if the Set includes the given element.
func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

// Size returns the number of elements in the Set.
func (s Set[T]) Size() int {
	return len(s)
}

// Clear removes all elements from the Set.
func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Slice converts the Set into a slice of its elements.
func (s Set[T]) Slice() []T {
	result := make([]T, 0, s.Size())
	for item := range s {
		result = append(result, item)
	}
	return result
}

// UnionSet creates a new Set containing all elements from both sets.
func UnionSet[T comparable](s1, s2 Set[T]) Set[T] {
	result := NewSet[T]()
	for item := range s1 {
		result.Add(item)
	}
	for item := range s2 {
		result.Add(item)
	}
	return result
}

// IntersectionSet creates a new Set containing only elements present in both sets.
func IntersectionSet[T comparable](s1, s2 Set[T]) Set[T] {
	result := NewSet[T]()
	for item := range s1 {
		if s2.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// DifferenceSet creates a new Set containing elements from s1 that are not in s2 (s1 - s2).
func DifferenceSet[T comparable](s1, s2 Set[T]) Set[T] {
	result := NewSet[T]()
	for item := range s1 {
		if !s2.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
