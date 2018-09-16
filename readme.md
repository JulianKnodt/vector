# Vector

A simple vector library for golang, which includes basic operations like
+, -, •, x, o, linear interpolation, min/max of two vectors, (square) magnitude, distance
between, normalizing, scalar multiplication, powers of vectors, and whether a vector intersects
a triangle created by a set of three other vectors.

Based around:
```go
type Vec3 [3]float64
```

## Examples

```go
&Vec3{1.0, 2.0, 3.0}.UnitSet().SMulSet(math.PI).Dot(otherVector)
```

Functions suffixed with set modify the pointer that it is being applied to,
and should be used whenever possible. Functions not sufixed with set create a new
pointer that should have set used on it if it's not needed later.

## Contributions

- Ensure `make test` continues to pass
- Check benchmarks with `make bench`
- Stay uniform with existing API
