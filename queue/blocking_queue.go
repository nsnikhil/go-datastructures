package queue

import (
	"errors"
	"time"
)

const max = 1<<63 - 62135596801

type BlockingQueue struct {
	*LinkedQueue
}

func NewBlockingQueue() (*BlockingQueue, error) {
	lq, err := NewLinkedQueue()
	if err != nil {
		return nil, err
	}

	return &BlockingQueue{LinkedQueue: lq}, nil
}

func (bq *BlockingQueue) Remove() (interface{}, error) {
	return removeBlocking(bq, time.Duration(max))
}

func (bq *BlockingQueue) RemoveWithTimeout(duration time.Duration) (interface{}, error) {
	e, err := removeBlocking(bq, duration)

	return e, err
}

func removeBlocking(bq *BlockingQueue, duration time.Duration) (interface{}, error) {
	c := make(chan interface{})
	e := make(chan error)

	go func() {
		for bq.ll.Size() == 0 {

		}

		v, err := bq.ll.RemoveFirst()
		if err != nil {
			e <- err
			return
		}

		c <- v

	}()

	select {
	case v := <-c:
		return v, nil
	case err := <-e:
		return nil, err
	case <-time.After(duration):
		return nil, errors.New("timed out")
	}

}
