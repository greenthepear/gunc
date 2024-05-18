package gunc

import (
	"testing"
)

func TestFilter(t *testing.T) {
	s1 := Range(0, 10)
	s1 = Filter(s1, func(e int) bool { return e%3 != 0 })
	checkTargetSlice(t, "filter #1", s1, []int{1, 2, 4, 5, 7, 8, 10})

	s2 := Range(-100, 100)
	s2 = Filter(s2, func(e int) bool { return e >= -5 && e <= 5 })
	checkTargetSlice(
		t, "filter #2", s2, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})
}
