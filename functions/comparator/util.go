package comparator

import (
	"github.com/nsnikhil/go-datastructures/liberr"
	"github.com/nsnikhil/go-datastructures/utils"
	"math"
	"strings"
)

type IntegerComparator struct{}

func NewIntegerComparator() IntegerComparator {
	return IntegerComparator{}
}

func (ic IntegerComparator) Compare(one interface{}, two interface{}) (int, error) {
	if utils.GetTypeName(one) != utils.GetTypeName(0) {
		return math.MinInt32, liberr.TypeMismatchError("int", utils.GetTypeName(one))
	}

	if utils.GetTypeName(two) != utils.GetTypeName(0) {
		return math.MinInt32, liberr.TypeMismatchError("int", utils.GetTypeName(two))
	}

	return one.(int) - two.(int), nil
}

type StringComparator struct{}

func NewStringComparator() StringComparator {
	return StringComparator{}
}

func (tc StringComparator) Compare(one interface{}, two interface{}) (int, error) {
	if utils.GetTypeName(one) != utils.GetTypeName("") {
		return math.MinInt32, liberr.TypeMismatchError("string", utils.GetTypeName(one))
	}

	if utils.GetTypeName(two) != utils.GetTypeName("") {
		return math.MinInt32, liberr.TypeMismatchError("string", utils.GetTypeName(two))
	}

	return strings.Compare(one.(string), two.(string)), nil
}
