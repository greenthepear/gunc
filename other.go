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

func Pipeline[M PipeFunction[I, S, E], I InnerFunction[S, E], S ~[]E, E any](
	param S, functions ...With[M, I, S, E]) S {
	for _, f := range functions {
		mf := f.main
		f1, ok := any(mf).(func(S) S)
		if ok {
			param = f1(param)
			continue
		}
		f2, ok := any(mf).(func(S, func(E) E) S)
		if ok {
			param = f2(param, any(f.inner).(func(E) E))
			continue
		}
		f3, ok := any(f).(func(S, func(int, E) E) S)
		if ok {
			param = f3(param, any(f.inner).(func(int, E) E))
			continue
		}

		panic("function does not implement generic interface")
	}
	return param
}
