package liberror

import "fmt"

type InvalidOperation struct {
	message string
}

func NewInvalidOperation(message string) InvalidOperation {
	return InvalidOperation{message: message}
}

func (ivo InvalidOperation) Error() string {
	return fmt.Sprintf("invalid operation: %s", ivo.message)
}
