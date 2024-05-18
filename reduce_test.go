package gunc

import (
	"testing"
)

func TestReduce(t *testing.T) {
	s1 := Range(0, 5)
	sr := Reduce(s1, func(a, b int) int { return a + b })
	checkTargetVal(t, "reduce #1", sr, 15)

	s2 := Range(1, 5)
	sr = Reduce(s2, func(a, b int) int { return b - a })
	checkTargetVal(t, "reduce #2", sr, 3)

	s3 := []int{1, 5, 15, 25}
	sr = ReduceRight(s3, func(a, b int) int { return a - b })
	checkTargetVal(t, "reduceright #1", sr, 4)

	//Scan
	scanr := Scan(s1, func(a, b int) int { return a + b })
	checkTargetSlice(t, "scan #1", scanr, []int{0, 1, 3, 6, 10, 15})

	scanr = Scan(s2, func(a, b int) int { return b - a })
	checkTargetSlice(t, "scan #2", scanr, []int{1, 1, 2, 2, 3})
}
