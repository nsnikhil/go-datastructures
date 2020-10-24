package liberr

import "fmt"

var TypeMismatchError = func(expected, got string) error {
	return fmt.Errorf("expected %s got %s", expected, got)
}

var IndexOutOfBoundError = func(i int) error {
	return fmt.Errorf("index %d is out of bound", i)
}

var InvalidOperationError = func(reason string) error {
	return fmt.Errorf("invalid operation: %s", reason)
}

var NotFondError = func(e interface{}, s string) error {
	return fmt.Errorf("element %v not found in the %s", e, s)
}

var NotFondErrorInList = func(e interface{}) error {
	return NotFondError(e, "list")
}

var EmptyListError = InvalidOperationError("list is empty")
