# fgo
Functional GO!

![cover](./cover.jpeg)

## Never use default struct constructor since golang has zero values for all types which may cause implicit invalid state!

# Features
## Algebraic Data Types
```go
type (
    Maybe[T], Result[T], Slice[T], HashMap[K, V], 
    List[T], SkipList[T], Ref[T](non-null reference), 
    Numbers, PriorityQueue[T](max heap), Set[T],
    Ptr[T], Lazy[T], Tuple[A, B]...
)
```
## Higher-Kinded Types
`types.HKT[Kind, A]`

## Typeclasses
```go
type (
    Monoid, Functor, Eq, Ord, Hash
    Iterable, Implement, Monad, Default,
    Foldable, Enum, Logging, Apply, Async...
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
// functor.go
type Functor[F_ Kind, A, B any] interface {
    Fmap(types.HKT[F_, A], func(A) B) types.HKT[F_, B]
}

// unsafe_functor.go
package unsafe
type Functor[F_ Kind] interface {
    Fmap(types.HKT[F_, any], func(any) any) types.HKT[F_, any]
}

type FromUnsafe[F_ Kind, A, B any] struct {
    hkt.Pipe[F_, A, B]
    unsafe Functor[F_]
}

func (self FromUnsafe[F_, A, B]) Fmap(ma types.HKT[F_, A], f func(A) B) types.HKT[F_, B] {
    mb := self.unsafe.Fmap(self.Boxed(ma), func(x any) any {
        return f(x.(A))
    })
    return self.Unboxed(mb)
}
```
