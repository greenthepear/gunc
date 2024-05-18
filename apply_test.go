package gunc

import (
	"slices"
	"testing"
)

func errorTargetVsExpected(
	t *testing.T, whatsWrong string, result any, target any) {
	t.Errorf(
		"Worng after %s!\nshould be:\t%v\nis:\t\t%v",
		whatsWrong, target, result)
}

func checkTargetVal[E comparable](
	t *testing.T, whatsAreWeChecking string, result E, target E) {
	if result != target {
		errorTargetVsExpected(t, whatsAreWeChecking, result, target)
	}
}

func checkTargetSlice[S ~[]E, E comparable](
	t *testing.T, whatsAreWeChecking string, result S, target S) {
	if !slices.Equal(result, target) {
		errorTargetVsExpected(t, whatsAreWeChecking, result, target)
	}
}

func TestApply(t *testing.T) {
	s := RangeWithStep(0, 10, 0.5)
	rs := string(Apply(s, func(e float64) rune {
		return 'a' + rune(e)
	}))
	checkTargetVal(t, "apply #1", rs, "aabbccddeeffgghhiijjk")

	s2 := RangeWithStep(100, 0, 10)
	s2 = ApplyWithIndex(s2, func(index int, elem int) int {
		return index - 2
	})
	checkTargetSlice(
		t, "apply #2", s2, []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8})
}
