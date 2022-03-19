package stream

import (
	"github.com/nsnikhil/go-datastructures/base"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/consumer"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/functions/predicate"
	"github.com/nsnikhil/go-datastructures/functions/supplier"
)

type Stream[T comparable] interface {
	AllMatch(p predicate.Predicate[T]) bool

	AnyMatch(p predicate.Predicate[T]) bool

	//Collect(c collector.Collector[T]) T

	Count() int

	Distinct() Stream[T]

	DropWhile(p predicate.Predicate[T]) Stream[T]

	TakeWhile(p predicate.Predicate[T]) Stream[T]

	Empty() bool

	Filter(p predicate.Predicate[T]) Stream[T]

	Iterator() iterator.Iterator[T]

	Generate(s supplier.Supplier[T]) Stream[T]

	Iterate(s T, uo operator.UnaryOperator[T]) Stream[T]

	Limit(c int) Stream[T]

	//Map(f function.Function[T]) Stream[T]

	Max(c comparator.Comparator[T]) base.Optional[T]

	Min(c comparator.Comparator[T]) base.Optional[T]

	Of(e ...T) Stream[T]

	Peek(c consumer.Consumer[T]) Stream[T]

	//Reduce(bo operator.BinaryOperator[T]) base.Optional[T]

	Skip(n int) Stream[T]

	Sorted(c comparator.Comparator[T]) Stream[T]
}
