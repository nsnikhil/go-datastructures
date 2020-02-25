package comparator

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"testing"
)

type obj struct {
	id     int
	values []int
}

func newObj(id int, values ...int) obj {
	return obj{
		id:     id,
		values: values,
	}
}

type testObjComparator struct{}

func newTestObjComparator() *testObjComparator {
	return &testObjComparator{}
}

func (tc *testObjComparator) Compare(one interface{}, two interface{}) (int, error) {
	if reflect.TypeOf(one).Name() != reflect.TypeOf(obj{}).Name() {
		return math.MinInt32, fmt.Errorf("invalid type : expected obj got %s", reflect.TypeOf(one).Name())
	}

	if reflect.TypeOf(two).Name() != reflect.TypeOf(obj{}).Name() {
		return math.MinInt32, fmt.Errorf("invalid type : expected obj got %s", reflect.TypeOf(two).Name())
	}

	a := one.(obj)
	b := two.(obj)

	const (
		greater = 1
		smaller = -1
		equals  = 0
	)

	idDiff := a.id - b.id
	lenDiff := len(a.values) - len(b.values)

	if idDiff == 0 && lenDiff == 0 {
		return equals, nil
	}

	if idDiff > 0 && lenDiff > 0 {
		return greater, nil
	}

	return smaller, nil

}

func TestCompare(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int, error)
		expectedResult func(res int) bool
		expectedError  error
	}{
		{
			name: "test integer comparator for greater",
			actualResult: func() (int, error) {
				ic := NewIntegerComparator()
				return ic.Compare(2, 1)
			},
			expectedResult: func(res int) bool {
				return res > 0
			},
		},
		{
			name: "test integer comparator for smaller",
			actualResult: func() (int, error) {
				ic := NewIntegerComparator()
				return ic.Compare(1, 2)
			},
			expectedResult: func(res int) bool {
				return res < 0
			},
		},
		{
			name: "test integer comparator for equals",
			actualResult: func() (int, error) {
				ic := NewIntegerComparator()
				return ic.Compare(2, 2)
			},
			expectedResult: func(res int) bool {
				return res == 0
			},
		},
		{
			name: "test integer comparator return error for type mismatch",
			actualResult: func() (int, error) {
				ic := NewIntegerComparator()
				return ic.Compare(2, "invalid")
			},
			expectedResult: func(res int) bool {
				return res == math.MinInt32
			},
			expectedError: errors.New("invalid type : expected int got string"),
		},
		{
			name: "test string comparator for greater",
			actualResult: func() (int, error) {
				ic := NewStringComparator()
				return ic.Compare("ab", "a")
			},
			expectedResult: func(res int) bool {
				return res > 0
			},
		},
		{
			name: "test string comparator for smaller",
			actualResult: func() (int, error) {
				ic := NewStringComparator()
				return ic.Compare("a", "ab")
			},
			expectedResult: func(res int) bool {
				return res < 0
			},
		},
		{
			name: "test string comparator for equals",
			actualResult: func() (int, error) {
				ic := NewStringComparator()
				return ic.Compare("a", "a")
			},
			expectedResult: func(res int) bool {
				return res == 0
			},
		},
		{
			name: "test string comparator return error for type mismatch",
			actualResult: func() (int, error) {
				ic := NewStringComparator()
				return ic.Compare("ab", 1)
			},
			expectedResult: func(res int) bool {
				return res == math.MinInt32
			},
			expectedError: errors.New("invalid type : expected string got int"),
		},
		{
			name: "test obj comparator for greater",
			actualResult: func() (int, error) {
				ic := newTestObjComparator()
				a := newObj(2, 1, 2)
				b := newObj(1, 1)
				return ic.Compare(a, b)
			},
			expectedResult: func(res int) bool {
				return res > 0
			},
		},
		{
			name: "test obj comparator for smaller when id is smaller",
			actualResult: func() (int, error) {
				ic := newTestObjComparator()
				a := newObj(2, 1, 2)
				b := newObj(4, 1)
				return ic.Compare(a, b)
			},
			expectedResult: func(res int) bool {
				return res < 0
			},
		},
		{
			name: "test obj comparator for smaller when there are less values",
			actualResult: func() (int, error) {
				ic := newTestObjComparator()
				a := newObj(2, 1)
				b := newObj(1, 1, 2)
				return ic.Compare(a, b)
			},
			expectedResult: func(res int) bool {
				return res < 0
			},
		},
		{
			name: "test obj comparator for equals",
			actualResult: func() (int, error) {
				ic := newTestObjComparator()
				a := newObj(1, 1, 2)
				b := newObj(1, 1, 2)
				return ic.Compare(a, b)
			},
			expectedResult: func(res int) bool {
				return res == 0
			},
		},
		{
			name: "test obj comparator return error for type mismatch",
			actualResult: func() (int, error) {
				ic := newTestObjComparator()
				a := newObj(1, 1, 2)
				return ic.Compare(a, "ab")
			},
			expectedResult: func(res int) bool {
				return res == math.MinInt32
			},
			expectedError: errors.New("invalid type : expected obj got string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.True(t, testCase.expectedResult(res))
		})
	}
}
