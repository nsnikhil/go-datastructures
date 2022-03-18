package set

import (
	"errors"
	"github.com/nsnikhil/erx"
)

var emptySetError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("emptySetError"),
		operation,
		errors.New("set is empty"),
	)
}

var emptyIteratorError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("emptyIteratorError"),
		operation,
		errors.New("iterator is empty"),
	)
}

var emptyArgsListError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("emptyArgsListError"),
		operation,
		errors.New("argument list is empty"),
	)
}
