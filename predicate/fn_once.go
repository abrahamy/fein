package predicate

type FnOnce[T any] struct {
	callable func() T
	called   bool
	value    T
}

func NewFnOnce[T any](f func() T) FnOnce[T] {
	var p FnOnce[T]
	p.callable = f
	return p
}

func (p FnOnce[T]) Call() T {
	if p.called {
		return p.value
	}
	p.value, p.called = p.callable(), true
	return p.value
}
