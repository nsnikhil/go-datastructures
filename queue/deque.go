package queue

import (
	"github.com/nsnikhil/erx"
	"github.com/nsnikhil/go-datastructures/internal"
)

type Deque[T comparable] struct {
	*LinkedQueue[T]
}

func NewDeque[T comparable]() *Deque[T] {
	return &Deque[T]{
		NewLinkedQueue[T](),
	}
}

func (dq *Deque[T]) AddFirst(e T) {
	dq.ll.AddFirst(e)
}

func (dq *Deque[T]) RemoveLast() (T, error) {
	v, err := dq.ll.RemoveLast()
	if err != nil {
		return internal.ZeroValueOf[T](), erx.WithArgs(erx.Kind("Deque.RemoveLast"), err)
	}

	return v, nil
}

func (dq *Deque[T]) PeekLast() (T, error) {
	v, err := dq.ll.GetLast()
	if err != nil {
		return internal.ZeroValueOf[T](), erx.WithArgs(erx.Kind("Deque.PeekLast"), err)
	}

	return v, nil
}
