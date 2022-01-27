package list

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/erx"
)

var invalidArgsError = func(msg string, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("invalidArgsError"),
		operation,
		errors.New(msg),
	)
}

var emptyIteratorError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("emptyIteratorError"),
		operation,
		errors.New("iterator is empty"),
	)
}

var emptyListError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("emptyListError"),
		operation,
		errors.New("list is empty"),
	)
}

var invalidIndexError = func(index int64, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("invalidIndexError"),
		operation,
		fmt.Errorf("invalid index %d", index),
	)
}

//TODO: REMOVE INTERFACE FROM HERE
var elementNotFoundError = func(element interface{}, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("elementNotFoundError"),
		operation,
		fmt.Errorf("element %v not found in the list", element),
	)
}
