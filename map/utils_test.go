package gmap

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
	"reflect"
	"testing"
)

func TestToBytes(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() ([]byte, error)
		expectedResult []byte
		expectedError  error
	}{
		{
			name: "convert 1 to byte array",
			actualResult: func() ([]byte, error) {
				return toBytes(1)
			},
			expectedResult: []byte{0x31, 0xa},
		},
		{
			name: "convert a to byte array",
			actualResult: func() ([]byte, error) {
				return toBytes("a")
			},
			expectedResult: []byte{0x22, 0x61, 0x22, 0xa},
		},
		{
			name: "convert struct to byte array",
			actualResult: func() ([]byte, error) {
				type a struct{}
				return toBytes(a{})
			},
			expectedResult: []byte{0x7b, 0x7d, 0xa},
		},
		{
			name: "chan to bytes return error",
			actualResult: func() ([]byte, error) {
				return toBytes(make(chan int))
			},
			expectedResult: []byte(nil),
			expectedError:  &json.UnsupportedTypeError{Type: reflect.TypeOf(make(chan int))},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestHashCode(t *testing.T) {
	hs := sha3.New512()

	testCases := []struct {
		name           string
		actualResult   func() (uint32, error)
		expectedResult uint32
		expectedError  error
	}{
		{
			name: "return hashcode of 1",
			actualResult: func() (uint32, error) {
				return hashCode(&hs, 1)
			},
			expectedResult: 0x51e0aa1b,
		},
		{
			name: "return hashcode of 2",
			actualResult: func() (uint32, error) {
				return hashCode(&hs, 2)
			},
			expectedResult: 0xdbc2d0a3,
		},
		{
			name: "return hashcode of a",
			actualResult: func() (uint32, error) {
				return hashCode(&hs, "a")
			},
			expectedResult: 0xc91243a7,
		},
		{
			name: "return hashcode of b",
			actualResult: func() (uint32, error) {
				return hashCode(&hs, "b")
			},
			expectedResult: 0x2f12281e,
		},
		{
			name: "return hashcode of struct a",
			actualResult: func() (uint32, error) {
				type a struct{}
				return hashCode(&hs, a{})
			},
			expectedResult: 0x22b7192f,
		},
		{
			name: "return hashcode of struct b",
			actualResult: func() (uint32, error) {
				type b struct{}
				return hashCode(&hs, b{})
			},
			expectedResult: 0x22b7192f,
		},
		{
			name: "return hashcode of chan return error",
			actualResult: func() (uint32, error) {
				return hashCode(&hs, make(chan int))
			},
			expectedResult: 0x7fffffff,
			expectedError:  &json.UnsupportedTypeError{Type: reflect.TypeOf(make(chan int))},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestIndexOf(t *testing.T) {
	hs := sha3.New256()
	var cp float64 = 32

	testCases := []struct {
		name           string
		actualResult   func() ([]int, error)
		expectedResult []int
		expectedError  error
	}{
		{
			name: "return hash for first 100 numbers",
			actualResult: func() ([]int, error) {
				res := make([]int, 100)

				for i := 0; i < 100; i++ {
					idx, err := indexOf(&hs, i, cp)
					if err != nil {
						return nil, err
					}

					res[i] = idx
				}

				return res, nil
			},
			expectedResult: []int{6, 11, 31, 25, 30, 11, 17, 26, 0, 4, 18, 7, 18, 12, 29, 25, 15, 28, 18, 8, 4, 16, 10, 18, 17, 0, 25, 11, 30, 12, 20, 26, 30, 7, 11, 20, 15, 21, 0, 20, 21, 14, 15, 26, 21, 31, 1, 16, 4, 24, 12, 20, 3, 24, 8, 30, 16, 3, 18, 16, 10, 30, 3, 12, 19, 10, 3, 14, 28, 4, 8, 26, 21, 18, 24, 5, 29, 4, 1, 18, 9, 1, 11, 1, 30, 11, 10, 31, 1, 12, 14, 6, 24, 30, 26, 2, 30, 11, 27, 19},
		},
		{
			name: "return hash for other 100 numbers",
			actualResult: func() ([]int, error) {
				res := make([]int, 100)

				for i := 0; i < 100; i += 2 {
					idx, err := indexOf(&hs, (200+i)*(1<<16), cp)
					if err != nil {
						return nil, err
					}

					res[i] = idx
				}

				return res, nil
			},
			expectedResult: []int{29, 0, 7, 0, 29, 0, 22, 0, 14, 0, 17, 0, 17, 0, 16, 0, 30, 0, 1, 0, 10, 0, 0, 0, 4, 0, 4, 0, 25, 0, 20, 0, 23, 0, 21, 0, 19, 0, 1, 0, 25, 0, 8, 0, 0, 0, 1, 0, 21, 0, 21, 0, 12, 0, 16, 0, 21, 0, 8, 0, 19, 0, 19, 0, 19, 0, 20, 0, 28, 0, 11, 0, 24, 0, 4, 0, 8, 0, 23, 0, 9, 0, 4, 0, 12, 0, 24, 0, 25, 0, 16, 0, 20, 0, 16, 0, 6, 0, 2, 0},
		},
		{
			name: "return hash for all ascii values",
			actualResult: func() ([]int, error) {
				res := make([]int, 128)

				for i := 0; i < 128; i++ {
					idx, err := indexOf(&hs, int32(i), cp)
					if err != nil {
						return nil, err
					}

					res[i] = idx
				}

				return res, nil
			},
			expectedResult: []int{6, 11, 31, 25, 30, 11, 17, 26, 0, 4, 18, 7, 18, 12, 29, 25, 15, 28, 18, 8, 4, 16, 10, 18, 17, 0, 25, 11, 30, 12, 20, 26, 30, 7, 11, 20, 15, 21, 0, 20, 21, 14, 15, 26, 21, 31, 1, 16, 4, 24, 12, 20, 3, 24, 8, 30, 16, 3, 18, 16, 10, 30, 3, 12, 19, 10, 3, 14, 28, 4, 8, 26, 21, 18, 24, 5, 29, 4, 1, 18, 9, 1, 11, 1, 30, 11, 10, 31, 1, 12, 14, 6, 24, 30, 26, 2, 30, 11, 27, 19, 31, 16, 3, 18, 16, 26, 29, 11, 31, 30, 5, 6, 19, 1, 23, 9, 21, 12, 14, 15, 26, 10, 4, 7, 5, 28, 31, 7},
		},
		{
			name: "return hash indexOf of struct",
			actualResult: func() ([]int, error) {
				type a struct{ I int }

				res := make([]int, 100)

				for i := 0; i < 100; i++ {
					v := i

					if v%2 != 0 {
						v = 0
					}

					idx, err := indexOf(&hs, a{I: v}, cp)
					if err != nil {
						return nil, err
					}

					res[i] = idx
				}

				return res, nil
			},
			expectedResult: []int{9, 9, 16, 9, 2, 9, 7, 9, 18, 9, 16, 9, 19, 9, 27, 9, 0, 9, 14, 9, 18, 9, 15, 9, 21, 9, 26, 9, 31, 9, 28, 9, 7, 9, 3, 9, 10, 9, 19, 9, 3, 9, 0, 9, 22, 9, 4, 9, 26, 9, 3, 9, 16, 9, 18, 9, 30, 9, 16, 9, 7, 9, 20, 9, 14, 9, 16, 9, 20, 9, 8, 9, 0, 9, 26, 9, 29, 9, 1, 9, 20, 9, 4, 9, 4, 9, 10, 9, 17, 9, 31, 9, 4, 9, 12, 9, 2, 9, 29, 9},
		},
		{
			name: "return error when hashing channel",
			actualResult: func() ([]int, error) {
				idx, err := indexOf(&hs, make(chan int), cp)
				return []int{idx}, err
			},
			expectedResult: []int{-1},
			expectedError:  &json.UnsupportedTypeError{Type: reflect.TypeOf(make(chan int))},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
