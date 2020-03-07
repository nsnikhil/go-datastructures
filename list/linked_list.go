package list

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/liberror"
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

type LinkedList struct {
	typeURL string
	root    *node
}

func NewLinkedList(data ...interface{}) (*LinkedList, error) {
	if len(data) == 0 {
		return &LinkedList{
			typeURL: na,
		}, nil
	}

	typeURL := getTypeName(data[0])

	for i := 1; i < len(data); i++ {
		if getTypeName(data[i]) != typeURL {
			return nil, liberror.NewTypeMismatchError(typeURL, getTypeName(data[i]))
		}
	}

	ll := &LinkedList{
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

func (ll *LinkedList) Add(e interface{}) error {
	if ll.typeURL == na {
		ll.root = newNode(e)
		ll.typeURL = getTypeName(e)
		return nil
	}

	if err := ll.isValidType(e); err != nil {
		return err
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

func (ll *LinkedList) AddAt(i int, e interface{}) error {
	if ll.typeURL == na {
		ll.root = newNode(e)
		ll.typeURL = getTypeName(e)
		return nil
	}

	if err := ll.isValidIndex(i); err != nil {
		return err
	}

	if err := ll.isValidType(e); err != nil {
		return err
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

func (ll *LinkedList) AddAll(l ...interface{}) error {
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

		if err := ll.isValidType(l[idx]); err != nil {
			return err
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

func (ll *LinkedList) AddFirst(e interface{}) error {
	temp := ll.root

	if ll.typeURL == na {
		ll.typeURL = getTypeName(e)
	} else {
		if err := ll.isValidType(e); err != nil {
			return err
		}
	}

	ll.root = newNode(e)
	ll.root.next = temp

	return nil
}

func (ll *LinkedList) AddLast(e interface{}) error {
	return ll.Add(e)
}

func (ll *LinkedList) Clear() {
	ll.root = nil
}

func (ll *LinkedList) Clone() (List, error) {
	if ll.IsEmpty() {
		return NewLinkedList()
	}
	return ll.SubList(0, ll.Size())
}

func (ll *LinkedList) Contains(e interface{}) (bool, error) {
	_, err := newFinder().search(ll, e)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ll *LinkedList) ContainsAll(l ...interface{}) (bool, error) {
	for _, e := range l {
		if _, err := ll.Contains(e); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (ll *LinkedList) Get(i int) interface{} {
	if err := ll.isValidIndex(i); ll.IsEmpty() || err != nil {
		return nil
	}

	temp := ll.root

	for i != 0 {
		i--
		temp = temp.next
	}

	return temp.data
}

func (ll *LinkedList) GetFirst() interface{} {
	return ll.Get(0)
}

func (ll *LinkedList) GetLast() interface{} {
	return ll.Get(ll.Size() - 1)
}

func (ll *LinkedList) IndexOf(e interface{}) (int, error) {
	if ll.IsEmpty() {
		return -1, fmt.Errorf("list is empty")
	}

	if err := ll.isValidType(e); err != nil {
		return invalidIndex, err
	}

	i, _ := newFinder().search(ll, e)
	if i == invalidIndex {
		return invalidIndex, fmt.Errorf("element %v not found in the list", e)
	}

	return i, nil

}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Size() == 0
}

func (ll *LinkedList) Iterator() iterator.Iterator {
	return newLinkedListIterator(ll)
}

func (ll *LinkedList) LastIndexOf(e interface{}) (int, error) {
	if ll.IsEmpty() {
		return -1, fmt.Errorf("list is empty")
	}

	if err := ll.isValidType(e); err != nil {
		return invalidIndex, err
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

func (ll *LinkedList) Remove(e interface{}) (bool, error) {
	if ll.IsEmpty() {
		return false, fmt.Errorf("list is empty")
	}

	if err := ll.isValidType(e); err != nil {
		return false, err
	}

	i, err := ll.IndexOf(e)
	if err != nil || i == invalidIndex {
		return false, err
	}

	if _, err := ll.RemoveAt(i); err != nil {
		return false, err
	}

	return true, nil
}

func (ll *LinkedList) RemoveAt(i int) (interface{}, error) {
	if ll.IsEmpty() {
		return nil, fmt.Errorf("list is empty")
	}

	if err := ll.isValidIndex(i); err != nil {
		return nil, err
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

func (ll *LinkedList) RemoveAll(l ...interface{}) (bool, error) {
	return filterLinkedList(ll, false, l...)
}

func (ll *LinkedList) RemoveFirst() (interface{}, error) {
	return ll.RemoveAt(0)
}

func (ll *LinkedList) RemoveFirstOccurrence(e interface{}) (bool, error) {
	return ll.Remove(e)
}

func (ll *LinkedList) RemoveLast() (interface{}, error) {
	return ll.RemoveAt(ll.Size() - 1)
}

func (ll *LinkedList) RemoveLastOccurrence(e interface{}) (bool, error) {
	i, err := ll.LastIndexOf(e)
	if err != nil {
		return false, err
	}

	if _, err := ll.RemoveAt(i); err != nil {
		return false, err
	}

	return true, nil
}

func (ll *LinkedList) ReplaceAll(uo operator.UnaryOperator) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("type mismatch : %v", r)
		}
	}()

	temp := ll.root

	for temp != nil {
		e := uo.Apply(temp.data)
		//TODO CAN YOU JUST USE ERR INSTEAD OF NEW VARIABLE
		if typeError := ll.isValidType(e); typeError != nil {
			err = typeError
			return
		}

		temp.data = e
		temp = temp.next
	}

	return
}

func (ll *LinkedList) RetainAll(l ...interface{}) (bool, error) {
	return filterLinkedList(ll, true, l...)
}

func (ll *LinkedList) Set(i int, e interface{}) (interface{}, error) {
	if ll.IsEmpty() {
		return nil, fmt.Errorf("list is empty")
	}

	if err := ll.isValidIndex(i); err != nil {
		return nil, err
	}

	if err := ll.isValidType(e); err != nil {
		return nil, err
	}

	temp := ll.root

	for i-1 > 0 {
		i--
		temp = temp.next
	}

	temp.data = e

	return temp.data, nil
}

func (ll *LinkedList) Size() int {
	count := 0
	temp := ll.root

	for temp != nil {
		temp = temp.next
		count++
	}

	return count
}

func (ll *LinkedList) Sort(c comparator.Comparator) {
	al := ll.ToArrayList()
	al.Sort(c)

	it := al.Iterator()
	temp := ll.root

	for temp != nil && it.HasNext() {
		temp.data = it.Next()
		temp = temp.next
	}
}

func (ll *LinkedList) SubList(s int, e int) (List, error) {
	if e < s {
		return nil, fmt.Errorf("end cannot be smaller than start")
	}

	if err := ll.isValidIndex(s); err != nil {
		return nil, err
	}

	//TODO USE IS VALID INDEX METHOD HERE
	if e < 0 || e > ll.Size() {
		return nil, liberror.NewIndexOutOfBoundError(e)
	}

	tempLL, err := NewLinkedList()
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

func (ll *LinkedList) isValidIndex(i int) error {
	if i < 0 || i >= ll.Size() {
		return liberror.NewIndexOutOfBoundError(i)
	}
	return nil
}

func (ll *LinkedList) isValidType(e interface{}) error {
	if getTypeName(e) != ll.typeURL {
		return liberror.NewTypeMismatchError(ll.typeURL, getTypeName(e))
	}
	return nil
}

func filterLinkedList(ll *LinkedList, inverse bool, l ...interface{}) (bool, error) {
	if ll.IsEmpty() {
		return false, fmt.Errorf("list is empty")
	}

	for _, e := range l {
		if err := ll.isValidType(e); err != nil {
			return false, err
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
	ll       *LinkedList
}

func newLinkedListIterator(ll *LinkedList) *linkedListIterator {
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

func (ll *LinkedList) ToArrayList() *ArrayList {
	var e []interface{}
	it := ll.Iterator()

	for it.HasNext() {
		e = append(e, it.Next())
	}

	al, _ := NewArrayList(e...)
	return al
}
