package collector

import (
	"github.com/nsnikhil/go-datastructures/functions/consumer"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/functions/supplier"
)

type Collector interface {
	Accumulator() consumer.BiConsumer
	Combiner() operator.BinaryOperator
	Supplier() supplier.Supplier
}
