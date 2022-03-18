package heap

import (
	"errors"
	"github.com/nsnikhil/erx"
)

var emptyHeapError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("emptyHeapError"),
		operation,
		errors.New("heap is empty"),
	)
}

var emptyIteratorError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("emptyIteratorError"),
		operation,
		errors.New("iterator is empty"),
	)
}
