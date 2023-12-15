# Structs
fgo provides lots of useful struct types and many types implements fgo interfaces(say, Iterable).

| name      | description      |
|-----------|------------------|
| Slice[T]  | Additional slice abilities |
| List[T]   | Linked list of T |
| Maybe[T]  | T or None, like rust's Option |
| Result[T] | value T or error, like rust's Result<T, &dyn Error>|
| Ref[T]    | Non-Null reference |
| PriorityQueue[T] | max heap |