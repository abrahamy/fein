package set

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New[int]()
	assert.NotNil(t, s)
	assert.NotNil(t, s.inner)
}

func TestFromSlice(t *testing.T) {
	s := FromSlice([]int{1, 2, 2, 3, 1, 1})
	assert.Equal(t, s.Len(), 3)
	assert.ElementsMatch(t, s.Elems(), []int{1, 2, 3})
}

func TestLen(t *testing.T) {
	s := New[int]()
	assert.Equal(t, s.Len(), 0)

	s.Add(1, 2, 3)
	assert.Equal(t, s.Len(), 3)

	assert.Equal(t, FromSlice([]int{4, 5, 5}).Len(), 2)
}

func TestAdd(t *testing.T) {
	s := New[int]()
	s.Add(1, 2, 3)
	assert.Equal(t, s.Len(), 3)

	s.Add(4, 5)
	assert.Equal(t, s.Len(), 5)

	assert.ElementsMatch(t, s.Elems(), []int{1, 2, 3, 4, 5})
}

func TestElems(t *testing.T) {
	emptySet := New[int]()
	intSet := FromSlice([]int{1, 2, 2, 3, 4, 4, 3})
	assert.Equal(t, len(emptySet.Elems()), 0)
	assert.Equal(t, len(intSet.Elems()), 4)
	assert.ElementsMatch(t, intSet.Elems(), []int{1, 2, 3, 4})
}

func TestUnion(t *testing.T) {
	emptySet := New[int]()
	oneTwoThree := FromSlice([]int{1, 2, 3})
	fourFive := FromSlice([]int{4, 5})
	sixSevenEight := FromSlice([]int{6, 7, 8})

	assert.ElementsMatch(t, emptySet.Union(oneTwoThree).Elems(), oneTwoThree.Elems())
	assert.ElementsMatch(t, oneTwoThree.Union(fourFive).Elems(), []int{1, 2, 3, 4, 5})
	assert.ElementsMatch(t, oneTwoThree.Union(fourFive, sixSevenEight).Elems(), []int{1, 2, 3, 4, 5, 6, 7, 8})
}

func TestString(t *testing.T) {
	emptySet := New[int]()
	noneEmpty := FromSlice([]int{1, 2, 3})
	assert.Equal(t, strings.Compare(emptySet.String(), fmt.Sprintf("Set%v", emptySet.Elems())), 0)
	assert.Equal(t, strings.Compare(noneEmpty.String(), fmt.Sprintf("Set%v", noneEmpty.Elems())), 0)
}

func TestContains(t *testing.T) {
	emptySet := New[int]()
	noneEmpty := FromSlice([]int{1, 2, 3})

	assert.False(t, emptySet.Contains(1))
	assert.False(t, noneEmpty.Contains(5))
	assert.True(t, noneEmpty.Contains(1))
}
