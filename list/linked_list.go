package list

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
)

type node struct {
	data interface{}
	next *node
}

func newNode(data interface{}) *node {
	return &node{
		data: data,
	}
}

type linkedList struct {
	typeURL string
	root    *node
}

func newLinkedList(data ...interface{}) (*linkedList, error) {
	if len(data) == 0 {
		return &linkedList{
			typeURL: na,
		}, nil
	}

	typeURL := getTypeName(data[0])

	for i := 1; i < len(data); i++ {
		if getTypeName(data[i]) != typeURL {
			return nil, fmt.Errorf("type mismatch : expected %s got %s", typeURL, getTypeName(data[i]))
		}
	}

	ll := &linkedList{
		typeURL: typeURL,
		root:    newNode(data[0]),
	}

	temp := ll.root
	for i := 1; i < len(data); i++ {
		temp.next = newNode(data[i])
		temp = temp.next
	}

	return ll, nil
}

func (ll *linkedList) Add(e interface{}) error {
	if ll.typeURL == na {
		ll.root = newNode(e)
		ll.typeURL = getTypeName(e)
		return nil
	}

	if !ll.isValidType(e) {
		return fmt.Errorf("type mismatch : expected %s got %s", ll.typeURL, getTypeName(e))
	}

	if ll.IsEmpty() {
		ll.root = newNode(e)
		return nil
	}

	temp := ll.root

	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode(e)

	return nil
}

func (ll *linkedList) AddAt(i int, e interface{}) error {
	if ll.IsEmpty() && ll.typeURL == na {
		ll.root = newNode(e)
		ll.typeURL = getTypeName(e)
		return nil
	}

	if !ll.isValidIndex(i) {
		return fmt.Errorf("invalid index %d", i)
	}

	if !ll.isValidType(e) {
		return fmt.Errorf("type mismatch : expected %s got %s", ll.typeURL, getTypeName(e))
	}

	tempNode := newNode(e)

	if i == 0 {
		tempNode.next = ll.root
		ll.root = tempNode
		return nil
	}

	temp := ll.root

	for i > 1 {
		i--
		temp = temp.next
	}

	tempNext := temp.next
	temp.next = tempNode
	tempNode.next = tempNext

	return nil
}

func (ll *linkedList) AddAll(l ...interface{}) error {
	if len(l) == 0 {
		return nil
	}

	for i := 0; i < len(l)-1; i++ {
		if getTypeName(l[i]) != getTypeName(l[i+1]) {
			return fmt.Errorf("type mismatch : all elements must be of same type")
		}
	}

	idx := 0
	var temp *node

	if ll.typeURL == na {
		ll.root = newNode(l[idx])
		ll.typeURL = getTypeName(l[idx])

		idx++
		temp = ll.root
	} else {

		if !ll.isValidType(l[idx]) {
			return fmt.Errorf("type mismatch : expected %s got %s", ll.typeURL, getTypeName(l[idx]))
		}

		if ll.IsEmpty() {
			ll.root = newNode(l[idx])
			idx++
		}

		temp = ll.root

		for temp.next != nil {
			temp = temp.next
		}

	}

	for ; idx < len(l); idx++ {
		temp.next = newNode(l[idx])
		temp = temp.next
	}

	return nil
}

func (ll *linkedList) Clear() {
	ll.root = nil
}

func (ll *linkedList) Contains(e interface{}) (bool, error) {
	_, err := newFinder().search(ll, e)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ll *linkedList) ContainsAll(l ...interface{}) (bool, error) {
	for _, e := range l {
		if _, err := ll.Contains(e); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (ll *linkedList) Get(i int) interface{} {
	if ll.IsEmpty() || !ll.isValidIndex(i) {
		return nil
	}

	temp := ll.root

	for i != 0 {
		i--
		temp = temp.next
	}

	return temp.data
}

func (ll *linkedList) IndexOf(e interface{}) (int, error) {
	if ll.IsEmpty() {
		return -1, fmt.Errorf("list is empty")
	}

	if !ll.isValidType(e) {
		return invalidIndex, fmt.Errorf("type mismatch : expected %s got %s", ll.typeURL, getTypeName(e))
	}

	i, _ := newFinder().search(ll, e)
	if i == invalidIndex {
		return invalidIndex, fmt.Errorf("element %v not found in the list", e)
	}

	return i, nil

}

func (ll *linkedList) IsEmpty() bool {
	return ll.Size() == 0
}

func (ll *linkedList) Iterator() iterator.Iterator {
	return newLinkedListIterator(ll)
}

func (ll *linkedList) LastIndexOf(e interface{}) (int, error) {
	if ll.IsEmpty() {
		return -1, fmt.Errorf("list is empty")
	}

	if !ll.isValidType(e) {
		return invalidIndex, fmt.Errorf("type mismatch : expected %s got %s", ll.typeURL, getTypeName(e))
	}

	i := invalidIndex
	count := 0
	temp := ll.root

	for temp != nil {
		if temp.data == e {
			i = count
		}
		temp = temp.next
		count++
	}

	if i == invalidIndex {
		return invalidIndex, fmt.Errorf("element %v not found in the list", e)
	}

	return i, nil
}

func (ll *linkedList) Remove(e interface{}) (bool, error) {
	if ll.IsEmpty() {
		return false, fmt.Errorf("list is empty")
	}

	if !ll.isValidType(e) {
		return false, fmt.Errorf("type mismatch : expected %s got %s", ll.typeURL, getTypeName(e))
	}

	i, err := ll.IndexOf(e)
	if err != nil || i == invalidIndex {
		return false, err
	}

	if ll.root.data == e {
		ll.root = ll.root.next
		return true, nil
	}

	temp := ll.root

	for temp.next.data != e {
		temp = temp.next
	}

	temp.next = temp.next.next

	return true, nil
}

func (ll *linkedList) RemoveAt(i int) (interface{}, error) {
	if ll.IsEmpty() {
		return nil, fmt.Errorf("list is empty")
	}

	if !ll.isValidIndex(i) {
		return nil, fmt.Errorf("invalid index %d", i)
	}

	temp := ll.root

	if i == 0 {
		ll.root = ll.root.next
		return temp.data, nil
	}

	for i-1 > 0 {
		temp = temp.next
		i--
	}

	e := temp.next

	temp.next = temp.next.next

	return e.data, nil
}

func (ll *linkedList) RemoveAll(l ...interface{}) (bool, error) {
	return filterLinkedList(ll, false, l...)
}

func (ll *linkedList) ReplaceAll(uo operator.UnaryOperator) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("type mismatch : %v", r)
		}
	}()

	temp := ll.root

	for temp != nil {
		e := uo.Apply(temp.data)
		if !ll.isValidType(e) {
			err = fmt.Errorf("type mismatch : expected %s got %s", ll.typeURL, getTypeName(e))
			return
		}

		temp.data = e
		temp = temp.next
	}

	return
}

func (ll *linkedList) RetainAll(l ...interface{}) (bool, error) {
	return filterLinkedList(ll, true, l...)
}

func (ll *linkedList) Set(i int, e interface{}) (interface{}, error) {
	if ll.IsEmpty() {
		return nil, fmt.Errorf("list is empty")
	}

	if !ll.isValidIndex(i) {
		return nil, fmt.Errorf("invalid index %d", i)
	}

	if !ll.isValidType(e) {
		return nil, fmt.Errorf("type mismatch : expected %s got %s", ll.typeURL, getTypeName(e))
	}

	temp := ll.root

	for i-1 > 0 {
		i--
		temp = temp.next
	}

	temp.data = e

	return temp.data, nil
}

func (ll *linkedList) Size() int {
	count := 0
	temp := ll.root

	for temp != nil {
		temp = temp.next
		count++
	}

	return count
}

func (ll *linkedList) Sort(c comparator.Comparator) {
	al := ll.toArrayList()
	al.Sort(c)

	it := al.Iterator()
	temp := ll.root

	for temp != nil && it.HasNext() {
		temp.data = it.Next()
		temp = temp.next
	}
}

func (ll *linkedList) SubList(s int, e int) (List, error) {
	if e < s {
		return nil, fmt.Errorf("end cannot be smaller than start")
	}

	if !ll.isValidIndex(s) {
		return nil, fmt.Errorf("invalid index %d", s)
	}

	if e < 0 || e > ll.Size() {
		return nil, fmt.Errorf("invalid index %d", e)
	}

	tempLL, err := newLinkedList()
	if err != nil {
		return nil, err
	}

	temp := ll.root

	n := s

	for s > 0 {
		s--
		temp = temp.next
	}

	for temp != nil && n < e {
		if err = tempLL.Add(temp.data); err != nil {
			return nil, err
		}
		temp = temp.next
		n++
	}

	return tempLL, nil
}

func (ll *linkedList) isValidIndex(i int) bool {
	return i >= 0 && i < ll.Size()
}

func (ll *linkedList) isValidType(e interface{}) bool {
	return getTypeName(e) == ll.typeURL
}

func filterLinkedList(ll *linkedList, inverse bool, l ...interface{}) (bool, error) {
	if len(l) == 0 {
		return false, nil
	}

	if ll.IsEmpty() {
		return false, fmt.Errorf("list is empty")
	}

	for _, e := range l {
		if !ll.isValidType(e) {
			return false, fmt.Errorf("type mismatch : expected %s got %s", ll.typeURL, getTypeName(e))
		}
	}

	dataMap := make(map[interface{}]bool)
	for _, e := range l {
		dataMap[e] = true
	}

	temp := ll.root
	var prev *node
	isLast := false

	for temp != nil {

		isPresent := false

		if inverse {
			isPresent = !dataMap[temp.data]
		} else {
			isPresent = dataMap[temp.data]
		}

		if isPresent {

			/*
				since we cannot perform *temp = *temp.next if temp is the last element in the list
				we break out of the loop
			*/
			if temp.next != nil {
				*temp = *temp.next
			} else {
				isLast = true
				break
			}

		} else {
			prev = temp
			temp = temp.next
		}

	}

	/*
		if the last element was one of the elements to be removed to change that element prev to point to nil
		also if list only contains one element and that element has to be removed then set root as nil
	*/
	if isLast {
		if prev != nil {
			prev.next = nil
		} else {
			ll.root = nil
		}
	}

	return true, nil
}

type linkedListIterator struct {
	currNode *node
	ll       *linkedList
}

func newLinkedListIterator(ll *linkedList) *linkedListIterator {
	return &linkedListIterator{
		currNode: ll.root,
		ll:       ll,
	}
}

func (ll *linkedListIterator) HasNext() bool {
	return ll.currNode != nil
}

func (ll *linkedListIterator) Next() interface{} {
	if ll.currNode == nil {
		return nil
	}

	e := ll.currNode.data

	ll.currNode = ll.currNode.next

	return e
}

func (ll *linkedList) toArrayList() *ArrayList {
	var e []interface{}
	it := ll.Iterator()

	for it.HasNext() {
		e = append(e, it.Next())
	}

	al, _ := NewArrayList(e...)
	return al
}
