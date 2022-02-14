package gmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testType struct{}

func TestCreateNewPair(t *testing.T) {
	type object struct{ data int }

	intRunePair := NewPair[int, rune](1, 'a')

	assert.Equal(t, &Pair[int, rune]{1, 'a'}, intRunePair)

	stringObject := NewPair[string, object]("a", object{data: 1})

	assert.Equal(t, &Pair[string, object]{"a", object{data: 1}}, stringObject)

}

func TestPairGetKey(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   interface{}
		expectedResult interface{}
	}{
		{
			name:           "return key as 1",
			actualResult:   NewPair(1, 'a').First(),
			expectedResult: 1,
		},
		{
			name:           "return key as t",
			actualResult:   NewPair('t', testType{}).First(),
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
			actualResult:   NewPair(1, 'a').Second(),
			expectedResult: 'a',
		},
		{
			name:           "return key as testType",
			actualResult:   NewPair('t', testType{}).Second(),
			expectedResult: testType{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult)
		})
	}
}
