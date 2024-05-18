package gunc

// Apply applies function fn to each element of the slice and returns it.
//
// The returning slice can be of a differnt type.
func Apply[S ~[]E, E any, R any](slice S, fn func(E) R) []R {
	r := make([]R, len(slice))
	for i, e := range slice {
		r[i] = fn(e)
	}
	return r
}

// Like [Apply], but function fn takes the index of the element
// as a parameter.
func ApplyWithIndex[S ~[]E, E any, R any](slice S, fn func(int, E) R) []R {
	r := make([]R, len(slice))
	for i, e := range slice {
		r[i] = fn(i, e)
	}
	return r
}

// Like [Apply] but works on the slice directly/in-place.
func ApplyInPlace[S ~[]E, E any](slice *S, fn func(E) E) {
	for i, e := range *slice {
		(*slice)[i] = fn(e)
	}
}

// Like [ApplyInPlace], but function fn takes the index of the element
// as a parameter.
func ApplyInPlaceWithIndex[S ~[]E, E any](slice *S, fn func(int, E) E) {
	for i, e := range *slice {
		(*slice)[i] = fn(i, e)
	}
}
