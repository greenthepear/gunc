package gunc

// A number is a number
type Number interface {
	float32 | float64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64
}

// Creates a slice of numbers with a specific step
// in an inclusive range [from, to]. Supports descending ranges.
//
// Panics if step is negative
// (should still be positive for descending ranges).
func RangeWithStep[N Number](from N, to N, step N) []N {
	r := make([]N, 0)
	if step < 0 {
		panic("negative step would cause an infinite loop")
	}
	if from < to {
		for i := from; i <= to; i += step {
			r = append(r, i)
		}
	} else {
		for i := from; i >= to; i -= step {
			r = append(r, i)
		}
	}
	return r
}

// Alias for [RangeWithStep](from, to, 1).
func Range[N Number](from N, to N) []N {
	return RangeWithStep(from, to, 1)
}
