package comparator

import (
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/nsnikhil/go-datastructures/utils"
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
		return math.MinInt32, liberror.NewTypeMismatchError("int", utils.GetTypeName(one))
	}

	if reflect.TypeOf(two).Name() != reflect.TypeOf(0).Name() {
		return math.MinInt32, liberror.NewTypeMismatchError("int", utils.GetTypeName(two))
	}

	return one.(int) - two.(int), nil
}

type StringComparator struct{}

func NewStringComparator() StringComparator {
	return StringComparator{}
}

func (tc StringComparator) Compare(one interface{}, two interface{}) (int, error) {
	if reflect.TypeOf(one).Name() != reflect.TypeOf("").Name() {
		return math.MinInt32, liberror.NewTypeMismatchError("string", utils.GetTypeName(one))
	}

	if reflect.TypeOf(two).Name() != reflect.TypeOf("").Name() {
		return math.MinInt32, liberror.NewTypeMismatchError("string", utils.GetTypeName(two))
	}

	return strings.Compare(one.(string), two.(string)), nil
}
