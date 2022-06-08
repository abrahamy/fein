package set_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/abrahamy/fein/set"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := set.New[int]()
	assert.NotNil(t, s)
}

func TestFromSlice(t *testing.T) {
	s := set.FromSlice([]int{1, 2, 2, 3, 1, 1})
	assert.Equal(t, s.Len(), 3)
	assert.ElementsMatch(t, s.Members(), []int{1, 2, 3})
}

func TestLen(t *testing.T) {
	s := set.New[int]()
	assert.Equal(t, s.Len(), 0)

	s.Add(1, 2, 3)
	assert.Equal(t, s.Len(), 3)

	assert.Equal(t, set.FromSlice([]int{4, 5, 5}).Len(), 2)
}

func TestAdd(t *testing.T) {
	s := set.New[int]()
	s.Add(1, 2, 3)
	assert.Equal(t, s.Len(), 3)

	s.Add(4, 5)
	assert.Equal(t, s.Len(), 5)

	assert.ElementsMatch(t, s.Members(), []int{1, 2, 3, 4, 5})
}

func TestMembers(t *testing.T) {
	emptySet := set.New[int]()
	intSet := set.FromSlice([]int{1, 2, 2, 3, 4, 4, 3})
	assert.Equal(t, len(emptySet.Members()), 0)
	assert.Equal(t, len(intSet.Members()), 4)
	assert.ElementsMatch(t, intSet.Members(), []int{1, 2, 3, 4})
}

func TestUnion(t *testing.T) {
	emptySet := set.New[int]()
	oneTwoThree := set.FromSlice([]int{1, 2, 3})
	fourFive := set.FromSlice([]int{4, 5})
	sixSevenEight := set.FromSlice([]int{6, 7, 8})

	assert.ElementsMatch(t, emptySet.Union(oneTwoThree).Members(), oneTwoThree.Members())
	assert.ElementsMatch(t, oneTwoThree.Union(fourFive).Members(), []int{1, 2, 3, 4, 5})
	assert.ElementsMatch(t, oneTwoThree.Union(fourFive, sixSevenEight).Members(), []int{1, 2, 3, 4, 5, 6, 7, 8})
}

func TestString(t *testing.T) {
	emptySet := set.New[int]()
	noneEmpty := set.FromSlice([]int{1})
	assert.Equal(t, strings.Compare(emptySet.String(), fmt.Sprintf("Set%v", emptySet.Members())), 0)
	assert.Equal(t, strings.Compare(noneEmpty.String(), fmt.Sprintf("Set%v", noneEmpty.Members())), 0)
}

func TestContains(t *testing.T) {
	emptySet := set.New[int]()
	noneEmpty := set.FromSlice([]int{1, 2, 3})

	assert.False(t, emptySet.Contains(1))
	assert.False(t, noneEmpty.Contains(5))
	assert.True(t, noneEmpty.Contains(1))
}
