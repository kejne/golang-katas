package chop_test

import (
	testee "github.com/kejne/golang-katas/kata02-karate-chop/chop"
	a "github.com/stretchr/testify/assert"
	"testing"
)

func TestChop(t *testing.T) {
	a.Equal(t, -1, testee.Chop(3, make([]int, 0)))
	a.Equal(t, -1, testee.Chop(3, []int{1}))
	a.Equal(t, 0, testee.Chop(1, []int{1}))
	a.Equal(t, 0, testee.Chop(1, []int{1, 3, 5}))
	a.Equal(t, 1, testee.Chop(3, []int{1, 3, 5}))
	a.Equal(t, 2, testee.Chop(5, []int{1, 3, 5}))
	a.Equal(t, -1, testee.Chop(0, []int{1, 3, 5}))
	a.Equal(t, -1, testee.Chop(2, []int{1, 3, 5}))
	a.Equal(t, -1, testee.Chop(4, []int{1, 3, 5}))
	a.Equal(t, -1, testee.Chop(6, []int{1, 3, 5}))
	a.Equal(t, 0, testee.Chop(1, []int{1, 3, 5, 7}))
	a.Equal(t, 1, testee.Chop(3, []int{1, 3, 5, 7}))
	a.Equal(t, 2, testee.Chop(5, []int{1, 3, 5, 7}))
	a.Equal(t, 3, testee.Chop(7, []int{1, 3, 5, 7}))
	a.Equal(t, -1, testee.Chop(0, []int{1, 3, 5, 7}))
	a.Equal(t, -1, testee.Chop(2, []int{1, 3, 5, 7}))
	a.Equal(t, -1, testee.Chop(4, []int{1, 3, 5, 7}))
	a.Equal(t, -1, testee.Chop(6, []int{1, 3, 5, 7}))
	a.Equal(t, -1, testee.Chop(8, []int{1, 3, 5, 7}))
}
