package option

import (
	"fein/predicate"
)

type Option[T any] struct {
	value T
	some  bool
}

func Some[T any](value T) Option[T] {
	var option Option[T]
	option.value, option.some = value, true
	return option
}

func None[T any]() Option[T] {
	var option Option[T]
	return option
}

func (o Option[T]) IsSome() bool {
	return o.some
}

func (o Option[T]) IsNone() bool {
	return !o.some
}

func (o Option[T]) Unwrap() T {
	if o.IsSome() {
		return o.value
	}

	panic("Cannot unwrap an option of variant None!")
}

func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.IsSome() {
		return o.value
	}
	return defaultValue
}

func (o Option[T]) UnwrapOrElse(f predicate.FnOnce[T]) T {
	if o.IsSome() {
		return o.value
	}
	return f.Call()
}
