package liberror

import "fmt"

type IndexOutOfBoundError struct {
	outOfBoundIndex int
}

func NewIndexOutOfBoundError(outOfBoundIndex int) IndexOutOfBoundError {
	return IndexOutOfBoundError{outOfBoundIndex: outOfBoundIndex}
}

func (iob IndexOutOfBoundError) Error() string {
	return fmt.Sprintf("index %d is out of bound", iob.outOfBoundIndex)
}
