package base

import (
	"github.com/nsnikhil/go-datastructures/functions/consumer"
	"github.com/nsnikhil/go-datastructures/functions/runnable"
	"github.com/nsnikhil/go-datastructures/functions/supplier"
)

type Optional interface {
	Empty() bool

	Get() (interface{}, error)

	IfPresent(c consumer.Consumer)

	IfPresentOrElse(c consumer.Consumer, r runnable.Runnable)

	IsPresent() bool

	OrElse(e interface{}) interface{}

	OrElseGet(s supplier.Supplier) interface{}
}
