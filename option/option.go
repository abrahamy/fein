package option

import (
	"fmt"
	"reflect"
)

// Rust inspired, see: https://doc.rust-lang.org/nightly/core/option/enum.Option.html
type Option[T any] struct {
	value T
	some  bool
}

// Construct an option with the given value.
func Some[T any](value T) Option[T] {
	var option Option[T]
	option.value, option.some = value, true
	return option
}

// Construct an option without a value.
func None[T any]() Option[T] {
	var option Option[T]
	return option
}

// Returns None if the option is None, otherwise returns optb.
func (o Option[T]) And(other Option[any]) Option[any] {
	if o.Some() {
		return other
	}
	return None[any]()
}

// Returns None if the option is None, otherwise calls f with the wrapped value and returns the result.
func (o Option[T]) AndThen(f func(T) Option[any]) Option[any] {
	if o.None() {
		return None[any]()
	}
	return f(o.Unwrap())
}

// Returns true if other is equal to this option.
func (o Option[T]) Equal(other Option[T]) bool {
	return (o.None() && other.None()) || (o.Some() && other.Some() && reflect.DeepEqual(o.Unwrap(), other.Unwrap()))
}

// Returns true if the given option has a value.
func (o Option[T]) Some() bool {
	return o.some
}

// Returns false if the given option does not have a value.
func (o Option[T]) None() bool {
	return !o.Some()
}

func (o Option[T]) Unwrap() T {
	if o.Some() {
		return o.value
	}

	panic("Cannot unwrap an option of variant None!")
}

func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.Some() {
		return o.value
	}
	return defaultValue
}

func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.Some() {
		return o.value
	}
	return f()
}

func (o Option[T]) String() string {
	if o.Some() {
		return fmt.Sprintf("Some(%v)", o.value)
	}
	return "None"
}
