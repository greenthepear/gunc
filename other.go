package gunc

import "slices"

// Reverse wraps [slices.Reverse], but returns the slice instead.
func Reverse[S ~[]E, E any](s S) S {
	slices.Reverse(s)
	return s
}
