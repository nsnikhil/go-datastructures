package list

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/functions/predicate"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/nsnikhil/go-datastructures/utils"
)

type factors struct {
	upperLoadFactor float64
	lowerLoadFactor float64

	scalingFactor int
	capacity      int
}

type ArrayList struct {
	typeURL string
	*factors
	size int
	data []interface{}
}

func NewArrayList(data ...interface{}) (*ArrayList, error) {
	al := &ArrayList{
		factors: &factors{upperLoadFactor: upperLoadFactor, lowerLoadFactor: lowerLoadFactor, scalingFactor: scalingFactor, capacity: initialCapacity},
		size:    nought,
		typeURL: utils.NA,
		data:    make([]interface{}, initialCapacity),
	}

	if len(data) == 0 {
		return al, nil
	}

	if err := addAll(al, data...); err != nil {
		return nil, err
	}

	return al, nil
}

func (al *ArrayList) Add(e interface{}) error {
	return addAll(al, e)
}

func (al *ArrayList) AddAt(i int, e interface{}) error {
	if al.IsEmpty() && al.typeURL == utils.NA {
		al.typeURL = utils.GetTypeName(e)
		return addAll(al, e)
	}

	if err := al.isValidIndex(i); err != nil {
		return err
	}

	if err := al.isValidType(e); err != nil {
		return err
	}

	checkIncreaseCapacity(al)

	for j := al.Size(); j > i; j-- {
		al.data[j] = al.data[j-1]
	}

	al.data[i] = e

	al.size++

	return nil
}

func (al *ArrayList) AddAll(l ...interface{}) error {
	return addAll(al, l...)
}

func addAll(al *ArrayList, l ...interface{}) error {
	if len(l) == 0 {
		return nil
	}

	for i := 0; i < len(l)-1; i++ {
		if utils.GetTypeName(l[i]) != utils.GetTypeName(l[i+1]) {
			return fmt.Errorf("type mismatch : all elements must be of same type")
		}
	}

	eleType := utils.GetTypeName(l[0])

	if al.typeURL == utils.NA {
		al.typeURL = eleType
	} else if al.typeURL != eleType {
		return liberror.NewTypeMismatchError(al.typeURL, eleType)
	}

	for _, e := range l {
		if err := add(al, e); err != nil {
			return nil
		}
	}

	return nil
}

func add(al *ArrayList, e interface{}) error {
	checkIncreaseCapacity(al)

	if al.typeURL != utils.GetTypeName(e) {
		return liberror.NewTypeMismatchError(al.typeURL, utils.GetTypeName(e))
	}

	al.data[al.size] = e
	al.size++

	return nil
}

func (al *ArrayList) Clear() {
	for i := 0; i < al.Size(); i++ {
		al.data[i] = nil
	}
	al.size = nought
}

func (al *ArrayList) Clone() (List, error) {
	if al.IsEmpty() {
		return NewArrayList()
	}
	return al.SubList(nought, al.Size())
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
		return utils.InvalidIndex, fmt.Errorf("list is empty")
	}

	if err := al.isValidType(e); err != nil {
		return utils.InvalidIndex, err
	}

	i, _ := newFinder(concurrent).search(al, e)
	if i == utils.InvalidIndex {
		return utils.InvalidIndex, fmt.Errorf("element %v not found in the list", e)
	}

	return i, nil
}

func (al *ArrayList) IsEmpty() bool {
	return al.Size() == nought
}

func (al *ArrayList) Iterator() iterator.Iterator {
	return newArrayListIterator(al)
}

func (al *ArrayList) LastIndexOf(e interface{}) (int, error) {
	if al.IsEmpty() {
		return utils.InvalidIndex, fmt.Errorf("list is empty")
	}

	if err := al.isValidType(e); err != nil {
		return utils.InvalidIndex, err
	}

	i := al.Size() - 1
	for i >= 0 {
		if al.Get(i) == e {
			return i, nil
		}
		i--
	}

	return utils.InvalidIndex, fmt.Errorf("element %v not found in the list", e)
}

func (al *ArrayList) Remove(e interface{}) (bool, error) {
	if al.IsEmpty() {
		return false, fmt.Errorf("list is empty")
	}

	if err := al.isValidType(e); err != nil {
		return false, err
	}

	i, err := al.IndexOf(e)
	if err != nil || i == utils.InvalidIndex {
		return false, err
	}

	removeAt(al, i)

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

func removeAt(al *ArrayList, i int) {
	for j := i; j < al.Size(); j++ {
		al.data[j] = al.data[j+1]
	}

	al.size--
	checkDecreaseCapacity(al)
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

	if to < 0 || to > al.Size() {
		return false, liberror.NewIndexOutOfBoundError(to)
	}

	idx := to
	for i := from; i < al.Size(); i++ {
		if idx < al.Size() {
			al.data[i] = al.data[idx]
			idx++
		} else {
			al.data[i] = interface{}(nil)
		}
	}

	al.size -= to - from

	checkDecreaseCapacity(al)

	return true, nil
}

func (al *ArrayList) Replace(old, new interface{}) error {
	if al.IsEmpty() {
		return errors.New("list is empty")
	}

	oldType := utils.GetTypeName(old)
	newType := utils.GetTypeName(new)

	if al.typeURL != oldType {
		return liberror.NewTypeMismatchError(al.typeURL, oldType)
	}

	if al.typeURL != newType {
		return liberror.NewTypeMismatchError(al.typeURL, newType)
	}

	id, err := al.IndexOf(old)
	if err != nil {
		return err
	}

	al.data[id] = new

	return nil
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
	return al.size
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

	if e < 0 || e > al.Size() {
		return nil, liberror.NewIndexOutOfBoundError(e)
	}

	tempList, err := NewArrayList()
	if err != nil {
		return nil, err
	}

	for i := s; i < e; i++ {
		if err := tempList.Add(al.Get(i)); err != nil {
			return nil, err
		}
	}

	return tempList, nil
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
	if utils.GetTypeName(e) != al.typeURL {
		return liberror.NewTypeMismatchError(al.typeURL, utils.GetTypeName(e))
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

		if i == utils.InvalidIndex {
			continue
		}

		idx[i] = true
	}

	sz := al.Size()
	tempData := make([]interface{}, al.capacity)
	k := 0
	for i := 0; i < sz; i++ {
		if inverse {
			if idx[i] {
				tempData[k] = al.Get(i)
				k++
			}
		} else {
			if !idx[i] {
				tempData[k] = al.Get(i)
				k++
			}
		}
	}

	al.data = tempData

	if inverse {
		al.size = len(l)
	} else {
		al.size -= len(l)
	}

	checkDecreaseCapacity(al)

	return true, nil

}

func checkIncreaseCapacity(al *ArrayList) {
	if al.size >= int(float64(al.capacity)*al.upperLoadFactor) {
		al.capacity *= al.scalingFactor
		al.data = resize(al.capacity, al.data)
	}
}

func checkDecreaseCapacity(al *ArrayList) {
	if al.capacity != initialCapacity && al.size <= int(float64(al.capacity)*al.lowerLoadFactor) {
		al.capacity /= al.scalingFactor
		al.data = resize(al.capacity, al.data)
	}
}

func resize(capacity int, data []interface{}) []interface{} {
	temp := make([]interface{}, capacity)

	sz := len(data)
	if len(temp) < sz {
		sz = len(temp)
	}

	for i := 0; i < sz; i++ {
		temp[i] = data[i]
	}

	return temp
}
