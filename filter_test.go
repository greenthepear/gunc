package gunc

import (
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	s1 := Create(0, 10)
	s1 = Filter(s1, func(e int) bool { return e%3 != 0 })
	target := []int{1, 2, 4, 5, 7, 8, 10}
	if !slices.Equal(s1, target) {
		t.Errorf(
			"Worng after apply!\nShould be:\t%v\nis:\t%v",
			target, s1)
	}

	s2 := Create(-100, 100)
	s2 = Filter(s2, func(e int) bool { return e >= -5 && e <= 5 })
	target = []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	if !slices.Equal(s2, target) {
		t.Errorf(
			"Worng after apply!\nShould be:\t%v\nis:\t%v",
			target, s2)
	}
}
