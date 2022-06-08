package option_test

import (
	"testing"

	"github.com/abrahamy/fein/option"
	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {
	x := option.Some(2)
	y := option.None[any]()
	assert.True(t, x.And(y).Equal(option.None[any]()))

	x = option.None[int]()
	y = option.Some[any]("test")
	assert.True(t, x.And(y).Equal(option.None[any]()))

	x = option.Some(2)
	y = option.Some[any]("test")
	assert.True(t, x.And(y).Equal(option.Some[any]("test")))

	x = option.None[int]()
	y = option.None[any]()
	assert.True(t, x.And(y).Equal(option.None[any]()))
}

func TestEqual(t *testing.T) {
	assert.True(t, option.Some(2).Equal(option.Some(2)))
	assert.True(t, option.None[int]().Equal(option.None[int]()))
	assert.False(t, option.None[int]().Equal(option.Some(2)))
	assert.False(t, option.Some(2).Equal(option.None[int]()))
	assert.False(t, option.Some(2).Equal(option.Some(5)))
}
