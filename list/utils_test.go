package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTypeName(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() string
		expectedResult string
	}{
		{
			name: "return type name as int",
			actualResult: func() string {
				return getTypeName(1)
			},
			expectedResult: "int",
		},
		{
			name: "return type name as string",
			actualResult: func() string {
				return getTypeName("abc")
			},
			expectedResult: "string",
		},
		{
			name: "return type name as someObj",
			actualResult: func() string {
				type someObj struct{}
				return getTypeName(someObj{})
			},
			expectedResult: "someObj",
		},
		{
			name: "return type name as na for interface",
			actualResult: func() string {
				var e interface{}
				return getTypeName(e)
			},
			expectedResult: "na",
		},
		{
			name: "return type name as na for nil",
			actualResult: func() string {
				return getTypeName(nil)
			},
			expectedResult: "na",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}
