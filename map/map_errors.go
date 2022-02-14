package gmap

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/erx"
)

var emptyMapError = func(operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("emptyMapError"),
		operation,
		errors.New("map is empty"),
	)
}

var keyNotFoundError = func(key interface{}, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("keyNotFoundError"),
		operation,
		fmt.Errorf("key %v not found in the map", key),
	)
}

var valueMisMatchError = func(expected, got interface{}, operation erx.Operation) *erx.Erx {
	return erx.WithArgs(
		erx.Kind("valueMisMatchError"),
		operation,
		fmt.Errorf("value mismatch: expected %v, got %v", expected, got),
	)
}
