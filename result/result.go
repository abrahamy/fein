package result

import (
	"mayhap/option"
	"mayhap/predicate"
	"reflect"
)

type Result[T any, E any] struct {
	value T
	err   E
	isErr bool
}

/*
 *	Constructors
 */

// Constructor for Ok variant.
func Ok[T any, E any](value T) Result[T, E] {
	var rs Result[T, E]
	rs.value = value
	return rs
}

// Constructor for Err variant.
func Err[T any, E any](err E) Result[T, E] {
	var rs Result[T, E]
	rs.err, rs.isErr = err, true
	return rs
}

/*
 *	Methods
 */

// Converts from Result[T, E] to Option[T].
func (rs Result[T, E]) Ok() option.Option[T] {
	if rs.IsErr() {
		none := option.None[T]()
		return none
	}

	return option.Some(rs.value)
}

// Converts from Result[T, E] to Option[E].
func (rs Result[T, E]) Err() option.Option[E] {
	if !rs.IsErr() {
		none := option.None[E]()
		return none
	}

	return option.Some(rs.err)
}

/*
 Returns value if the result is Ok, otherwise returns the Err value of result.

 This implementation is not fully compatible to the equivalent Rust implementation
 since Go does not yet support generic methods nor union types. In the Rust implementation
 the `other` parameter can have different types, e.g. `other Result[T2, E2]` making the
 result of type `Result[T, E] | Result[T2, E2]`
*/
func (rs Result[T, E]) And(other Result[T, E]) Result[T, E] {
	if rs.IsErr() {
		return rs
	}

	return other
}

/*
 Calls f if the result is Ok, otherwise returns the Err value of the result.

 This method has the same limitations as Result::And, the actual types of the
 Rust version is func (rs Result[T, E]) AndThen(f Predicate[T, U]) Result[U, E]
*/
func (rs Result[T, E]) AndThen(f predicate.Predicate[T, any]) Result[any, E] {
	if rs.IsErr() {
		return Err[any](rs.err)
	}
	return Ok[any, E](f.Call(rs.value))
}

func (rs Result[T, E]) IsErr() bool {
	return rs.isErr
}

func (rs Result[T, E]) IsOK() bool {
	return !rs.IsErr()
}

func (rs Result[T, E]) Contains(value T) bool {
	return rs.IsOK() && reflect.DeepEqual(rs.value, value)
}

func (rs Result[T, E]) ContainsErr(err E) bool {
	return rs.IsErr() && reflect.DeepEqual(rs.err, err)
}

func (rs Result[T, E]) Expect(msg string) T {
	if rs.IsOK() {
		return rs.Ok().Unwrap()
	}

	panic(msg)
}

func (rs Result[T, E]) ExpectErr(msg string) E {
	if rs.IsErr() {
		return rs.Err().Unwrap()
	}

	panic(msg)
}

func (rs Result[T, E]) Map(f predicate.Predicate[T, any]) Result[any, E] {
	return rs.AndThen(f)
}
