package list

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testObj struct {
	key   int
	value int
}

func newTestObj(key, value int) testObj {
	return testObj{
		key:   key,
		value: value,
	}
}

func (to testObj) product() int {
	return to.key * to.value
}

func (to testObj) compare(other testObj) int {
	return testObjComparator{}.Compare(to, other)
}

type testObjComparator struct{}

func (testObjComparator) Compare(one testObj, two testObj) int {
	return one.product() - two.product()
}

func TestQuickSortIntegerList(t *testing.T) {
	al := NewArrayList[int](5, 4, 3, 2, 1)

	newQuickSorter[int]().sort(al, comparator.NewIntegerComparator())

	expectedList := NewArrayList[int](1, 2, 3, 4, 5)

	assert.Equal(t, expectedList, al)
}

func TestQuickSortStringList(t *testing.T) {
	al := NewArrayList[string]("e", "d", "c", "b", "a")

	newQuickSorter[string]().sort(al, comparator.NewStringComparator())

	expectedList := NewArrayList[string]("a", "b", "c", "d", "e")

	assert.Equal(t, expectedList, al)
}

func TestQuickSortObjectList(t *testing.T) {
	al := NewArrayList[testObj](newTestObj(2, 3), newTestObj(4, 6), newTestObj(1, 4))

	newQuickSorter[testObj]().sort(al, testObjComparator{})

	expectedList := NewArrayList[testObj](newTestObj(1, 4), newTestObj(2, 3), newTestObj(4, 6))

	assert.Equal(t, expectedList, al)
}
