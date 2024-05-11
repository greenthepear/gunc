package gunc

import (
	"slices"
	"testing"
)

func errorTargetVsExpected(
	t *testing.T, whatsWrong string, result any, target any) {
	t.Errorf(
		"Worng after %s!\nShould be:\t%v\nis:\t%v",
		whatsWrong, target, result)
}

func TestApply(t *testing.T) {
	s := CreateWithStep(0, 10, 0.5)
	rs := string(Apply(s, func(e float64) rune {
		return 'a' + rune(e)
	}))
	target := "aabbccddeeffgghhiijjk"
	if rs != target {
		errorTargetVsExpected(t, "apply #1", rs, target)
	}

	s2 := CreateWithStep(100, 0, 10)
	s2 = ApplyWithIndex(s2, func(index int, elem int) int {
		return index - 2
	})
	target2 := []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8}
	if !slices.Equal(s2, target2) {
		errorTargetVsExpected(t, "apply #2", s2, target2)
	}
}
