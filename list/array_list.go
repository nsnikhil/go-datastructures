package list

import (
	"datastructures/functions/comparator"
	"datastructures/functions/iterator"
	"datastructures/functions/operator"
	"fmt"
	"reflect"
)

const (
	na                = "na"
	typeMisMatchError = "every data in List should be of same type"
	invalidIndex      = -1
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

	typeURL := reflect.TypeOf(data[0]).Name()

	for i := 1; i < len(data); i++ {
		if reflect.TypeOf(data[i]).Name() != typeURL {
			return nil, fmt.Errorf(typeMisMatchError)
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
		al.typeURL = reflect.TypeOf(e).Name()
		return nil
	}

	if !isValidType(e, al) {
		return fmt.Errorf("type mismatch : expected %s got %s", al.typeURL, reflect.TypeOf(e).Name())
	}

	al.data = append(al.data, e)
	return nil
}

func (al *ArrayList) AddAt(i int, e interface{}) error {
	if !isValidType(e, al) {
		return fmt.Errorf("type mismatch : expected %s got %s", al.typeURL, reflect.TypeOf(e).Name())
	}

	if !isValidIndex(i, al) {
		return fmt.Errorf("invalid index %d", i)
	}

	al.data = append(al.data[:i], append([]interface{}{e}, al.data[i:]...)...)

	return nil
}

func (al *ArrayList) AddAll(l ...interface{}) error {
	for _, e := range l {
		if !isValidType(e, al) {
			return fmt.Errorf("failed to Add elements due to invalid type %s", reflect.TypeOf(e).Name())
		}
	}

	al.data = append(al.data, l...)

	return nil
}

func (al *ArrayList) Clear() {
	al.data = nil
}

//TODO SHOULD IT THROW ERROR WHEN TYPE MISMATCH ?
func (al *ArrayList) Contains(e interface{}) (bool, error) {
	if !isValidType(e, al) {
		return false, fmt.Errorf("type mismatch : expected %s got %s", al.typeURL, reflect.TypeOf(e).Name())
	}

	for _, d := range al.data {
		if d == e {
			return true, nil
		}
	}

	return false, fmt.Errorf("element %v not found", e)
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
	if al.Size() == 0 || !isValidIndex(i, al) {
		return nil
	}
	return al.data[i]
}

//TODO SHOULD IT THROW ERROR WHEN TYPE MISMATCH ?
func (al *ArrayList) IndexOf(e interface{}) (int, error) {
	if !isValidType(e, al) {
		return invalidIndex, fmt.Errorf("type mismatch : expected %s got %s", al.typeURL, reflect.TypeOf(e).Name())
	}

	i := searchList(al, e)
	if i == invalidIndex {
		return invalidIndex, fmt.Errorf("failed to find element %v in List", e)
	}

	return i, nil
}

func (al *ArrayList) IsEmpty() bool {
	return al.Size() == 0
}

func (al *ArrayList) Iterator() iterator.Iterator {
	return newArrayListIterator(al)
}

//TODO SHOULD IT THROW ERROR WHEN TYPE MISMATCH ?
func (al *ArrayList) LastIndexOf(e interface{}) (int, error) {
	if !isValidType(e, al) {
		return invalidIndex, fmt.Errorf("type mismatch : expected %s got %s", al.typeURL, reflect.TypeOf(e).Name())
	}

	i := al.Size() - 1
	for i >= 0 {
		if al.Get(i) == e {
			return i, nil
		}
		i--
	}

	return invalidIndex, fmt.Errorf("element %v is not present in List", e)
}

//TODO SHOULD IT THROW ERROR WHEN TYPE MISMATCH ?
func (al *ArrayList) Remove(e interface{}) (bool, error) {
	if !isValidType(e, al) {
		return false, fmt.Errorf("type mismatch : expected %s got %s", al.typeURL, reflect.TypeOf(e).Name())
	}

	i, err := al.IndexOf(e)
	if err != nil || i == invalidIndex {
		return false, err
	}

	al.data = append(al.data[0:i], al.data[i+1:]...)

	return true, nil
}

func (al *ArrayList) RemoveAt(i int) (interface{}, error) {
	if !isValidIndex(i, al) {
		return nil, fmt.Errorf("invalid index %d", i)
	}

	e := al.Get(i)
	removeAt(al, i)

	return e, nil
}

func (al *ArrayList) RemoveAll(l ...interface{}) (bool, error) {
	return filter(al, false, l...)
}

func (al *ArrayList) ReplaceAll(uo operator.UnaryOperator) error {
	sz := al.Size()
	for i := 0; i < sz; i++ {
		e := uo.Apply(al.Get(i))
		if !isValidType(e, al) {
			return fmt.Errorf("type mismatch : expected %s got %s", al.typeURL, reflect.TypeOf(e).Name())
		}

		if _, err := al.Set(i, e); err != nil {
			return err
		}
	}

	return nil
}

func (al *ArrayList) RetainAll(l ...interface{}) (bool, error) {
	return filter(al, true, l...)
}

func (al *ArrayList) Set(i int, e interface{}) (interface{}, error) {
	if !isValidIndex(i, al) {
		return nil, fmt.Errorf("failed to Set value %v due to invalid index %d", e, i)
	}

	al.data[i] = e
	return e, nil
}

func (al *ArrayList) Size() int {
	return len(al.data)
}

func (al *ArrayList) Sort(c comparator.Comparator) {
	sortList(al, c)
}

func (al *ArrayList) SubList(s int, e int) (List, error) {
	if !isValidIndex(s, al) {
		return nil, fmt.Errorf("invalid index %d", s)
	}

	if !isValidIndex(e, al) {
		return nil, fmt.Errorf("invalid index %d", e)
	}

	tempData := make([]interface{}, 0)
	for i := s; i <= e; i++ {
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
	ali.currentIndex += 1

	return e
}

func isValidIndex(i int, al *ArrayList) bool {
	return i >= 0 && i < al.Size()
}

func isValidType(e interface{}, al *ArrayList) bool {
	return reflect.TypeOf(e).Name() == al.typeURL
}

func filter(al *ArrayList, inverse bool, l ...interface{}) (bool, error) {
	for _, e := range l {
		if !isValidType(e, al) {
			return false, fmt.Errorf("type mismatch : expected %s got %s", al.typeURL, reflect.TypeOf(e).Name())
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
