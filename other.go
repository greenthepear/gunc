package gunc

import "slices"

// Reverse wraps [slices.Reverse], but returns the slice instead.
func Reverse[S ~[]E, E any](s S) S {
	slices.Reverse(s)
	return s
}

type InnerFunction[S ~[]E, E any] interface {
	func(E) E | func(int, E) E
}

type PipeFunction[I InnerFunction[S, E], S ~[]E, E any] interface {
	func(S) S | func(S, I) S
}
type With[M PipeFunction[I, S, E], I InnerFunction[S, E], S ~[]E, E any] struct {
	main  M
	inner I
}

func (p With[M, I, S, E]) With(
	mainFunction M, innerFunction I) With[M, I, S, E] {
	return With[M, I, S, E]{mainFunction, innerFunction}
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
