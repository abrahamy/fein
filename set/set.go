package set

import "fmt"

// Python inspired Set data structure, see: https://docs.python.org/3/library/stdtypes.html#set
type Set[T comparable] struct {
	inner map[T]bool
}

// Constructor: make a new empty set.
func New[T comparable]() Set[T] {
	s := Set[T]{
		make(map[T]bool, 0),
	}
	return s
}

// Constructor: make a new set with initial elements.
func FromSlice[T comparable](elem []T) Set[T] {
	s := New[T]()
	s.Add(elem...)
	return s
}

// Add element(s) to the set and return the number of new items.
func (s Set[T]) Add(elem ...T) int {
	previousSize := len(s.inner)
	for _, e := range elem {
		s.inner[e] = true
	}

	return len(s.inner) - previousSize
}

// Return the number of elements in the set (cardinality of the set).
func (s Set[T]) Len() int {
	return len(s.inner)
}

// Returns a slice containing all members of the set.
func (s Set[T]) Members() []T {
	members := make([]T, 0)
	for k := range s.inner {
		members = append(members, k)
	}
	return members
}

// Return a new set with elements from the set s and all others.
func (s Set[T]) Union(other ...Set[T]) Set[T] {
	u := New[T]()
	u.Add(s.Members()...)
	for _, o := range other {
		u.Add(o.Members()...)
	}
	return u
}

func (s Set[T]) String() string {
	return fmt.Sprintf("Set%v", s.Members())
}

func (s Set[T]) Empty() bool {
	return s.Len() == 0
}

// Returns true if elements are members of the set.
func (s Set[T]) Contains(elem ...T) bool {
	for _, e := range elem {
		if _, ok := s.inner[e]; !ok {
			return false
		}
	}
	return true
}

// Return a new set with elements common to the set and other.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	members := make([]T, 0)
	for _, elem := range s.Members() {
		if other.Contains(elem) {
			members = append(members, elem)
		}
	}
	return FromSlice(members)
}

// Return a new set with elements in the set that are not in the others.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	members := make([]T, 0)
	for _, elem := range s.Members() {
		if !other.Contains(elem) {
			members = append(members, elem)
		}
	}
	return FromSlice(members)
}

// Return a new set with elements in either the set or other but not both.
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	members := s.Difference(other).Members()
	members = append(members, other.Difference(s).Members()...)
	return FromSlice(members)
}

// Return true if the set has no elements in common with other.
// Sets are disjoint if and only if their intersection is the empty set.
func (s Set[T]) Disjoint(other Set[T]) bool {
	return s.Intersection(other).Empty()
}

// Returns true if every element in the set is also in other.
func (s Set[T]) Subset(other Set[T]) bool {
	return other.Contains(s.Members()...)
}

// Returns true if every element in the set is also in other and
// other contains at least one element not in set.
func (s Set[T]) ProperSubset(other Set[T]) bool {
	return s.Subset(other) && s.Len() < other.Len()
}

// Returns true if every element in other is in the set.
func (s Set[T]) Superset(other Set[T]) bool {
	return s.Contains(other.Members()...)
}

// Return true if the set is a proper superset of other,
// that is, set >= other and set != other.
func (s Set[T]) ProperSuperset(other Set[T]) bool {
	return s.Superset(other) && s.Len() > other.Len()
}

// Remove element elem from the set. Panics if elem is not contained in the set.
func (s Set[T]) Remove(elem T) {
	if !s.Contains(elem) {
		panic(fmt.Sprintf("set %s does not contain the element %v.", s.String(), elem))
	}
	delete(s.inner, elem)
}

// Remove element elem from the set if it is present. Never panics.
func (s Set[T]) Discard(elem ...T) {
	for _, e := range elem {
		delete(s.inner, e)
	}
}

// Remove and return an arbitrary element from the set. Panics if the set is empty.
func (s Set[T]) Pop() T {
	if s.Empty() {
		panic("cannot pop item from an empty set!")
	}
	members := s.Members()
	val := members[len(members)-1]
	delete(s.inner, val)
	return val
}

// Remove all elements from the set.
func (s Set[T]) Clear() {
	s.inner = make(map[T]bool, 0)
	_ = s.inner
}

// Update the set, adding elements from all others.
func (s Set[T]) Update(other ...Set[T]) {
	for _, o := range other {
		s.Add(o.Members()...)
	}
}

// Update the set, removing elements found in others.
func (s Set[T]) DifferenceUpdate(other ...Set[T]) {
	intersections := New[T]()
	for _, o := range other {
		intersections.Union(s.Intersection(o))
	}
	members := s.Difference(intersections).Members()
	s.Clear()
	s.Add(members...)
}

// Update the set, keeping only elements found in it and all others.
func (s Set[T]) IntersectionUpdate(other ...Set[T]) {
	// @todo: implement
}

// Update the set, keeping only elements found in either set, but not in both.
func (s Set[T]) SymmetricDifferenceUpdate(other Set[T]) {
	members := s.SymmetricDifference(other).Members()
	s.Clear()
	s.Add(members...)
}
