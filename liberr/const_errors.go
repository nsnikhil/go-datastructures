package liberr

import (
	"fmt"
)

var TypeMismatchError = func(expected, got string) error {
	return fmt.Errorf("expected %s got %s", expected, got)
}
