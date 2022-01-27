package queue

import (
	"errors"
	"time"
)

const max = 1<<63 - 62135596801

type BlockingQueue[T comparable] struct {
	*LinkedQueue[T]
}

func NewBlockingQueue[T comparable]() *BlockingQueue[T] {
	return &BlockingQueue[T]{LinkedQueue: NewLinkedQueue[T]()}
}

func (bq *BlockingQueue[T]) Remove() (T, error) {
	return removeBlocking(bq, time.Duration(max))
}

func (bq *BlockingQueue[T]) RemoveWithTimeout(duration time.Duration) (T, error) {
	e, err := removeBlocking(bq, duration)

	return e, err
}

func removeBlocking[T comparable](bq *BlockingQueue[T], duration time.Duration) (T, error) {
	c := make(chan T)
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
		return *new(T), err
	case <-time.After(duration):
		return *new(T), errors.New("timed out")
	}

}
