package erx_test

import (
	"errors"
	"erx"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestCreateErxWithArgsSuccess(t *testing.T) {
	testCases := map[string][]interface{}{
		"test create erx with only cause":                         {cause()},
		"test create erx with cause and kind":                     {cause(), kind()},
		"test create erx with cause and operation":                {cause(), operation()},
		"test create erx with cause and severity":                 {cause(), severity()},
		"test create erx with cause, kind and operation":          {cause(), kind(), operation()},
		"test create erx with cause, kind and severity":           {cause(), kind(), severity()},
		"test create erx with cause, operation and severity":      {cause(), operation(), severity()},
		"test create erx with cause, kind operation and severity": {cause(), kind(), operation(), severity()},
	}

	for name, args := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.NotNil(t, erx.WithArgs(args...))
		})
	}
}

func TestCreateErxWithArgsFailure(t *testing.T) {
	testCases := map[string][]interface{}{
		"test create failure when no args are provided":   {},
		"test create failure when cause is not provided":  {kind(), operation(), severity()},
		"test create failure when cause is redundant":     {cause(), cause()},
		"test create failure when operation is redundant": {operation(), operation()},
		"test create failure when kind is redundant":      {kind(), kind()},
		"test create failure when severity is redundant":  {severity(), severity()},
		"test create failure for invalid data":            {cause(), randInt(1, 8)},
	}

	for name, args := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Nil(t, erx.WithArgs(args...))
		})
	}
}

func TestGetKind(t *testing.T) {
	kind := kind()

	testCases := map[string]struct {
		args         []interface{}
		expectedKind erx.Kind
	}{
		"test get kind for flat erx": {
			args:         []interface{}{cause(), kind},
			expectedKind: kind,
		},
		"test get kind for nested erx": {
			args:         []interface{}{operation(), erx.WithArgs(cause(), kind)},
			expectedKind: kind,
		},
		"test return empty string when kind is missing": {
			args:         []interface{}{cause()},
			expectedKind: "",
		},
		"test return empty string when err is nil": {
			args:         nil,
			expectedKind: "",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedKind, erx.WithArgs(testCase.args...).Kind())
		})
	}
}

func TestGetOperations(t *testing.T) {
	op1 := operation()
	op2 := operation()
	op3 := operation()

	testCases := map[string]struct {
		args               []interface{}
		expectedOperations []erx.Operation
	}{
		"test get operation for flat erx": {
			args:               []interface{}{op1, cause()},
			expectedOperations: []erx.Operation{op1},
		},
		"test get operation for nested erx": {
			args:               []interface{}{op1, erx.WithArgs(op2, erx.WithArgs(op3, cause()))},
			expectedOperations: []erx.Operation{op1, op2, op3},
		},
		"test get operation for nested erx when op is missing in middle": {
			args:               []interface{}{op1, erx.WithArgs(erx.WithArgs(op3, cause()))},
			expectedOperations: []erx.Operation{op1, op3},
		},
		"test get operation return empty array when op is not present": {
			args:               []interface{}{cause()},
			expectedOperations: []erx.Operation{},
		},
		"test get operation return empty array erx is nil": {
			args:               []interface{}{},
			expectedOperations: []erx.Operation{},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedOperations, erx.WithArgs(testCase.args...).Operations())
		})
	}
}

func TestErxError(t *testing.T) {
	msg := randString(randInt(8, 14))

	testCases := map[string]struct {
		args           []interface{}
		expectedResult string
	}{
		"test error for flat erx": {
			args:           []interface{}{errors.New(msg)},
			expectedResult: msg,
		},
		"test error for nested erx": {
			args:           []interface{}{operation(), erx.WithArgs(errors.New(msg))},
			expectedResult: msg,
		},
		"test error return empty string for nil erx": {
			args:           []interface{}{},
			expectedResult: "",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, erx.WithArgs(testCase.args...).Error())
		})
	}
}

func TestErxToString(t *testing.T) {
	msg := randString(randInt(8, 14))
	cause := errors.New(msg)
	kind := kind()
	severity := severity()
	operation := operation()

	testCases := map[string]struct {
		args           []interface{}
		expectedResult string
	}{
		"test to string for flat erx with only cause": {
			args: []interface{}{cause},
			expectedResult: fmt.Sprintf(
				`{"cause":"%s"}`,
				msg,
			),
		},
		"test to string for flat erx with cause and operation": {
			args: []interface{}{cause, operation},
			expectedResult: fmt.Sprintf(
				`{"operation":"%s","cause":"%s"}`,
				operation, msg,
			),
		},
		"test to string for flat erx with cause kind and operation": {
			args: []interface{}{cause, kind, operation},
			expectedResult: fmt.Sprintf(
				`{"kind":"%s","operation":"%s","cause":"%s"}`,
				kind, operation, msg,
			),
		},
		"test to string for flat erx with cause, kind, operation and severity": {
			args: []interface{}{cause, kind, operation, severity},
			expectedResult: fmt.Sprintf(
				`{"kind":"%s","operation":"%s","severity":"%s","cause":"%s"}`,
				kind, operation, severity, msg,
			),
		},
		"test to string for nested erx with kind": {
			args: []interface{}{operation, severity, erx.WithArgs(kind, cause)},
			expectedResult: fmt.Sprintf(
				`{"operation":"%s","severity":"%s","cause":{"kind":"%s","cause":"%s"}}`,
				operation, severity, kind, msg,
			),
		},
		"test to string for nested erx with kind and severity": {
			args: []interface{}{operation, severity, erx.WithArgs(kind, severity, cause)},
			expectedResult: fmt.Sprintf(
				`{"operation":"%s","severity":"%s","cause":{"kind":"%s","severity":"%s","cause":"%s"}}`,
				operation, severity, kind, severity, msg,
			),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, erx.WithArgs(testCase.args...).String())
		})
	}
}

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func cause() error {
	return errors.New(
		randString(randInt(8, 12)),
	)
}

func kind() erx.Kind {
	kinds := []erx.Kind{
		erx.Kind(randString(8)),
		erx.Kind(randString(12)),
		erx.Kind(randString(10)),
	}

	return kinds[randInt(0, len(kinds)-1)]
}

func operation() erx.Operation {
	return erx.Operation(
		randString(randInt(8, 12)),
	)
}

func severity() erx.Severity {
	severities := []erx.Severity{
		erx.Severity(randString(8)),
		erx.Severity(randString(12)),
		erx.Severity(randString(10)),
	}

	return severities[randInt(0, len(severities)-1)]
}

func randString(n int) string {
	rand.Seed(time.Now().UnixNano())

	sz := len(letters)

	sb := strings.Builder{}
	sb.Grow(n)

	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rand.Intn(sz)])
	}

	return sb.String()
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
