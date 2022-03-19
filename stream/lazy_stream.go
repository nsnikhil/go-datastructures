package stream

import (
	"github.com/nsnikhil/go-datastructures/base"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/consumer"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/functions/predicate"
	"github.com/nsnikhil/go-datastructures/functions/supplier"
	"github.com/nsnikhil/go-datastructures/queue"
	"github.com/nsnikhil/go-datastructures/utils"
)

type LazyStream[T comparable] struct {
	operations queue.Queue[T]
}

func NewLazyStream[T comparable]() Stream[T] {
	return &LazyStream[T]{
		operations: queue.NewDeque[T](),
	}
}

func (ls *LazyStream[T]) AllMatch(p predicate.Predicate[T]) bool {
	return false
}

func (ls *LazyStream[T]) AnyMatch(p predicate.Predicate[T]) bool {
	return false
}

//func (ls *LazyStream[T]) Collect(c collector.Collector[T]) T {
//	return nil
//}

func (ls *LazyStream[T]) Count() int {
	return utils.InvalidIndex
}

func (ls *LazyStream[T]) Distinct() Stream[T] {
	return nil
}

func (ls *LazyStream[T]) DropWhile(p predicate.Predicate[T]) Stream[T] {
	return nil
}

func (ls *LazyStream[T]) TakeWhile(p predicate.Predicate[T]) Stream[T] {
	return nil
}

func (ls *LazyStream[T]) Empty() bool {
	return false
}

func (ls *LazyStream[T]) Filter(p predicate.Predicate[T]) Stream[T] {
	return nil
}

func (ls *LazyStream[T]) Iterator() iterator.Iterator[T] {
	return nil
}

func (ls *LazyStream[T]) Generate(s supplier.Supplier[T]) Stream[T] {
	return nil
}

func (ls *LazyStream[T]) Iterate(s T, uo operator.UnaryOperator[T]) Stream[T] {
	return nil
}

func (ls *LazyStream[T]) Limit(c int) Stream[T] {
	return nil
}

//func (ls *LazyStream[T]) Map(f function.Function[T]) Stream[T] {
//	return nil
//}

func (ls *LazyStream[T]) Max(c comparator.Comparator[T]) base.Optional[T] {
	return nil
}

func (ls *LazyStream[T]) Min(c comparator.Comparator[T]) base.Optional[T] {
	return nil
}

func (ls *LazyStream[T]) Of(e ...T) Stream[T] {
	return nil
}

func (ls *LazyStream[T]) Peek(c consumer.Consumer[T]) Stream[T] {
	return nil
}

//func (ls *LazyStream[T]) Reduce(bo operator.BinaryOperator) base.Optional[T] {
//	return nil
//}

func (ls *LazyStream[T]) Skip(n int) Stream[T] {
	return nil
}

func (ls *LazyStream[T]) Sorted(c comparator.Comparator[T]) Stream[T] {
	return nil
}
