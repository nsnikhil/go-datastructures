package queue

import (
	"github.com/nsnikhil/erx"
	"github.com/nsnikhil/go-datastructures/internal"
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
	return bq.removeBlocking(time.Duration(max))
}

func (bq *BlockingQueue[T]) RemoveWithTimeout(duration time.Duration) (T, error) {
	e, err := bq.removeBlocking(duration)

	return e, err
}

func (bq *BlockingQueue[T]) removeBlocking(duration time.Duration) (T, error) {
	c := make(chan T)
	e := make(chan error)

	go func() {
		for bq.ll.Size() == 0 {

		}

		v, err := bq.ll.RemoveFirst()
		if err != nil {
			e <- erx.WithArgs(erx.Kind("BlockingQueue.removeBlocking"), err)
			return
		}

		c <- v

	}()

	select {
	case v := <-c:
		return v, nil
	case err := <-e:
		return internal.ZeroValueOf[T](), err
	case <-time.After(duration):
		return internal.ZeroValueOf[T](), blockingQueueTimedOutError("BlockingQueue.removeBlocking")
	}

}
