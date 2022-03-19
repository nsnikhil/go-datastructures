package tree

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/erx"
)

var emptyTreeError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("emptyTreeError"),
		operation,
		errors.New("tree is empty"),
	)
}

var isNotBinaryTreeError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("isNotBinaryTreeError"),
		operation,
		errors.New("tree is not binary tree"),
	)
}

//TODO: REMOVE INTERFACE FROM HERE
var elementNotFoundError = func(element interface{}, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("elementNotFoundError"),
		operation,
		fmt.Errorf("element %v not found in the tree", element),
	)
}

//TODO: REMOVE INTERFACE FROM HERE
var noPreOrderSuccessorError = func(element interface{}, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("noPreOrderSuccessorError"),
		operation,
		fmt.Errorf("no pre order successor found for %v", element),
	)
}

//TODO: REMOVE INTERFACE FROM HERE
var noPostOrderSuccessorError = func(element interface{}, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("noPostOrderSuccessorError"),
		operation,
		fmt.Errorf("no post order successor found for %v", element),
	)
}

//TODO: REMOVE INTERFACE FROM HERE
var noInOrderSuccessorError = func(element interface{}, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("noInOrderSuccessorError"),
		operation,
		fmt.Errorf("no in order successor found for %v", element),
	)
}

//TODO: REMOVE INTERFACE FROM HERE
var noLevelOrderSuccessorError = func(element interface{}, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("noLevelOrderSuccessorError"),
		operation,
		fmt.Errorf("no level order successor found for %v", element),
	)
}
