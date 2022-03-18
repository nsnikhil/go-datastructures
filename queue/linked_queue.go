package queue

import (
	"github.com/nsnikhil/erx"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/nsnikhil/go-datastructures/list"
)

type LinkedQueue[T comparable] struct {
	ll *list.LinkedList[T]
}

func NewLinkedQueue[T comparable]() *LinkedQueue[T] {
	return &LinkedQueue[T]{
		ll: list.NewLinkedList[T](),
	}
}

func (lq *LinkedQueue[T]) Add(e T) {
	lq.ll.AddLast(e)
}

func (lq *LinkedQueue[T]) Remove() (T, error) {
	v, err := lq.ll.RemoveFirst()
	if err != nil {
		return internal.ZeroValueOf[T](), erx.WithArgs(erx.Kind("LinkedQueue.Remove"), err)
	}

	return v, nil
}

func (lq *LinkedQueue[T]) Peek() (T, error) {
	v, err := lq.ll.GetFirst()
	if err != nil {
		return internal.ZeroValueOf[T](), erx.WithArgs(erx.Kind("LinkedQueue.Peek"), err)
	}

	return v, nil
}

func (lq *LinkedQueue[T]) Empty() bool {
	return lq.ll.IsEmpty()
}

func (lq *LinkedQueue[T]) Size() int64 {
	return lq.ll.Size()
}

func (lq *LinkedQueue[T]) Clear() {
	lq.ll.Clear()
}
