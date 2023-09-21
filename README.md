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

type FolderExample[A any] struct {
    monoid Monoid[A]
}

func (example FolderExample) foldl(xs []A, acc func([]A, A) []A) []A {
    m := example.monoid
    ans := m.Empty()
    for _, x := range xs {
        ans = m.Combine(ans, x)
    }
    return ans
}
```
