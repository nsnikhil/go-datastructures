package erx

import (
	"fmt"
	"strings"
)

//Erx type to create new error with kind, operation, severity and stack erros
type Erx struct {
	internalErx
}

type internalErx struct {
	Kind      Kind      `json:"kind,omitempty"`
	Operation Operation `json:"operation,omitempty"`
	Severity  Severity  `json:"severity,omitempty"`
	Cause     error     `json:"cause"`
}

func (e *Erx) Error() string {
	if e == nil || e.internalErx.Cause == nil {
		return ""
	}

	return e.internalErx.Cause.Error()
}

//Kind returns the first error kind in stack else empty string
func (e *Erx) Kind() Kind {
	if e == nil {
		return ""
	}

	if len(e.internalErx.Kind) != 0 {
		return e.internalErx.Kind
	}

	t, ok := e.internalErx.Cause.(*Erx)
	if !ok {
		return ""
	}

	return t.Kind()
}

//Operations returns all the operations in the stack
func (e *Erx) Operations() []Operation {
	ops := make([]Operation, 0)
	erxOperations(e, &ops)
	return ops
}

func erxOperations(e *Erx, ops *[]Operation) {
	if e == nil {
		return
	}

	if len(e.internalErx.Operation) != 0 {
		*ops = append(*ops, e.internalErx.Operation)
	}

	t, ok := e.internalErx.Cause.(*Erx)
	if !ok {
		return
	}

	erxOperations(t, ops)
}

func (e *Erx) String() string {
	return marshal(e)
}

func marshal(e *Erx) string {
	builder := new(strings.Builder)
	builder.WriteRune('{')

	wk, wo, ws := false, false, false

	if len(e.internalErx.Kind) != 0 {
		builder.WriteString(fmt.Sprintf(`"kind":"%s"`, e.internalErx.Kind))
		wk = true
	}

	if len(e.internalErx.Operation) != 0 {
		if wk {
			builder.WriteRune(',')
			wk = false
		}

		builder.WriteString(fmt.Sprintf(`"operation":"%s"`, e.internalErx.Operation))
		wo = true
	}

	if len(e.internalErx.Severity) != 0 {
		if wk || wo {
			builder.WriteRune(',')
			wk = false
			wo = false
		}

		builder.WriteString(fmt.Sprintf(`"severity":"%s"`, e.internalErx.Severity))
		ws = true
	}

	if e.Cause != nil {
		if wk || wo || ws {
			builder.WriteRune(',')
			wk = false
			wo = false
			ws = false
		}

		t, ok := e.Cause.(*Erx)
		if ok {
			builder.WriteString(fmt.Sprintf(`"cause":%s`, marshal(t)))
		} else {
			builder.WriteString(fmt.Sprintf(`"cause":"%s"`, e.internalErx.Cause.Error()))
		}

	}

	builder.WriteRune('}')
	return builder.String()
}

//WithArgs create new kind, severity, operation and cause
func WithArgs(args ...interface{}) *Erx {
	if len(args) == 0 {
		return nil
	}

	e := &Erx{}

	for _, arg := range args {
		switch t := arg.(type) {
		case Operation:
			if len(e.internalErx.Operation) != 0 {
				return nil
			}

			e.internalErx.Operation = t
		case Kind:
			if len(e.internalErx.Kind) != 0 {
				return nil
			}

			e.internalErx.Kind = t
		case Severity:
			if len(e.internalErx.Severity) != 0 {
				return nil
			}

			e.internalErx.Severity = t
		case error:
			if e.internalErx.Cause != nil {
				return nil
			}

			e.internalErx.Cause = t
		default:
			return nil
		}
	}

	if e.internalErx.Cause == nil {
		return nil
	}

	return e
}
