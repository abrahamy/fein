package result_test

import (
	"math"
	"testing"

	"github.com/abrahamy/fein/result"
	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	ok := result.Ok[int, string](1)
	assert.NotNil(t, ok)
}

func TestErr(t *testing.T) {
	err := result.Err[int]("Error!!!")
	assert.NotNil(t, err)
}

func TestAnd(t *testing.T) {
	ok := result.Ok[int, string](1)
	alsoOk := result.Ok[int, string](2)

	err := result.Err[int]("this is an error!")
	alsoErr := result.Err[int]("this is also an error!")

	assert.Equal(t, ok.And(alsoOk), alsoOk)
	assert.Equal(t, ok.And(err), err)
	assert.Equal(t, err.And(ok), err)
	assert.Equal(t, err.And(alsoErr), err)
}

func TestAndThen(t *testing.T) {
	ok := result.Ok[float64, string](7)
	callable := func(i float64) any {
		var val any = math.Pow(i, 2)
		return val
	}
	var expected float64 = 49
	actual := ok.AndThen(callable).Ok().Unwrap().(float64)
	assert.Equal(t, actual, expected)

	err := result.Err[float64]("this is an error!")
	actualErr := err.AndThen(callable).Err().Unwrap()
	assert.Equal(t, actualErr, "this is an error!")
}

func TestIsErr(t *testing.T) {
	ok := result.Ok[int, error](1)
	assert.False(t, ok.IsErr())

	err := result.Err[any]("this is an error!")
	assert.True(t, err.IsErr())
}

func TestIsOkay(t *testing.T) {
	ok := result.Ok[int, string](1)
	assert.True(t, ok.IsOk())

	err := result.Err[int]("this is an error!")
	assert.False(t, err.IsOk())
}

func TestContains(t *testing.T) {
	ok := result.Ok[int, string](1)
	assert.True(t, ok.Contains(1))
	assert.False(t, ok.Contains(2))

	err := result.Err[int]("this is an error!")
	assert.False(t, err.Contains(1))
}

func TestContainsErr(t *testing.T) {
	ok := result.Ok[int, string](1)
	assert.False(t, ok.ContainsErr("this is an error!"))

	err := result.Err[int]("this is an error!")
	assert.True(t, err.ContainsErr("this is an error!"))
	assert.False(t, err.ContainsErr("A different error!"))
}

func TestExpect(t *testing.T) {
	msg := "Got an error!"
	ok := result.Ok[int, string](1)
	assert.Equal(t, ok.Expect(msg), 1)

	err := result.Err[int]("this is an error!")
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, msg, r)
		}
	}()
	err.Expect(msg)
}

func TestExpectErr(t *testing.T) {
	msg := "Did not get an error!"
	err := result.Err[int]("this is an error!")
	assert.Equal(t, err.ExpectErr(msg), "this is an error!")

	ok := result.Ok[int, string](1)
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, msg, r)
		}
	}()
	ok.ExpectErr(msg)
}
