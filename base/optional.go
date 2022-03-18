package base

import (
	"github.com/nsnikhil/go-datastructures/functions/consumer"
	"github.com/nsnikhil/go-datastructures/functions/runnable"
	"github.com/nsnikhil/go-datastructures/functions/supplier"
)

type Optional[T comparable] interface {
	Empty() bool

	Get() (T, error)

	IfPresent(c consumer.Consumer[T])

	IfPresentOrElse(c consumer.Consumer[T], r runnable.Runnable)

	IsPresent() bool

	OrElse(e T) T

	OrElseGet(s supplier.Supplier[T]) T
}
