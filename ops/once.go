package ops

import "fmt"

// Rust inspired, see: https://doc.rust-lang.org/nightly/core/ops/trait.FnOnce.html
type fnOnce[T any] struct {
	f func() T
}

func FnOnce[T any](f func() T) fnOnce[T] {
	var instance fnOnce[T]
	instance.f = f
	return instance
}

func (self *fnOnce[T]) Call() T {
 if self.f == nil {
  panic("Function of type FnOnce called more than once!")
	}
	
 result := self.f()
	self.f = nil
	return result
}

// This is an alias for FnOnce::Call method 
func (self *fnOnce[T]) Apply() T {
 return self.Call()
}

func (self fnOnce[T]) String() string {
	return fmt.Sprintf("FnOnce(%T)", self.f)
}
