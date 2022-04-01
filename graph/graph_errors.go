package graph

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/erx"
)

var nodeNotFoundError = func(node interface{}, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("nodeNotFoundError"),
		operation,
		fmt.Errorf("node %v not found in the graph", node),
	)
}

var edgeNotFoundError = func(from, to interface{}, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("edgeNotFoundError"),
		operation,
		fmt.Errorf("edge %v to %v not found in the graph", from, to),
	)
}

var emptyIteratorError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("emptyIteratorError"),
		operation,
		errors.New("iterator is empty"),
	)
}

var pathNotFoundError = func(from, to interface{}, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("pathNotFoundError"),
		operation,
		fmt.Errorf("path %v to %v not found in the graph", from, to),
	)
}
