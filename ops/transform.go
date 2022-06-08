package ops

import "fmt"

// Rust inspired, see: https://doc.rust-lang.org/nightly/core/ops/trait.FnOnce.html
type transform[U any, V any] struct {
	call func(U) V
}

func Transform[U any, V any](f func(U) V) transform[U, V] {
	var t transform[U, V]
	t.call = f
	return t
}

func (t transform[U, V]) Call(arg U) V {
	return t.call(arg)
}

func (t transform[U, V]) String() string {
	return fmt.Sprintf("Transform(%T)", t.call)
}
