package ops

import "fmt"

// Rust inspired, see: https://doc.rust-lang.org/nightly/core/ops/trait.FnOnce.html
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

func (p *FnOnce[T]) Call() T {
	if !p.called {
		p.value, p.called = p.callable(), true
	}
	return p.value
}

func (p FnOnce[T]) String() string {
	return fmt.Sprintf("FnOnce(%T)", p.callable)
}
