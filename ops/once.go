package ops

import "fmt"

// Rust inspired, see: https://doc.rust-lang.org/nightly/core/ops/trait.FnOnce.html
type fnOnce[T any] struct {
	call   func() T
	result T
}

func FnOnce[T any](f func() T) fnOnce[T] {
	var fn fnOnce[T]
	fn.call = f
	return fn
}

func (f *fnOnce[T]) Call() T {
	if f.call != nil {
		f.result = f.call()
		f.call = nil
	}
	return f.result
}

func (f fnOnce[T]) String() string {
	return fmt.Sprintf("FnOnce(%T)", f.call)
}
