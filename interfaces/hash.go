package interfaces

import "hash/fnv"

type Hash[T any] interface {
	Hash(T) uint64
}

type HashString struct{}

func (hash *HashString) Hash(s string) uint64 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return uint64(h.Sum32())
}

type HashInt struct{}

func (hash *HashInt) Hash(i int) uint64 {
	return uint64(i)
}
