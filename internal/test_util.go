package internal

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func AssertErrorEquals(t *testing.T, expected, actual error) {
	if expected == nil && actual == nil {
		return
	}

	assert.Equal(t, expected.Error(), actual.Error())
}

//TODO: REFACTOR
func AssertSliceEquals[T comparable](t *testing.T, a, b []T) {
	if len(a) != len(b) {
		assert.Fail(t, fmt.Sprintf("len of a : %b, is not equals to len of b: %v", len(a), len(b)))
	}

	keys := func(sl []T) map[T]bool {
		res := make(map[T]bool)

		for _, e := range sl {
			res[e] = true
		}

		return res
	}

	ak := keys(a)
	bk := keys(b)

	for v := range ak {
		if !bk[v] {
			assert.Fail(t, fmt.Sprintf("slice %v, and %v are not equals", a, b))
		}
	}
}

func AreMapsSame[K comparable, V any](a, b map[K]V, vc comparator.Comparator[V]) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		bv, ok := b[k]
		if !ok {
			return false
		}

		if vc.Compare(v, bv) != 0 {
			return false
		}
	}

	return true
}

type SliceGenerator struct {
	Min           int64
	Max           int64
	AbsoluteValue int64
	Size          int64
	Reverse       bool
	Random        bool
	Absolute      bool
}

func (sg SliceGenerator) Generate() []int64 {
	return generateSlice(sg.Min, sg.Max, sg.AbsoluteValue, sg.Size, sg.Random, sg.Reverse, sg.Absolute)
}

func generateSlice(min, max, absoluteValue, size int64, random, reverse, absolute bool) []int64 {
	getRandomNumber := func(min, max int64) int64 {
		rand.Seed(time.Now().UnixNano())
		return rand.Int63n(max-min) + min
	}

	res := make([]int64, size)
	for i := int64(0); i < size; i++ {
		idx := i

		if reverse {
			idx = size - i - 1
		}

		if absolute {
			res[idx] = absoluteValue
		} else if random {
			res[idx] = getRandomNumber(min, max)
		} else {
			res[idx] = i
		}

	}

	return res
}
