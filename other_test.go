package gunc

import (
	"slices"
	"testing"
)

func TestReverse(t *testing.T) {
	s := Range(1, 5)
	s = Reverse(s)
	target := []int{5, 4, 3, 2, 1}
	if !slices.Equal(s, target) {
		errorTargetVsExpected(t, "reverse", s, target)
	}
}
