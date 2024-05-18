package gunc

// Return value after applying dyadic function fn to each neighboring element
// from left to right, where a is the element on the left and b on the right.
func Reduce[S ~[]E, E any](slice S, fn func(E, E) E) E {
	var r E = slice[0]
	for i := 1; i < len(slice); i++ {
		r = fn(r, slice[i])
	}
	return r
}

// Like [Reduce] but from right to left. a is the element on the right while
// b is the one on the left.
func ReduceRight[S ~[]E, E any](slice S, fn func(E, E) E) E {
	var r E = slice[len(slice)-1]
	for i := len(slice) - 2; i >= 0; i-- {
		r = fn(r, slice[i])
	}
	return r
}

// Reduce, but keep intermediate values. Based on [APL's scan].
//
// [APL's scan]: https://aplwiki.com/wiki/Scan
func Scan[S ~[]E, E any](slice S, fn func(E, E) E) S {
	for i := 1; i < len(slice); i++ {
		slice[i] = fn(slice[i-1], slice[i])
	}
	return slice
}

// Like [Scan], but function fn takes the index of the element
// as a parameter.
func ScanWithIndex[S ~[]E, E any](
	slice S, fn func(e1 E, i1 int, e2 E, i2 int) E) S {

	for i := 1; i < len(slice); i++ {
		slice[i] = fn(slice[i-1], i-1, slice[i], i)
	}
	return slice
}
