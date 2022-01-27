package queue

type Deque[T comparable] struct {
	*LinkedQueue[T]
}

func NewDeque[T comparable]() *Deque[T] {
	return &Deque[T]{
		NewLinkedQueue[T](),
	}
}

func (dq *Deque[T]) AddFirst(e T) error {
	return dq.ll.AddFirst(e)
}

func (dq *Deque[T]) RemoveLast() (T, error) {
	return dq.ll.RemoveLast()
}

func (dq *Deque[T]) PeekLast() (T, error) {
	//TODO CHECK IF EMPTY QUEUE SHOULD RETURN ERROR
	return dq.ll.GetLast()
}
