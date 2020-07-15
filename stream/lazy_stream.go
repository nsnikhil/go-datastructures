package stream

import (
	"github.com/nsnikhil/go-datastructures/base"
	"github.com/nsnikhil/go-datastructures/functions/collector"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/consumer"
	"github.com/nsnikhil/go-datastructures/functions/function"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/functions/predicate"
	"github.com/nsnikhil/go-datastructures/functions/supplier"
	"github.com/nsnikhil/go-datastructures/queue"
	"github.com/nsnikhil/go-datastructures/utils"
)

type LazyStream struct {
	operations queue.Queue
}

func NewLazyStream() (Stream, error) {
	q, err := queue.NewDeque()
	if err != nil {
		return nil, err
	}

	ls := &LazyStream{
		operations: q,
	}

	return ls, nil
}

func (ls *LazyStream) AllMatch(p predicate.Predicate) bool {
	return false
}

func (ls *LazyStream) AnyMatch(p predicate.Predicate) bool {
	return false
}

func (ls *LazyStream) Collect(c collector.Collector) interface{} {
	return nil
}

func (ls *LazyStream) Count() int {
	return utils.InvalidIndex
}

func (ls *LazyStream) Distinct() Stream {
	return nil
}

func (ls *LazyStream) DropWhile(p predicate.Predicate) Stream {
	return nil
}

func (ls *LazyStream) TakeWhile(p predicate.Predicate) Stream {
	return nil
}

func (ls *LazyStream) Empty() bool {
	return false
}

func (ls *LazyStream) Filter(p predicate.Predicate) Stream {
	return nil
}

func (ls *LazyStream) Iterator() iterator.Iterator {
	return nil
}

func (ls *LazyStream) Generate(s supplier.Supplier) Stream {
	return nil
}

func (ls *LazyStream) Iterate(s interface{}, uo operator.UnaryOperator) Stream {
	return nil
}

func (ls *LazyStream) Limit(c int) Stream {
	return nil
}

func (ls *LazyStream) Map(f function.Function) Stream {
	return nil
}

func (ls *LazyStream) Max(c comparator.Comparator) base.Optional {
	return nil
}

func (ls *LazyStream) Min(c comparator.Comparator) base.Optional {
	return nil
}

func (ls *LazyStream) Of(e ...interface{}) Stream {
	return nil
}

func (ls *LazyStream) Peek(c consumer.Consumer) Stream {
	return nil
}

func (ls *LazyStream) Reduce(bo operator.BinaryOperator) base.Optional {
	return nil
}

func (ls *LazyStream) Skip(n int) Stream {
	return nil
}

func (ls *LazyStream) Sorted(c comparator.Comparator) Stream {
	return nil
}
