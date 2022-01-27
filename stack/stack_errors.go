package stack

import (
	"errors"
	"github.com/nsnikhil/erx"
)

var emptyStackError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("emptyStackError"),
		operation,
		errors.New("stack is empty"),
	)
}
