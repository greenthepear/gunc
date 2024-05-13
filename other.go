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

func Pipeline[S ~[]E, E any, F PipeFunction[S, E]](param S, functions ...func() S) S {
	for _, f := range functions {
		f1, ok := any(f).(func(S) S)
		if ok {
			param = f1(param)
			continue
		}
		f2, ok := any(f).(func(S, func(E) E) S)
		if ok {
			param = 
		}
		f3, ok3 := any(f).(func(S, func(int, E) E) S)
		param = f(param)
	}
	return param
}
