package gmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testType struct{}

func TestCreateNewPair(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   *Pair
		expectedResult *Pair
	}{
		{
			name:           "return pair of int and int32",
			actualResult:   NewPair(1, 'a'),
			expectedResult: &Pair{1, 'a'},
		},
		{
			name:           "return pair of int32 and struct",
			actualResult:   NewPair('t', testType{}),
			expectedResult: &Pair{'t', testType{}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult)
		})
	}
}

func TestPairGetKey(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   interface{}
		expectedResult interface{}
	}{
		{
			name:           "return key as 1",
			actualResult:   NewPair(1, 'a').GetKey(),
			expectedResult: 1,
		},
		{
			name:           "return key as t",
			actualResult:   NewPair('t', testType{}).GetKey(),
			expectedResult: 't',
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult)
		})
	}
}

func TestPairGetValue(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   interface{}
		expectedResult interface{}
	}{
		{
			name:           "return value as 1",
			actualResult:   NewPair(1, 'a').GetValue(),
			expectedResult: 'a',
		},
		{
			name:           "return key as testType",
			actualResult:   NewPair('t', testType{}).GetValue(),
			expectedResult: testType{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult)
		})
	}
}
