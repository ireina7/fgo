package generic

import (
	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/interfaces/collection"
	"github.com/ireina7/fgo/structs/option"
	"github.com/ireina7/fgo/structs/tuple"
	"github.com/ireina7/fgo/types"
	"github.com/zyedidia/generic/hashmap"
)

type HashMapKind types.Kind

type HashMap[K, V any] struct {
	hmap *hashmap.Map[K, V]
}

func (HashMap[K, V]) Kind(HashMapKind) {}
func (HashMap[K, V]) ElemType(V)       {}

func Make[K, V any](
	eq interfaces.Eq[K, K],
	hash interfaces.Hash[K],
	capacity uint64,
) HashMap[K, V] {
	return HashMap[K, V]{
		hmap: hashmap.New[K, V](capacity, eq.Equal, hash.Hash),
	}
}

func (hm HashMap[K, V]) Set(key K, v V) {
	hm.hmap.Put(key, v)
}

func (hm HashMap[K, V]) Get(key K) option.Option[V] {
	v, exist := hm.hmap.Get(key)
	if !exist {
		return option.Nothing[V]()
	}
	return option.Just(v)
}

func (hm HashMap[K, V]) Delete(key K) {
	hm.hmap.Remove(key)
}

func (hm HashMap[K, V]) Clear() {
	hm.hmap.Clear()
}

func (hm HashMap[K, V]) Len() int {
	return hm.hmap.Size()
}

func (hm HashMap[K, V]) HasKey(key K) bool {
	return !option.IsNone(hm.Get(key))
}

type hashMapIter[K, V any] struct {
	hm HashMap[K, V]
	ch chan tuple.Tuple2[K, V]
}

func (iter *hashMapIter[K, V]) Next() option.Option[tuple.Tuple2[K, V]] {
	for x := range iter.ch {
		return option.Just(x)
	}
	return option.Nothing[tuple.Tuple2[K, V]]()
}

func (hm HashMap[K, V]) Iter() collection.Iterator[tuple.Tuple2[K, V]] {
	ch := make(chan tuple.Tuple2[K, V])
	go func() {
		hm.hmap.Each(func(key K, v V) {
			ch <- tuple.Tuple2[K, V]{
				A: key,
				B: v,
			}
		})
		close(ch)
	}()
	return &hashMapIter[K, V]{
		hm: hm,
		ch: ch,
	}
}
