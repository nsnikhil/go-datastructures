package list

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/liberr"
	"github.com/nsnikhil/go-datastructures/utils"
	"sync"
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
	count   int
	mt      *sync.Mutex

	first *node
	last  *node
}

func NewLinkedList(data ...interface{}) (*LinkedList, error) {
	ll := &LinkedList{typeURL: utils.NA, mt: new(sync.Mutex), count: 0}

	err := ll.AddAll(data...)
	if err != nil {
		return nil, err
	}

	return ll, nil
}

func (ll *LinkedList) Add(e interface{}) error {
	return ll.addAt(e, ll.Size())
}

func (ll *LinkedList) AddAt(i int, e interface{}) error {
	return ll.addAt(e, i)
}

func (ll *LinkedList) AddAll(data ...interface{}) error {
	sz := len(data)
	if sz == 0 {
		return nil
	}

	err := utils.AreAllSameType(data...)
	if err != nil {
		return err
	}

	lsz := ll.Size()

	for i := 0; i < sz; i++ {
		err := ll.addAt(data[i], i+lsz)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ll *LinkedList) AddFirst(e interface{}) error {
	return ll.addAt(e, 0)
}

func (ll *LinkedList) AddLast(e interface{}) error {
	return ll.addAt(e, ll.Size())
}

func (ll *LinkedList) Clear() {
	ll.first = nil
	ll.last = nil
	ll.count = 0
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

//TODO: OPTIMIZE, SHOULD FIND ALL THE ELEMENTS IN ONE PASS
func (ll *LinkedList) ContainsAll(l ...interface{}) (bool, error) {
	for _, e := range l {
		if _, err := ll.Contains(e); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (ll *LinkedList) Get(i int) interface{} {
	curr, err := ll.traverseTo(i)
	if err != nil {
		return nil
	}

	return curr.data
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
	err := checks(ll, []checkData{{empty, nil}, {typeURL, e}})
	if err != nil {
		return -1, err
	}

	curr := ll.first
	idx := 0

	for curr != nil {
		if curr.data == e {
			return idx, nil
		}

		curr = curr.next
		idx++
	}

	return -1, liberr.NotFondErrorInList(e)
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

//TODO: OPTIMIZE, MAKE IT SHORT
func (ll *LinkedList) LastIndexOf(e interface{}) (int, error) {
	err := checks(ll, []checkData{{empty, nil}, {typeURL, e}})
	if err != nil {
		return -1, err
	}

	curr := ll.last
	idx := ll.Size() - 1

	for curr != nil {
		if curr.data == e {
			return idx, nil
		}

		curr = curr.prev
		idx--
	}

	return -1, liberr.NotFondErrorInList(e)
}

func (ll *LinkedList) Remove(e interface{}) (bool, error) {
	err := checks(ll, []checkData{{empty, nil}, {typeURL, e}})
	if err != nil {
		return false, err
	}

	i, err := ll.IndexOf(e)
	if err != nil {
		return false, err
	}

	if _, err := ll.removeAt(i); err != nil {
		return false, err
	}

	return true, nil
}

func (ll *LinkedList) RemoveAt(i int) (interface{}, error) {
	data, err := ll.removeAt(i)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ll *LinkedList) RemoveAll(l ...interface{}) (bool, error) {
	return filterLinkedList(ll, false, l...)
}

func (ll *LinkedList) RemoveFirst() (interface{}, error) {
	return ll.removeAt(0)
}

func (ll *LinkedList) RemoveFirstOccurrence(e interface{}) (bool, error) {
	return ll.Remove(e)
}

func (ll *LinkedList) RemoveLast() (interface{}, error) {
	return ll.removeAt(ll.Size() - 1)
}

func (ll *LinkedList) RemoveLastOccurrence(e interface{}) (bool, error) {
	i, err := ll.LastIndexOf(e)
	if err != nil {
		return false, err
	}

	if _, err := ll.removeAt(i); err != nil {
		return false, err
	}

	return true, nil
}

func (ll *LinkedList) Replace(old, new interface{}) error {
	err := checks(ll, []checkData{{empty, nil}, {typeURL, old}, {typeURL, new}})
	if err != nil {
		return err
	}

	curr := ll.first
	for curr != nil && curr.data != old {
		curr = curr.next
	}

	if curr == nil {
		return liberr.NotFondErrorInList(old)
	}

	curr.data = new
	return nil
}

func (ll *LinkedList) ReplaceAll(uo operator.UnaryOperator) error {
	temp := ll.first

	e := uo.Apply(temp.data)
	if err := ll.isValidType(e); err != nil {
		return err
	}

	temp.data = e
	temp = temp.next

	for temp != nil {
		temp.data = uo.Apply(temp.data)
		temp = temp.next
	}

	return nil
}

func (ll *LinkedList) RetainAll(l ...interface{}) (bool, error) {
	return filterLinkedList(ll, true, l...)
}

func (ll *LinkedList) Set(i int, e interface{}) (interface{}, error) {
	err := checks(ll, []checkData{{empty, nil}, {index, i}, {typeURL, e}})
	if err != nil {
		return nil, err
	}

	curr, err := ll.traverseTo(i)
	if err != nil {
		return nil, err
	}

	curr.data = e
	return curr.data, nil
}

func (ll *LinkedList) Size() int {
	return ll.count
}

//TODO: CHANGE TO LINKED LIST MERGE SORT
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

//TODO: OPTIMIZE
func (ll *LinkedList) SubList(s int, e int) (List, error) {
	if e < s {
		return nil, liberr.InvalidOperationError("end cannot be smaller than start")
	}

	if err := ll.isValidIndex(s); err != nil {
		return nil, err
	}

	//TODO USE IS VALID INDEX METHOD HERE
	if e < 0 || e > ll.Size() {
		return nil, liberr.IndexOutOfBoundError(e)
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

func filterLinkedList(ll *LinkedList, inverse bool, l ...interface{}) (bool, error) {
	if ll.IsEmpty() {
		return false, liberr.EmptyListError
	}

	for _, e := range l {
		if err := ll.isValidType(e); err != nil {
			return false, err
		}
	}

	ll.mt.Lock()
	defer ll.mt.Unlock()

	dataMap := make(map[interface{}]bool)
	for _, e := range l {
		dataMap[e] = true
	}

	curr := ll.first

	rc := 0

	for curr != nil {

		var shouldRemove bool

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
					rc++
					break
				}

			} else if curr.next == nil {

				ll.last = curr.prev
				curr.prev.next = curr.next

			} else {
				curr.prev.next = curr.next
				curr.next.prev = curr.prev
			}

			rc++
		}

		curr = curr.next
	}

	ll.count -= rc
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

// HELPER FUNCTION FROM HERE ON
func (ll *LinkedList) addAt(e interface{}, p int) error {
	ll.mt.Lock()
	defer ll.mt.Unlock()

	sz := ll.Size()

	if p < 0 || p > sz {
		return liberr.IndexOutOfBoundError(p)
	}

	et := utils.GetTypeName(e)
	nd := newNode(e)

	if ll.typeURL == utils.NA {
		ll.typeURL = et
	} else if ll.typeURL != et {
		return liberr.TypeMismatchError(ll.typeURL, et)
	}

	if ll.IsEmpty() {
		ll.first = nd
		ll.last = nd
		ll.count++
		return nil
	}

	if p == 0 {
		nd.next = ll.first
		ll.first.prev = nd
		ll.first = nd
		ll.count++
		return nil
	}

	if p == sz {
		ll.last.next = nd
		nd.prev = ll.last
		ll.last = nd
		ll.count++
		return nil
	}

	curr, err := ll.traverseTo(p)
	if err != nil {
		return err
	}

	cp := curr.prev

	nd.next = curr
	curr.prev = nd

	cp.next = nd
	nd.prev = cp

	ll.count++

	return nil
}

func (ll *LinkedList) removeAt(p int) (interface{}, error) {
	ll.mt.Lock()
	defer ll.mt.Unlock()

	sz := ll.Size()

	if sz == 0 {
		return nil, liberr.EmptyListError
	}

	if p < 0 || p >= sz {
		return nil, liberr.IndexOutOfBoundError(p)
	}

	if sz == 1 {
		data := ll.first.data
		ll.first = nil
		ll.last = nil
		ll.count = 0
		return data, nil
	}

	if p == 0 {
		data := ll.first.data
		ll.first = ll.first.next
		ll.first.prev = nil
		ll.count--
		return data, nil
	}

	if p == sz-1 {
		data := ll.last.data
		ll.last = ll.last.prev
		ll.last.next = nil
		ll.count--
		return data, nil
	}

	curr, err := ll.traverseTo(p)
	if err != nil {
		return nil, err
	}

	cp := curr.prev
	cn := curr.next

	cp.next = cn
	cn.prev = cp
	ll.count--

	return curr.data, nil
}

func (ll *LinkedList) traverseTo(p int) (*node, error) {
	sz := ll.Size()

	if p < 0 || p >= sz {
		return nil, liberr.IndexOutOfBoundError(p)
	}

	if p == 0 {
		return ll.first, nil
	}

	if p == sz-1 {
		return ll.last, nil
	}

	md := sz / 2
	df := sz - p

	var curr *node

	if df > md {

		curr = ll.first
		k := p

		for curr != nil && k > 0 {
			curr = curr.next
			k--
		}

	} else {

		curr := ll.last
		k := sz - p

		for curr != nil && k > 0 {
			curr = curr.prev
			k--
		}

	}

	if curr == nil {
		return nil, liberr.IndexOutOfBoundError(p)
	}

	return curr, nil
}

//TODO: RELOCATE CHECK
type check string

const (
	index   check = "index"
	typeURL check = "typeURL"
	empty   check = "empty"
)

var cm = map[check]func(ll *LinkedList, e interface{}) error{
	index: func(ll *LinkedList, e interface{}) error {
		t, ok := e.(int)
		if !ok {
			return liberr.InvalidOperationError(fmt.Sprintf("%v is not an integer", e))
		}

		return ll.isValidIndex(t)
	},

	typeURL: func(ll *LinkedList, e interface{}) error {
		return ll.isValidType(e)
	},

	empty: func(ll *LinkedList, e interface{}) error {
		if ll.Size() == 0 {
			return liberr.EmptyListError
		}

		return nil
	},
}

func (ll *LinkedList) isValidIndex(i int) error {
	if i < 0 || i >= ll.Size() {
		return liberr.IndexOutOfBoundError(i)
	}

	return nil
}

func (ll *LinkedList) isValidType(e interface{}) error {
	if et := utils.GetTypeName(e); et != ll.typeURL {
		return liberr.TypeMismatchError(ll.typeURL, et)
	}

	return nil
}

type checkData struct {
	check check
	data  interface{}
}

func checks(ll *LinkedList, cd []checkData) error {
	for _, v := range cd {
		if err := cm[v.check](ll, v.data); err != nil {
			return err
		}
	}

	return nil
}
