# fgo
Functional GO!

![cover](./cover.jpeg)

## Never use default struct constructor since golang has zero values for all types which may cause implicit invalid state!

# Features
## Algebraic Data Types
```go
type (
    Maybe, Result, Slice[T], HashMap[K, V], 
    List[T], SkipList[T], Ref[T](non-null reference), 
    Ptr[T], Lazy[T], Tuple[A, B]...
)
```
## Higher-Kinded Types
`types.HKT[Kind, A]`

## Typeclasses
```go
type (
    Monoid, Functor, Eq, Ord, 
    Iterable, Implement, Monad, 
    Foldable, Enum, Apply...
)
```

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
