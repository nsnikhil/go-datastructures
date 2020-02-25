package comparator

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

type IntegerComparator struct{}

func NewIntegerComparator() IntegerComparator {
	return IntegerComparator{}
}

func (ic IntegerComparator) Compare(one interface{}, two interface{}) (int, error) {
	if reflect.TypeOf(one).Name() != reflect.TypeOf(0).Name() {
		return math.MinInt32, fmt.Errorf("invalid type : expected int got %s", reflect.TypeOf(one).Name())
	}

	if reflect.TypeOf(two).Name() != reflect.TypeOf(0).Name() {
		return math.MinInt32, fmt.Errorf("invalid type : expected int got %s", reflect.TypeOf(two).Name())
	}

	return one.(int) - two.(int), nil
}

type StringComparator struct{}

func NewStringComparator() StringComparator {
	return StringComparator{}
}

func (tc StringComparator) Compare(one interface{}, two interface{}) (int, error) {
	if reflect.TypeOf(one).Name() != reflect.TypeOf("").Name() {
		return math.MinInt32, fmt.Errorf("invalid type : expected string got %s", reflect.TypeOf(one).Name())
	}

	if reflect.TypeOf(two).Name() != reflect.TypeOf("").Name() {
		return math.MinInt32, fmt.Errorf("invalid type : expected string got %s", reflect.TypeOf(two).Name())
	}

	return strings.Compare(one.(string), two.(string)), nil
}
