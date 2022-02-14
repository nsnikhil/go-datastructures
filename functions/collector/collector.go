package collector

import (
	"github.com/nsnikhil/go-datastructures/functions/consumer"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/functions/supplier"
)

type Collector[T comparable] interface {
	Accumulator() consumer.BiConsumer[T]
	Combiner() operator.BinaryOperator[T]
	Supplier() supplier.Supplier[T]
}
