# fgo
Functional GO!

![cover](./cover.jpeg)

# Features
- Algebraic Data Types
  + Option
  + Result
  + Slice
  + HashMap
  + List
  + SkipList
  + Ref
  + Ptr
  + Lazy
  + Tuple
  + ...
- Higher-Kinded Types
- Typeclasses
  + Monoid
  + Functor
  + Eq
  + Ord
  + Iterable
  + Implement
  + Monad
  + Foldable
  + Enum
  + Apply
  + ...


# Example
## Monoid
```go
type Monoid[A any] interface {
    Empty() A
    Combine(A, A) A
}

type foldExample[A any] struct {
    monoid Monoid[A]
}

func (example foldExample[A]) foldl(xs []A, acc func([]A, A) []A) []A {
    m := example.monoid
    ans := m.Empty()
    for _, x := range xs {
        ans = m.Combine(ans, x)
    }
    return ans
}
```
## Functor
```go
type Functor[F_, A, B any] interface {
    Fmap(types.HKT[F_, A], func(A) B) types.HKT[F_, B]
}

```
