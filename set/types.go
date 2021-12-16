package set

import "github.com/noxer/generic/util"

type Set[T comparable] map[T]struct{}

func (s Set[T]) Clone() Set[T] {
	c := make(Set[T], len(s))
	for item := range s {
		c[item] = struct{}{}
	}
	return c
}

func (s Set[T]) Put(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func (s Set[T]) Slice() []T {
	slice := make([]T, 0, len(s))

	for item := range s {
		slice = append(slice, item)
	}

	return slice
}

func (s Set[T]) Subtract(b Set[T]) {
	for item := range b {
		delete(s, item)
	}
}

func (s Set[T]) Any() T {
	for item := range s {
		return item
	}

	return util.Empty[T]()
}

func (s Set[T]) Len() int {
	return len(s)
}

func Intersect[T comparable](a, b Set[T]) Set[T] {
	c := make(Set[T])

	for item := range a {
		if _, ok := b[item]; ok {
			c[item] = struct{}{}
		}
	}

	return c
}

func Union[T comparable](a, b Set[T]) Set[T] {
	c := a.Clone()

	for item := range b {
		c[item] = struct{}{}
	}

	return c
}
