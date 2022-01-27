package comparator

import (
	"strings"
)

type IntegerComparator struct{}

func NewIntegerComparator() Comparator[int] {
	return IntegerComparator{}
}

func (ic IntegerComparator) Compare(one int, two int) int {
	return one - two
}

type StringComparator struct{}

func NewStringComparator() Comparator[string] {
	return StringComparator{}
}

func (tc StringComparator) Compare(one string, two string) int {
	return strings.Compare(one, two)
}
