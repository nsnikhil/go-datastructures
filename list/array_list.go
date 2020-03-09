package list

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/functions/predicate"
	"github.com/nsnikhil/go-datastructures/liberror"
)

type ArrayList struct {
	typeURL string
	data    []interface{}
}

func NewArrayList(data ...interface{}) (*ArrayList, error) {
	if len(data) == 0 {
		return &ArrayList{
			typeURL: na,
		}, nil
	}

	typeURL := getTypeName(data[0])

	for i := 1; i < len(data); i++ {
		if getTypeName(data[i]) != typeURL {
			return nil, liberror.NewTypeMismatchError(typeURL, getTypeName(data[i]))
		}
	}

	return &ArrayList{
		typeURL: typeURL,
		data:    data,
	}, nil
}

func (al *ArrayList) Add(e interface{}) error {

	if al.Size() == 0 && al.typeURL == na {
		al.data = append(al.data, e)
		al.typeURL = getTypeName(e)
		return nil
	}

	if err := al.isValidType(e); err != nil {
		return err
	}

	al.data = append(al.data, e)
	return nil
}

func (al *ArrayList) AddAt(i int, e interface{}) error {
	if al.IsEmpty() && al.typeURL == na {
		al.data = append(al.data, e)
		al.typeURL = getTypeName(e)
		return nil
	}

	if err := al.isValidIndex(i); err != nil {
		return err
	}

	if err := al.isValidType(e); err != nil {
		return err
	}

	al.data = append(al.data[:i], append([]interface{}{e}, al.data[i:]...)...)

	return nil
}

func (al *ArrayList) AddAll(l ...interface{}) error {
	if len(l) == 0 {
		return nil
	}

	for i := 0; i < len(l)-1; i++ {
		if getTypeName(l[i]) != getTypeName(l[i+1]) {
			return fmt.Errorf("type mismatch : all elements must be of same type")
		}
	}

	idx := 0

	if al.typeURL == na {
		al.data = append(al.data, l[idx])
		al.typeURL = getTypeName(l[idx])
		idx++
	} else {

		if err := al.isValidType(l[idx]); err != nil {
			return err
		}

		if al.IsEmpty() {
			al.data = append(al.data, l[idx])
			idx++
		}

	}

	al.data = append(al.data, l[idx:]...)

	return nil
}

func (al *ArrayList) Clear() {
	al.data = nil
}

func (al *ArrayList) Clone() (List, error) {
	if al.IsEmpty() {
		return NewArrayList()
	}
	return al.SubList(0, al.Size())
}

func (al *ArrayList) Contains(e interface{}) (bool, error) {

	_, err := newFinder(concurrent).search(al, e)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (al *ArrayList) ContainsAll(l ...interface{}) (bool, error) {
	for _, d := range l {
		ok, err := al.Contains(d)
		if !ok || err != nil {
			return false, err
		}
	}
	return true, nil
}

func (al *ArrayList) Get(i int) interface{} {
	if err := al.isValidIndex(i); al.IsEmpty() || err != nil {
		return nil
	}
	return al.data[i]
}

func (al *ArrayList) IndexOf(e interface{}) (int, error) {
	if al.IsEmpty() {
		return invalidIndex, fmt.Errorf("list is empty")
	}

	if err := al.isValidType(e); err != nil {
		return invalidIndex, err
	}

	i, _ := newFinder(concurrent).search(al, e)
	if i == invalidIndex {
		return invalidIndex, fmt.Errorf("element %v not found in the list", e)
	}

	return i, nil
}

func (al *ArrayList) IsEmpty() bool {
	return al.Size() == 0
}

func (al *ArrayList) Iterator() iterator.Iterator {
	return newArrayListIterator(al)
}

func (al *ArrayList) LastIndexOf(e interface{}) (int, error) {
	if al.IsEmpty() {
		return invalidIndex, fmt.Errorf("list is empty")
	}

	if err := al.isValidType(e); err != nil {
		return invalidIndex, err
	}

	i := al.Size() - 1
	for i >= 0 {
		if al.Get(i) == e {
			return i, nil
		}
		i--
	}

	return invalidIndex, fmt.Errorf("element %v not found in the list", e)
}

func (al *ArrayList) Remove(e interface{}) (bool, error) {
	if al.IsEmpty() {
		return false, fmt.Errorf("list is empty")
	}

	if err := al.isValidType(e); err != nil {
		return false, err
	}

	i, err := al.IndexOf(e)
	if err != nil || i == invalidIndex {
		return false, err
	}

	al.data = append(al.data[0:i], al.data[i+1:]...)

	return true, nil
}

func (al *ArrayList) RemoveAt(i int) (interface{}, error) {
	if al.IsEmpty() {
		return nil, fmt.Errorf("list is empty")
	}

	if err := al.isValidIndex(i); err != nil {
		return nil, err
	}

	e := al.Get(i)
	removeAt(al, i)

	return e, nil
}

func (al *ArrayList) RemoveAll(l ...interface{}) (bool, error) {
	return filterArrayList(al, false, l...)
}

func (al *ArrayList) RemoveIf(predicate predicate.Predicate) (bool, error) {
	var l []interface{}

	it := al.Iterator()

	for it.HasNext() {
		e := it.Next()
		if predicate.Test(e) {
			l = append(l, e)
		}
	}

	return filterArrayList(al, false, l...)
}

func (al *ArrayList) RemoveRange(from, to int) (bool, error) {
	if to < from {
		return false, fmt.Errorf("to cannot be smaller than from")
	}

	if err := al.isValidIndex(from); err != nil {
		return false, err
	}

	//TODO USE IS VALID INDEX METHOD HERE
	if to < 0 || to > al.Size() {
		return false, liberror.NewIndexOutOfBoundError(to)
	}

	al.data = append(al.data[:from], al.data[to:]...)

	return true, nil
}

func (al *ArrayList) ReplaceAll(uo operator.UnaryOperator) error {
	sz := al.Size()
	for i := 0; i < sz; i++ {

		e := uo.Apply(al.Get(i))

		if err := al.isValidType(e); err != nil {
			return err
		}

		if _, err := al.Set(i, e); err != nil {
			return err
		}
	}

	return nil
}

func (al *ArrayList) RetainAll(l ...interface{}) (bool, error) {
	return filterArrayList(al, true, l...)
}

func (al *ArrayList) Set(i int, e interface{}) (interface{}, error) {
	if al.IsEmpty() {
		return nil, fmt.Errorf("list is empty")
	}

	if err := al.isValidIndex(i); err != nil {
		return nil, err
	}

	if err := al.isValidType(e); err != nil {
		return nil, err
	}

	al.data[i] = e
	return al.data[i], nil
}

func (al *ArrayList) Size() int {
	return len(al.data)
}

func (al *ArrayList) Sort(c comparator.Comparator) {
	newSorter().sort(al, c)
}

func (al *ArrayList) SubList(s int, e int) (List, error) {
	if e < s {
		return nil, fmt.Errorf("end cannot be smaller than start")
	}

	if err := al.isValidIndex(s); err != nil {
		return nil, err
	}

	//TODO USE IS VALID INDEX METHOD HERE
	if e < 0 || e > al.Size() {
		return nil, liberror.NewIndexOutOfBoundError(e)
	}

	tempData := make([]interface{}, 0)
	for i := s; i < e; i++ {
		tempData = append(tempData, al.Get(i))
	}

	return NewArrayList(tempData...)
}

type arrayListIterator struct {
	currentIndex int
	al           *ArrayList
}

func newArrayListIterator(al *ArrayList) *arrayListIterator {
	return &arrayListIterator{
		currentIndex: 0,
		al:           al,
	}
}

func (ali *arrayListIterator) HasNext() bool {
	return ali.currentIndex != ali.al.Size()
}

func (ali *arrayListIterator) Next() interface{} {
	if ali.currentIndex >= ali.al.Size() {
		return nil
	}

	e := ali.al.Get(ali.currentIndex)
	ali.currentIndex++

	return e
}

func (al *ArrayList) isValidIndex(i int) error {
	if i < 0 || i >= al.Size() {
		return liberror.NewIndexOutOfBoundError(i)
	}
	return nil
}

func (al *ArrayList) isValidType(e interface{}) error {
	if getTypeName(e) != al.typeURL {
		return liberror.NewTypeMismatchError(al.typeURL, getTypeName(e))
	}
	return nil
}

func filterArrayList(al *ArrayList, inverse bool, l ...interface{}) (bool, error) {
	if al.IsEmpty() {
		return false, fmt.Errorf("list is empty")
	}

	for _, e := range l {
		if err := al.isValidType(e); err != nil {
			return false, err
		}
	}

	idx := make(map[int]bool, 0)
	for _, e := range l {
		i, _ := al.IndexOf(e)

		if i == invalidIndex {
			continue
		}

		idx[i] = true
	}

	sz := al.Size()
	tempData := make([]interface{}, 0)
	for i := 0; i < sz; i++ {
		if inverse {
			if idx[i] {
				tempData = append(tempData, al.Get(i))
			}
		} else {
			if !idx[i] {
				tempData = append(tempData, al.Get(i))
			}
		}
	}

	al.data = tempData

	return true, nil

}

func removeAt(al *ArrayList, i int) {
	al.data = append(al.data[0:i], al.data[i+1:]...)
}
