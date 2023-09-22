# fgo
Functional GOlang

# Features
- Algebraic Data Types
  + Option
  + Result
  + ...
- Higher-Kinded Types
- Typeclasses
  + Monoid
  + Functor
  + ...


# Example
## Monoid
```go
type Monoid[A any] interface {
    Empty() A
    Combine(A, A) A
}

type folderExample[A any] struct {
    monoid Monoid[A]
}

func (example folderExample) foldl(xs []A, acc func([]A, A) []A) []A {
    m := example.monoid
    ans := m.Empty()
    for _, x := range xs {
        ans = m.Combine(ans, x)
    }
    return ans
}
```
- Monoid whith HKT(Just for fun ;)
```go
func foldl[F_, A any](
    monoid Monoid[HKT[F_, A]],
    xs ...HKT[F_, A],
) HKT[F_, A] {
    zs := monoid.Empty()
    for _, ys := range xs {
	    zs = monoid.Combine(zs, ys)
    }
    return zs
}
```
## Functor
```go
type Functor[F_, A, B any] interface {
    Fmap(types.HKT[F_, A], func(A) B) types.HKT[F_, B]
}

```
