package predicate

type Predicate[U any, V any] struct {
	callable func(U) V
}

func New[U any, V any](f func(U) V) Predicate[U, V] {
	var p Predicate[U, V]
	p.callable = f
	return p
}

func (p Predicate[U, V]) Call(arg U) V {
	return p.callable(arg)
}
