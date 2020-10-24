package queue

import (
	"github.com/nsnikhil/go-datastructures/list"
)

type LinkedQueue struct {
	ll *list.LinkedList
}

func NewLinkedQueue() (*LinkedQueue, error) {
	ll, err := list.NewLinkedList()
	if err != nil {
		return nil, err
	}

	return &LinkedQueue{
		ll: ll,
	}, nil
}

func (lq *LinkedQueue) Add(e interface{}) error {
	return lq.ll.AddLast(e)
}

func (lq *LinkedQueue) Remove() (interface{}, error) {
	//TODO FIX "LIST IS EMPTY ERROR"
	return lq.ll.RemoveFirst()
}

func (lq *LinkedQueue) Peek() (interface{}, error) {
	//TODO CHECK IF EMPTY QUEUE SHOULD RETURN ERROR
	return lq.ll.GetFirst(), nil
}

func (lq *LinkedQueue) Empty() bool {
	return lq.ll.IsEmpty()
}

func (lq *LinkedQueue) Size() int {
	return lq.ll.Size()
}

func (lq *LinkedQueue) Clear() {
	lq.ll.Clear()
}
