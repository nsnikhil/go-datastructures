package stack

import (
	"github.com/nsnikhil/erx"
	"github.com/nsnikhil/go-datastructures/list"
)

type Stack[T comparable] struct {
	ll *list.LinkedList[T]
}

//TODO: SHOULD THE CONSTRUCTOR ACCEPT VAR ARGS?
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		ll: list.NewLinkedList[T](),
	}
}

func (s *Stack[T]) Push(e T) {
	s.ll.AddFirst(e)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		return *new(T), emptyStackError("Stack.Pop")
	}

	res, err := s.ll.RemoveFirst()
	if err != nil {
		return *new(T), erx.WithArgs(erx.Operation("Stack.Pop"), err)
	}

	return res, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.Empty() {
		return *new(T), emptyStackError("Stack.Peek")
	}

	res, err := s.ll.GetFirst()
	if err != nil {
		return *new(T), erx.WithArgs(erx.Operation("Stack.Peek"), err)
	}

	return res, nil
}

func (s *Stack[T]) Empty() bool {
	return s.ll.IsEmpty()
}

func (s *Stack[T]) Size() int64 {
	return s.ll.Size()
}

func (s *Stack[T]) Clear() {
	s.ll.Clear()
}
