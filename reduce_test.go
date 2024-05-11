package gunc

import (
	"slices"
	"testing"
)

func TestReduce(t *testing.T) {
	s1 := Range(0, 5)
	sr := Reduce(s1, func(a, b int) int { return a + b })
	target := 15
	if sr != target {
		errorTargetVsExpected(t, "reduce #1", sr, target)
	}

	s2 := Range(1, 5)
	sr = Reduce(s2, func(a, b int) int { return a * b })
	target = 120
	if sr != target {
		errorTargetVsExpected(t, "reduce #2", sr, target)
	}

	//Scan

	scanr := Scan(s1, func(a, b int) int { return a + b })
	scanTarget := []int{0, 1, 3, 6, 10, 15}
	if !slices.Equal(scanr, scanTarget) {
		errorTargetVsExpected(t, "scan #1", scanr, scanTarget)
	}

	scanr = Scan(s2, func(a, b int) int { return b - a })
	scanTarget = []int{1, 1, 2, 2, 3}
	if !slices.Equal(scanr, scanTarget) {
		errorTargetVsExpected(t, "scan #2", scanr, scanTarget)
	}
}
