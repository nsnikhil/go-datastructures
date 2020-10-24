package stack

import (
	"github.com/nsnikhil/go-datastructures/list"
)

type Stack struct {
	ll *list.LinkedList
}

func NewStack() (*Stack, error) {
	ll, err := list.NewLinkedList()
	if err != nil {
		return nil, err
	}

	return &Stack{
		ll: ll,
	}, nil
}

func (s *Stack) Push(e interface{}) error {
	return s.ll.AddFirst(e)
}

func (s *Stack) Pop() (interface{}, error) {
	return s.ll.RemoveFirst()
}

func (s *Stack) Peek() (interface{}, error) {
	return s.ll.GetFirst(), nil
}

func (s *Stack) Empty() bool {
	return s.ll.IsEmpty()
}

func (s *Stack) Size() int {
	return s.ll.Size()
}

func (s *Stack) Clear() {
	s.ll.Clear()
}

func (s *Stack) Search(e interface{}) (int, error) {
	//TODO FIX "LIST IS EMPTY ERROR" & "ELEMENT NOT FOUND IN THE LIST"
	return s.ll.IndexOf(e)
}
