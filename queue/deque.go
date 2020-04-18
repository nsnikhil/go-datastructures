package queue

type Deque struct {
	*LinkedQueue
}

func NewDeque() (*Deque, error) {
	lq, err := NewLinkedQueue()
	if err != nil {
		return nil, err
	}

	return &Deque{
		lq,
	}, nil
}

func (dq *Deque) AddFirst(e interface{}) error {
	return dq.ll.AddFirst(e)
}

func (dq *Deque) RemoveLast() (interface{}, error) {
	return dq.ll.RemoveLast()
}

func (dq *Deque) PeekLast() (interface{}, error) {
	//TODO CHECK IF EMPTY QUEUE SHOULD RETURN ERROR
	return dq.ll.GetLast(), nil
}
