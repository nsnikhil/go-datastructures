package queue

import (
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

func (lq *LinkedQueue[T]) Add(e T) error {
	return lq.ll.AddLast(e)
}

func (lq *LinkedQueue[T]) Remove() (T, error) {
	//TODO FIX "LIST IS EMPTY ERROR"
	return lq.ll.RemoveFirst()
}

func (lq *LinkedQueue[T]) Peek() (T, error) {
	//TODO CHECK IF EMPTY QUEUE SHOULD RETURN ERROR
	return lq.ll.GetFirst()
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
