package queue

import (
	"errors"
	"github.com/nsnikhil/erx"
)

var blockingQueueTimedOutError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("blockingQueueTimedOutError"),
		operation,
		errors.New("timed out"),
	)
}
