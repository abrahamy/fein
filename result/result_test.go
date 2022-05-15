package result

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	ok := Ok[int, error](1)
	assert.NotNil(t, ok)
}

func TestErr(t *testing.T) {
	err := Err[int](errors.New("Error!!!"))
	assert.NotNil(t, err)
}

func TestAnd(t *testing.T) {
	ok := Ok[int, error](1)
	alsoOk := Ok[int, error](2)

	err := Err[int](errors.New("this is an error!"))
	alsoErr := Err[int](errors.New("this is also an error!"))

	assert.Equal(t, ok.And(alsoOk), alsoOk)
	assert.Equal(t, ok.And(err), err)
	assert.Equal(t, err.And(ok), err)
	assert.Equal(t, err.And(alsoErr), err)
}

func TestIsErr(t *testing.T) {
	ok := Ok[int, error](1)
	assert.False(t, ok.IsErr())

	err := Err[any](errors.New("this is an error!"))
	assert.True(t, err.IsErr())
}
