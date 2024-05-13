package gunc

import "slices"

// Reverse wraps [slices.Reverse], but returns the slice instead.
func Reverse[S ~[]E, E any](s S) S {
	slices.Reverse(s)
	return s
}

type PipeFunction[S ~[]E, E any] interface {
	func(S) S |
		func(S, func(E) E) S |
		func(S, func(int, E) E) S
}

func Do[S ~[]E, E any, F PipeFunction[S, E]](f F) func(E) S {
	return func(param E) S {
		return f(param)
	}
}

func Pipeline[S ~[]E, E any, F PipeFunction[S, E]](param S, functions ...F) S {
	for _, f := range functions {
		param = f1(param)
	}
	return param
}
