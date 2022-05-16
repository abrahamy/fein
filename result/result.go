package result

import (
	"mayhap/option"
	"mayhap/predicate"
)

type Result[T any, E any] struct {
	value T
	err   E
	isErr bool
}

// Constructor for Ok
func Ok[T any, E any](value T) Result[T, E] {
	var rs Result[T, E]
	rs.value = value
	return rs
}

// Constructor for Err
func Err[T any, E any](err E) Result[T, E] {
	var rs Result[T, E]
	rs.err, rs.isErr = err, true
	return rs
}

// This implementation is not fully compatible to the equivalent
// implementation in rust since Go does not yet support generic
// methods or union types. In the Rust implementation the result
// `other` can have different types, e.g. `other Result[T2, E2]`
// making the result of type `Result[T, E] | Result[T2, E2]`
func (rs Result[T, E]) And(other Result[T, E]) Result[T, E] {
	if rs.IsErr() {
		return rs
	}

	return other
}

// The Rust equivalent has the following Go signature.
// func (rs Result[T, E]) AndThen(f Predicate[T, U]) Result[U, E]
func (rs Result[T, E]) AndThen(f predicate.Predicate[T, any]) Result[any, E] {
	if rs.IsErr() {
		return Err[any](rs.err)
	}
	return Ok[any, E](f.Call(rs.value))
}

func (rs Result[T, E]) IsErr() bool {
	return rs.isErr
}

func (rs Result[T, E]) Ok() option.Option[T] {
	if rs.IsErr() {
		none := option.None[T]()
		return none
	}

	return option.Some(rs.value)
}

func (rs Result[T, E]) Err() option.Option[E] {
	if !rs.IsErr() {
		none := option.None[E]()
		return none
	}

	return option.Some(rs.err)
}
