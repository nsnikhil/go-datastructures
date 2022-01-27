package internal

import (
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
