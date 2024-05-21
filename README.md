[![Go Reference](https://pkg.go.dev/badge/github.com/greenthepear/gunc.svg)](https://pkg.go.dev/github.com/greenthepear/gunc)

**gunc** is a collection of functional programming and array programming utilities for Go slices.

No dependencies, just the standard library and generics!

# Features

No more for loops for simple range number slices
```go
s1 := gunc.Range(-3, 3) // [-3, -2, -1, 0, 1, 2, 3]
s2 := gunc.RangeWithStep(0, 1, 0.2) // [0.0, 0.2, 0.4, 0.6, 0.8, 1.0]
```
The classics: apply, filter, reduce

```go
// Also check out gunc.ApplyInPlace()!
s1 = gunc.Apply(s1, func(e int) int { return e + 3 }) // [0, 1, 2, 3, 4, 5, 6]
s1 = gunc.Filter(s1, func(e int) bool { return e%3 == 0 }) // [0, 3, 6]
res := gunc.Reduce(s1, func(a, b int) int { return a + b }) // 9
```

Want the functions to use indexes also? `ApplyWithIndex()`, `FilterWithIndex()` and `ReduceWithIndex()` got you covered. Wanna keep `Reduce()`'s elements? `Scan()` is here from APL land.

# License
Do what you want with this, to be fancy I chose the **0BSD** license.
