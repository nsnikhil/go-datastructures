package list

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/nsnikhil/go-datastructures/utils"
)

type node struct {
	data interface{}
	next *node
	prev *node
}

func newNode(data interface{}) *node {
	return &node{
		data: data,
	}
}

type LinkedList struct {
	typeURL string
	first   *node
	last    *node
}

func NewLinkedList(data ...interface{}) (*LinkedList, error) {
	if len(data) == 0 {
		return &LinkedList{
			typeURL: utils.NA,
		}, nil
	}

	typeURL := utils.GetTypeName(data[0])

	for i := 1; i < len(data); i++ {
		if utils.GetTypeName(data[i]) != typeURL {
			return nil, liberror.NewTypeMismatchError(typeURL, utils.GetTypeName(data[i]))
		}
	}

	ll := &LinkedList{
		typeURL: typeURL,
		first:   newNode(data[0]),
	}

	curr := ll.first
	for i := 1; i < len(data); i++ {
		prev := curr

		curr.next = newNode(data[i])
		curr = curr.next
		curr.prev = prev
	}

	ll.last = curr

	return ll, nil
}

func (ll *LinkedList) Add(e interface{}) error {
	if ll.typeURL == utils.NA {
		ll.first = newNode(e)
		ll.last = ll.first
		ll.typeURL = utils.GetTypeName(e)
		return nil
	}

	if err := ll.isValidType(e); err != nil {
		return err
	}

	tempNode := newNode(e)

	if ll.IsEmpty() {
		ll.first = tempNode
		ll.last = tempNode
		return nil
	}

	temp := ll.last

	temp.next = tempNode
	temp.next.prev = temp

	ll.last = tempNode

	return nil

}

func (ll *LinkedList) AddAt(i int, e interface{}) error {
	if ll.typeURL == utils.NA {
		ll.first = newNode(e)
		ll.last = ll.first
		ll.typeURL = utils.GetTypeName(e)
		return nil
	}

	if err := ll.isValidIndex(i); err != nil {
		return err
	}

	if err := ll.isValidType(e); err != nil {
		return err
	}

	tempNode := newNode(e)
	curr := ll.first

	if i == 0 {
		tempNode.next = curr
		curr.prev = tempNode
		ll.first = tempNode
		return nil
	}

	for i > 1 {
		i--
		curr = curr.next
	}

	currNext := curr.next
	curr.next = tempNode

	tempNode.next = currNext
	tempNode.prev = curr

	currNext.prev = tempNode

	return nil
}

func (ll *LinkedList) AddAll(l ...interface{}) error {
	if len(l) == 0 {
		return nil
	}

	for i := 0; i < len(l)-1; i++ {
		if utils.GetTypeName(l[i]) != utils.GetTypeName(l[i+1]) {
			return fmt.Errorf("type mismatch : all elements must be of same type")
		}
	}

	idx := 0
	var curr *node

	if ll.typeURL == utils.NA {
		ll.first = newNode(l[idx])
		ll.last = ll.first
		ll.typeURL = utils.GetTypeName(l[idx])

		idx++
		curr = ll.first
	} else {

		if err := ll.isValidType(l[idx]); err != nil {
			return err
		}

		if ll.IsEmpty() {
			ll.first = newNode(l[idx])
			ll.last = ll.first
			idx++
		}

		curr = ll.first

		for curr.next != nil {
			curr = curr.next
		}

	}

	for ; idx < len(l); idx++ {
		prev := curr

		curr.next = newNode(l[idx])
		curr = curr.next
		curr.prev = prev
	}

	ll.last = curr

	return nil
}

func (ll *LinkedList) AddFirst(e interface{}) error {
	curr := ll.first

	if ll.typeURL == utils.NA {
		ll.typeURL = utils.GetTypeName(e)
	} else {
		if err := ll.isValidType(e); err != nil {
			return err
		}
	}

	tempNode := newNode(e)

	ll.first = tempNode
	ll.first.next = curr

	if curr != nil {
		curr.prev = tempNode
	} else {
		ll.last = tempNode
	}

	return nil
}

func (ll *LinkedList) AddLast(e interface{}) error {
	return ll.Add(e)
}

func (ll *LinkedList) Clear() {
	ll.first = nil
	ll.last = nil
}

func (ll *LinkedList) Clone() (List, error) {
	if ll.IsEmpty() {
		return NewLinkedList()
	}
	return ll.SubList(0, ll.Size())
}

func (ll *LinkedList) Contains(e interface{}) (bool, error) {
	_, err := newFinder(doubly).search(ll, e)
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

	temp := ll.first

	for i != 0 {
		i--
		temp = temp.next
	}

	return temp.data
}

func (ll *LinkedList) GetFirst() interface{} {
	if ll.IsEmpty() {
		return nil
	}

	return ll.first.data
}

func (ll *LinkedList) GetLast() interface{} {
	if ll.IsEmpty() {
		return nil
	}

	return ll.last.data
}

func (ll *LinkedList) IndexOf(e interface{}) (int, error) {
	if ll.IsEmpty() {
		return -1, fmt.Errorf("list is empty")
	}

	if err := ll.isValidType(e); err != nil {
		return utils.InvalidIndex, err
	}

	i, _ := newFinder(doubly).search(ll, e)
	if i == utils.InvalidIndex {
		return utils.InvalidIndex, fmt.Errorf("element %v not found in the list", e)
	}

	return i, nil

}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Size() == 0
}

func (ll *LinkedList) Iterator() iterator.Iterator {
	return newLinkedListIterator(ll)
}

func (ll *LinkedList) DescendingIterator() iterator.Iterator {
	return newLinkedListDescendingIterator(ll)
}

func (ll *LinkedList) LastIndexOf(e interface{}) (int, error) {
	if ll.IsEmpty() {
		return -1, fmt.Errorf("list is empty")
	}

	if err := ll.isValidType(e); err != nil {
		return utils.InvalidIndex, err
	}

	it := ll.DescendingIterator()

	i := utils.InvalidIndex
	count := ll.Size() - 1

	for it.HasNext() {
		t := it.Next()
		if t == e {
			i = count
			break
		}
		count--
	}

	if i == utils.InvalidIndex {
		return utils.InvalidIndex, fmt.Errorf("element %v not found in the list", e)
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
	if err != nil || i == utils.InvalidIndex {
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

	var curr *node
	curr = ll.first

	for curr != nil && i > 0 {
		curr = curr.next
		i--
	}

	if curr.prev == nil {
		ll.first = curr.next

		if curr.next != nil {
			curr.next.prev = curr.prev
		} else {
			ll.last = ll.first
		}

		return curr.data, nil
	}

	if curr.next == nil {
		ll.last = curr.prev
		curr.prev.next = curr.next

		return curr.data, nil
	}

	curr.prev.next = curr.next
	curr.next.prev = curr.prev

	return curr.data, nil
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
	if ll.IsEmpty() {
		return nil, fmt.Errorf("list is empty")
	}

	curr := ll.last
	val := curr.data

	if curr.prev == nil {
		ll.Clear()
		return val, nil
	}

	ll.last = curr.prev
	curr.prev.next = nil

	return val, nil
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

func (ll *LinkedList) ReplaceAll(uo operator.UnaryOperator) error {
	temp := ll.first

	for temp != nil {
		e := uo.Apply(temp.data)
		if err := ll.isValidType(e); err != nil {
			return err
		}

		temp.data = e
		temp = temp.next
	}

	return nil
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

	curr := ll.first

	for i > 0 {
		i--
		curr = curr.next
	}

	curr.data = e

	return curr.data, nil
}

func (ll *LinkedList) Size() int {
	count := 0
	temp := ll.first

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
	temp := ll.first

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

	temp := ll.first

	n := s

	for s > 0 {
		s--
		temp = temp.next
	}

	for temp != nil && n < e {
		if err = tempLL.AddLast(temp.data); err != nil {
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
	if utils.GetTypeName(e) != ll.typeURL {
		return liberror.NewTypeMismatchError(ll.typeURL, utils.GetTypeName(e))
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

	curr := ll.first

	for curr != nil {

		shouldRemove := false

		if inverse {
			shouldRemove = !dataMap[curr.data]
		} else {
			shouldRemove = dataMap[curr.data]
		}

		if shouldRemove {

			if curr.prev == nil {
				ll.first = curr.next

				if curr.next != nil {
					curr.next.prev = curr.prev
				} else {
					ll.last = ll.first
					break
				}

			} else if curr.next == nil {

				ll.last = curr.prev
				curr.prev.next = curr.next

			} else {
				curr.prev.next = curr.next
				curr.next.prev = curr.prev
			}

		}

		curr = curr.next

	}

	return true, nil
}

type linkedListIterator struct {
	currNode *node
	ll       *LinkedList
}

func newLinkedListIterator(ll *LinkedList) *linkedListIterator {
	return &linkedListIterator{
		currNode: ll.first,
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

type linkedListDescendingIterator struct {
	currNode *node
	ll       *LinkedList
}

func newLinkedListDescendingIterator(ll *LinkedList) *linkedListDescendingIterator {
	return &linkedListDescendingIterator{
		currNode: ll.last,
		ll:       ll,
	}
}

func (ll *linkedListDescendingIterator) HasNext() bool {
	return ll.currNode != nil
}

func (ll *linkedListDescendingIterator) Next() interface{} {
	if ll.currNode == nil {
		return nil
	}

	e := ll.currNode.data

	ll.currNode = ll.currNode.prev

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
