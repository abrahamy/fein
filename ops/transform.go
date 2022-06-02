package ops

// Rust inspired, see: https://doc.rust-lang.org/nightly/core/ops/trait.FnOnce.html
type Transform[U any, V any] struct {
	callable func(U) V
}

func NewTransform[U any, V any](f func(U) V) Transform[U, V] {
	var p Transform[U, V]
	p.callable = f
	return p
}

func (p Transform[U, V]) Call(arg U) V {
	return p.callable(arg)
}
