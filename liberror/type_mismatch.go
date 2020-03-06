package liberror

import "fmt"

type TypeMismatchError struct {
	expected string
	got      string
}

func NewTypeMismatchError(expected string, got string) TypeMismatchError {
	return TypeMismatchError{
		expected: expected,
		got:      got,
	}
}

func (tme TypeMismatchError) Error() string {
	return fmt.Sprintf("expected %s got %s", tme.expected, tme.got)
}
