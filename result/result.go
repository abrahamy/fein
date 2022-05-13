package result

type Result[T any, E error] struct {
	value T
	err   E
	isErr bool
}

// Constructor for Ok
func Ok[T any, E error](value T) Result[T, E] {
	var rs Result[T, E]
	rs.value = value
	return rs
}

// Constructor for Err
func Err[T any, E error](err E) Result[T, E] {
	var rs Result[T, E]
	rs.err = err
	rs.isErr = true
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

func (rs Result[T, E]) IsErr() bool {
	return rs.isErr
}
