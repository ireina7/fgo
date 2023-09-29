package generic

import (
	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/structs/option"
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
