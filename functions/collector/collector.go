package collector

import (
	"github.com/nsnikhil/go-datastructures/functions/consumer"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/functions/supplier"
)

type Collector[T any, K any, R any] interface {
	Accumulator() consumer.BiConsumer[T, K]
	Combiner() operator.BinaryOperator[T, K, R]
	Supplier() supplier.Supplier[T]
}
