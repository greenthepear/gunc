package gunc

// Filter creates a new slice with only the values for which fn returns true.
func Filter[S ~[]E, E any](slice S, fn func(E) bool) S {
	r := make(S, 0)
	for _, e := range slice {
		if fn(e) {
			r = append(r, e)
		}
	}
	return r
}

// Like [Filter], but function fn takes the index of the element as a parameter.
func FilterWithIndex[S ~[]E, E any](slice S, fn func(int, E) bool) S {
	r := make(S, 0)
	for i, e := range slice {
		if fn(i, e) {
			r = append(r, e)
		}
	}
	return r
}
