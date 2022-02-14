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
)

type Stream[T comparable] interface {
	AllMatch(p predicate.Predicate[T]) bool

	AnyMatch(p predicate.Predicate[T]) bool

	Collect(c collector.Collector[T]) interface{}

	Count() int

	Distinct() Stream

	DropWhile(p predicate.Predicate[T]) Stream

	TakeWhile(p predicate.Predicate[T]) Stream

	Empty() bool

	Filter(p predicate.Predicate[T]) Stream

	Iterator() iterator.Iterator

	Generate(s supplier.Supplier) Stream

	Iterate(s interface{}, uo operator.UnaryOperator) Stream

	Limit(c int) Stream

	Map(f function.Function) Stream

	Max(c comparator.Comparator) base.Optional

	Min(c comparator.Comparator) base.Optional

	Of(e ...interface{}) Stream

	Peek(c consumer.Consumer) Stream

	Reduce(bo operator.BinaryOperator) base.Optional

	Skip(n int) Stream

	Sorted(c comparator.Comparator) Stream
}
