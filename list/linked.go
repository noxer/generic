package list

import (
	"constraints"

	"github.com/noxer/generic"
)

type Linked[T any] struct {
	length int
	first  *singleLink[T]
}

func (s *Linked[T]) Iterator() generic.Iterator[T] {
	return &iterator[T]{curr: s.first}
}

type singleLink[T any] struct {
	n     *singleLink[T]
	value T
}

type iterator[T any] struct {
	curr *singleLink[T]
}

func (i *iterator[T]) Next() bool {
	if i.curr == nil {
		return false
	}

	i.curr = i.curr.n
	return i.curr != nil
}

func (i *iterator[T]) Val() (t T) {
	if i.curr != nil {
		t = i.curr.value
	}
	return
}

func (l *Linked[T]) Len() int {
	return l.length
}

func (l *Linked[T]) Slice() []T {
	slice := make([]T, 0, l.Len())
	iterator := l.Iterator()
	for iterator.Next() {
		slice = append(slice, iterator.Val())
	}
	return slice
}

func (l *Linked[T]) Any() (t T) {
	if l.Len() > 0 {
		t = l.first.value
	}

	return
}

func (l *Linked[T]) Put(items ...T) {
	for _, item := range items {
		newLink := &singleLink[T]{
			n:     l.first,
			value: item,
		}

		l.first = newLink
	}

	l.length += len(items)
}

type ComparableLinked[T comparable] struct {
	Linked[T]
}

func (l *ComparableLinked[T]) Contains(item T) bool {
	iter := l.Iterator()
	for iter.Next() {
		if iter.Val() == item {
			return true
		}
	}

	return false
}

func (l *ComparableLinked[T]) Remove(item T) {
	next := &l.first
	for *next != nil {
		if (*next).value == item {
			*next = (*next).n
			return
		}

		*next = (*next).n
	}
}

type OrderedLinked[T constraints.Ordered] struct {
	ComparableLinked[T]
}

func (l *OrderedLinked[T]) First() (t T) {
	if l.first != nil {
		t = l.first.value
	}

	return
}

func (l *OrderedLinked[T]) Last() (t T) {
	if l.first == nil {
		return
	}

	cur := l.first
	for cur.n != nil {
		cur = cur.n
	}
	return cur.value
}

func (l *OrderedLinked[T]) Get(i int) (t T) {
	iter := l.Iterator()

	for n := 0; n < i; n++ {
		if !iter.Next() {
			return t
		}
	}

	return iter.Val()
}

func (l *OrderedLinked[T]) Less(i, j int) bool {
	return l.Get(i) < l.Get(j)
}

var _ generic.ComparableCollection[int] = (*ComparableLinked[int])(nil)
