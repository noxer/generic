package generic

import (
	"constraints"
)

type Cloner[T any] interface {
	Clone() T
}

type Collection[T any] interface {
	Len() int
	Slice() []T
	Any() T
	Iterator() Iterator[T]
}

type MutableCollection[T any] interface {
	Collection[T]
	Put(...T)
}

type ComparableCollection[T comparable] interface {
	Collection[T]
	Contains(T) bool
	Remove(T)
}

type MutableComparableCollection[T comparable] interface {
	ComparableCollection[T]
	MutableCollection[T]
}

type OrderedCollection[T constraints.Ordered] interface {
	ComparableCollection[T]
	First() T
	Last() T
	Get(i int) T
	Less(i, j int) bool
}

type MutableOrderedCollection[T constraints.Ordered] interface {
	OrderedCollection[T]
	MutableComparableCollection[T]

	Set(i int, item T)
	Append(item ...T)
	Swap(i, j int)
}

type Iterator[T any] interface {
	Next() bool
	Val() T
}

func Contains[T comparable](coll Collection[T], item T) bool {
	if compColl, ok := coll.(ComparableCollection[T]); ok {
		return compColl.Contains(item)
	}

	iter := coll.Iterator()
	for iter.Next() {
		if iter.Val() == item {
			return true
		}
	}

	return false
}
