package linklist

import (
	"github.com/ireina7/fgo/interfaces/collection"
	"github.com/ireina7/fgo/structs/maybe"
	"github.com/ireina7/fgo/types"
	"github.com/ireina7/fgo/util"
)

type ListKind types.Type

type List[A any] struct {
	head ListNode[A]
}

func (List[A]) Kind(ListKind) {}
func (List[A]) ElemType(A)    {}

func Make[A any](xs ...A) List[A] {
	head := ListNode[A]{}
	node := &head
	for _, x := range xs {
		node.NextNode = &ListNode[A]{}
		node.NextNode.Value = x
		node = node.NextNode
	}
	return List[A]{head: *head.NextNode}
}

func (list List[A]) Iter() collection.Iterator[A] {
	return &listIter[A]{
		node: &list.head,
	}
}

func (list List[A]) Len() int {
	count := 0
	collection.For[A](list, func(x A) {
		count += 1
	})
	return count
}

func Clone[A any](xs List[A]) List[A] {
	return util.TODO[List[A]]("List.Clone")
}

func (list List[A]) Append(x A) List[A] {
	return util.TODO[List[A]]("List.Append")
}

// func Map[A, B any](list List[A], f func(A) B) List[B] {
// 	ans :=
// 	return nil
// }

type listIter[A any] struct {
	node *ListNode[A]
}

func (iter *listIter[A]) Next() maybe.Maybe[A] {
	if iter.node == nil {
		return maybe.From[A](nil)
	}
	val := iter.node.Value
	iter.node = iter.node.NextNode
	return maybe.From(&val)
}

type ListNode[A any] struct {
	Value    A
	NextNode *ListNode[A]
}
