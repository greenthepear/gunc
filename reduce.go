package gunc

// Return value after applying dyadic function fn to each neighboring element.
func Reduce[S ~[]E, E any](slice S, fn func(E, E) E) E {
	var r E = slice[0]
	for i := 1; i < len(slice)-1; i++ {
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
