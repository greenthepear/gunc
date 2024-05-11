package gunc

// Apply applies function fn to each element of the slice and returns it.
//
// Return can be of different type than the slice.
func Apply[S ~[]E, E any, R any](slice S, fn func(E) R) []R {
	r := make([]R, len(slice))
	for i, e := range slice {
		r[i] = fn(e)
	}
	return r
}

// Like [Apply], but function fn takes the index of the element as a parameter.
func ApplyWithIndex[S ~[]E, E any, R any](slice S, fn func(int, E) R) []R {
	r := make([]R, len(slice))
	for i, e := range slice {
		r[i] = fn(i, e)
	}
	return r
}
