package gunc

import (
	"testing"
)

func TestReverse(t *testing.T) {
	s := Range(1, 5)
	s = Reverse(s)
	checkTargetSlice(t, "reverse", s, []int{5, 4, 3, 2, 1})
}
