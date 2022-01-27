package list

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/functions/predicate"
)

type factors struct {
	upperLoadFactor float64
	lowerLoadFactor float64

	scalingFactor int64
	capacity      int64
}

type ArrayList[T comparable] struct {
	*factors
	size int64
	data []T
}

func NewArrayList[T comparable](elements ...T) *ArrayList[T] {
	al := &ArrayList[T]{
		factors: &factors{
			upperLoadFactor: upperLoadFactor,
			lowerLoadFactor: lowerLoadFactor,
			scalingFactor:   scalingFactor,
			capacity:        initialCapacity,
		},
		size: nought,
		data: make([]T, initialCapacity),
	}

	al.addAllFrom(nought, elements...)

	return al
}

func (al *ArrayList[T]) Add(element T) {
	al.addAllFrom(al.size, element)
}

func (al *ArrayList[T]) AddAt(index int64, element T) error {
	return al.addAllFrom(index, element)
}

//TODO: SHOULD IT FAIL WHEN ARGS ARE EMPTY?
func (al *ArrayList[T]) AddAll(elements ...T) {
	al.addAllFrom(al.size, elements...)
}

//TODO: SHOULD WE JUST RESET THE SIZE OR ALSO RE-INITIALIZE THE DATA TO CLAIM UNUSED SPACE?
func (al *ArrayList[T]) Clear() {
	al.size = nought
	al.data = make([]T, initialCapacity)
}

func (al *ArrayList[T]) Clone() List[T] {
	if al.IsEmpty() {
		return NewArrayList[T]()
	}

	clonedList, _ := al.SubList(nought, al.size-1)

	return clonedList
}

func (al *ArrayList[T]) Contains(element T) bool {
	return newDoublyFinder[T]().search(al, element) != -1
}

func (al *ArrayList[T]) ContainsAll(elements ...T) bool {
	if al.IsEmpty() {
		return false
	}

	dataSet := make(map[T]bool)
	for i := int64(0); i < al.size; i++ {
		dataSet[al.data[i]] = true
	}

	for _, ele := range elements {
		if !al.Contains(ele) {
			return false
		}
	}

	return true
}

//TODO: WHAT TO RETURN AS DEFAULT VALUE?
func (al *ArrayList[T]) Get(index int64) (T, error) {
	if al.IsEmpty() {
		return *new(T), emptyListError("ArrayList.Get")
	}

	if ok := al.isValidIndex(index); !ok {
		return *new(T), invalidIndexError(index, "ArrayList.Get")
	}

	return al.data[index], nil
}

func (al *ArrayList[T]) IndexOf(element T) int64 {
	return newLinearFinder[T]().search(al, element)
}

func (al *ArrayList[T]) IsEmpty() bool {
	return al.size == nought
}

func (al *ArrayList[T]) LastIndexOf(element T) int64 {
	return newReverseFinder[T]().search(al, element)
}

func (al *ArrayList[T]) Remove(element T) error {
	if al.IsEmpty() {
		return emptyListError("ArrayList.Remove")
	}

	index := newLinearFinder[T]().search(al, element)
	if index == invalidIndex {
		return elementNotFoundError(element, "ArrayList.Remove")
	}

	al.removeAllFrom(index, index)

	return nil
}

func (al *ArrayList[T]) RemoveAt(index int64) (T, error) {
	if al.IsEmpty() {
		return *new(T), emptyListError("ArrayList.RemoveAt")
	}

	if ok := al.isValidIndex(index); !ok {
		return *new(T), invalidIndexError(index, "ArrayList.RemoveAt")
	}

	e, _ := al.Get(index)

	al.removeAllFrom(index, index)

	return e, nil
}

func (al *ArrayList[T]) RemoveAll(l ...T) error {
	return al.filterArrayList(false, l...)
}

func (al *ArrayList[T]) RemoveIf(predicate predicate.Predicate[T]) error {
	var l []T

	it := al.Iterator()

	for it.HasNext() {
		e, _ := it.Next()
		if predicate.Test(e) {
			l = append(l, e)
		}
	}

	return al.filterArrayList(false, l...)
}

func (al *ArrayList[T]) RemoveRange(from, to int64) error {
	return al.removeAllFrom(from, to)
}

func (al *ArrayList[T]) Replace(old, new T) error {
	if al.IsEmpty() {
		return emptyListError("ArrayList.Replace")
	}

	idx := al.IndexOf(old)
	if idx == invalidIndex {
		return elementNotFoundError(old, "ArrayList.Replace")
	}

	al.data[idx] = new

	return nil
}

func (al *ArrayList[T]) ReplaceAll(uo operator.UnaryOperator[T]) {
	sz := al.size

	for i := int64(0); i < sz; i++ {
		al.data[i] = uo.Apply(al.data[i])
	}
}

func (al *ArrayList[T]) RetainAll(l ...T) error {
	return al.filterArrayList(true, l...)
}

func (al *ArrayList[T]) Set(index int64, element T) (T, error) {
	if al.IsEmpty() {
		return *new(T), emptyListError("ArrayList.Set")
	}

	if ok := al.isValidIndex(index); !ok {
		return *new(T), invalidIndexError(index, "ArrayList.Set")
	}

	al.data[index] = element

	return al.data[index], nil
}

func (al *ArrayList[T]) Size() int64 {
	return al.size
}

func (al *ArrayList[T]) Sort(c comparator.Comparator[T]) {
	newQuickSorter[T]().sort(al, c)
}

func (al *ArrayList[T]) SubList(start int64, end int64) (List[T], error) {
	if end < start {
		return nil, invalidArgsError("end cannot be smaller than start", "ArrayList.SubList")
	}

	if ok := al.isValidIndex(start); !ok {
		return nil, invalidIndexError(start, "ArrayList.SubList")
	}

	if end < 0 || end >= al.size {
		return nil, invalidIndexError(end, "ArrayList.SubList")
	}

	tempList := NewArrayList[T]()

	for i := start; i <= end; i++ {
		tempList.Add(al.data[i])
	}

	return tempList, nil
}

type arrayListIterator[T comparable] struct {
	currentIndex int64
	al           *ArrayList[T]
}

func newArrayListIterator[T comparable](al *ArrayList[T]) iterator.Iterator[T] {
	return &arrayListIterator[T]{
		currentIndex: 0,
		al:           al,
	}
}

func (al *ArrayList[T]) Iterator() iterator.Iterator[T] {
	return newArrayListIterator[T](al)
}

func (ali *arrayListIterator[T]) HasNext() bool {
	return ali.currentIndex != ali.al.Size()
}

func (ali *arrayListIterator[T]) Next() (T, error) {
	if ali.currentIndex >= ali.al.Size() {
		return *new(T), emptyIteratorError("arrayListIterator.Next")
	}

	e, _ := ali.al.Get(ali.currentIndex)

	ali.currentIndex++

	return e, nil
}

type arrayListDescendingIterator[T comparable] struct {
	currentIndex int64
	al           *ArrayList[T]
}

func newArrayListDescendingIterator[T comparable](al *ArrayList[T]) iterator.Iterator[T] {
	return &arrayListDescendingIterator[T]{
		currentIndex: al.size - 1,
		al:           al,
	}
}

func (al *ArrayList[T]) DescendingIterator() iterator.Iterator[T] {
	return newArrayListDescendingIterator[T](al)
}

func (ali *arrayListDescendingIterator[T]) HasNext() bool {
	return ali.currentIndex >= 0
}

func (ali *arrayListDescendingIterator[T]) Next() (T, error) {
	if ali.currentIndex < 0 {
		return *new(T), emptyIteratorError("arrayListDescendingIterator.Next")
	}

	e, _ := ali.al.Get(ali.currentIndex)

	ali.currentIndex--

	return e, nil
}

func (al *ArrayList[T]) isValidIndex(i int64) bool {
	return i >= 0 && i < al.size
}

func (al *ArrayList[T]) filterArrayList(inverse bool, l ...T) error {
	if al.IsEmpty() {
		return emptyListError("ArrayList.filterArrayList")
	}

	idx := make(map[int64]bool)

	for _, e := range l {
		i := al.IndexOf(e)

		if i == invalidIndex {
			continue
		}

		idx[i] = true
	}

	sz := al.Size()
	tempData := make([]T, al.capacity)

	k := 0

	for i := int64(0); i < sz; i++ {
		if inverse {
			if idx[i] {
				el, _ := al.Get(i)
				tempData[k] = el
				k++
			}
		} else {
			if !idx[i] {
				el, _ := al.Get(i)
				tempData[k] = el
				k++
			}
		}
	}

	al.data = tempData

	if inverse {
		al.size = int64(len(l))
	} else {
		al.size -= int64(len(l))
	}

	al.checkAndDecreaseCapacity()

	return nil

}

func (al *ArrayList[T]) addAllFrom(index int64, elements ...T) error {
	//TODO: HERE THE AL.SIZE VALUE IS CONSIDERED A VALID INDEX WHERE AS IT IS NOT IN OTHER INVALID INDEX CASES
	if index < 0 || index > al.size {
		return invalidIndexError(index, "ArrayList.addAllFrom")
	}

	newElementsCount := int64(len(elements))

	al.checkAndIncreaseCapacity(newElementsCount)

	for i := al.size - 1; i >= index; i-- {
		al.data[i+newElementsCount] = al.data[i]
	}

	for _, e := range elements {
		al.data[index] = e
		index++
	}

	al.size += newElementsCount

	return nil
}

func (al *ArrayList[T]) removeAllFrom(startIndex, endIndex int64) error {
	if endIndex < startIndex {
		return invalidArgsError("end cannot be smaller than start", "ArrayList.removeAllFrom")
	}

	if ok := al.isValidIndex(startIndex); !ok {
		return invalidIndexError(startIndex, "ArrayList.removeAllFrom")
	}

	if ok := al.isValidIndex(endIndex); !ok {
		return invalidIndexError(endIndex, "ArrayList.removeAllFrom")
	}

	sz := (endIndex - startIndex) + 1

	i := startIndex
	j := endIndex + 1

	for j < al.size {
		al.data[i] = al.data[j]
		i++
		j++
	}

	al.size -= sz

	al.checkAndDecreaseCapacity()

	return nil
}

//TODO: REFACTOR, AND BREAKS THE SINGLE RESPONSIBILITY
func (al *ArrayList[T]) checkAndIncreaseCapacity(newElementsCount int64) {
	willExceedCapacity := func(newElementsCount int64, al *ArrayList[T]) bool {
		return al.size+newElementsCount >= int64(float64(al.capacity)*al.upperLoadFactor)
	}

	if !willExceedCapacity(newElementsCount, al) {
		return
	}

	for willExceedCapacity(newElementsCount, al) {
		al.capacity *= al.scalingFactor
	}

	al.data = resize(al.capacity, al.data)
}

//TODO: REFACTOR, AND BREAKS THE SINGLE RESPONSIBILITY
func (al *ArrayList[T]) checkAndDecreaseCapacity() {
	canDecreaseCapacity := func(al *ArrayList[T]) bool {
		return al.capacity != initialCapacity && al.size <= int64(float64(al.capacity)*al.lowerLoadFactor)
	}

	if !canDecreaseCapacity(al) {
		return
	}

	for canDecreaseCapacity(al) {
		al.capacity /= al.scalingFactor
	}

	al.data = resize(al.capacity, al.data)
}

func resize[T any](capacity int64, data []T) []T {
	temp := make([]T, capacity)

	sz := len(data)
	if len(temp) < sz {
		sz = len(temp)
	}

	for i := 0; i < sz; i++ {
		temp[i] = data[i]
	}

	return temp
}
