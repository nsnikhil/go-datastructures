package list

import (
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/functions/operator"
	"github.com/nsnikhil/go-datastructures/functions/predicate"
	"github.com/nsnikhil/go-datastructures/internal"
)

type node[T comparable] struct {
	element T
	next    *node[T]
	prev    *node[T]
}

func newNode[T comparable](element T) *node[T] {
	return &node[T]{
		element: element,
	}
}

type LinkedList[T comparable] struct {
	size  int64
	first *node[T]
	last  *node[T]
}

func NewLinkedList[T comparable](elements ...T) *LinkedList[T] {
	ll := &LinkedList[T]{size: nought}

	ll.AddAll(elements...)

	return ll
}

func (ll *LinkedList[T]) Add(element T) {
	ll.addAt(element, ll.size)
}

func (ll *LinkedList[T]) AddAt(index int64, element T) error {
	return ll.addAt(element, index)
}

//TODO: SHOULD IT FAIL WHEN ARGS ARE EMPTY
func (ll *LinkedList[T]) AddAll(elements ...T) {
	lsz := ll.size

	for i := int64(0); i < int64(len(elements)); i++ {
		ll.addAt(elements[i], i+lsz)
	}

	return
}

func (ll *LinkedList[T]) AddFirst(element T) {
	ll.addAt(element, 0)
}

func (ll *LinkedList[T]) AddLast(element T) {
	ll.addAt(element, ll.size)
}

func (ll *LinkedList[T]) Clear() {
	ll.first = nil
	ll.last = nil
	ll.size = 0
}

func (ll *LinkedList[T]) Clone() List[T] {
	if ll.IsEmpty() {
		return NewLinkedList[T]()
	}

	e, _ := ll.SubList(0, ll.size-1)
	return e
}

func (ll *LinkedList[T]) Contains(element T) bool {
	return newDoublyFinder[T]().search(ll, element) != -1
}

func (ll *LinkedList[T]) ContainsAll(elements ...T) bool {
	if ll.IsEmpty() {
		return false
	}

	elementSet := make(map[T]bool)

	curr := ll.first

	for curr != nil {
		elementSet[curr.element] = true
		curr = curr.next
	}

	for _, ele := range elements {
		if !elementSet[ele] {
			return false
		}
	}

	return true
}

func (ll *LinkedList[T]) Get(index int64) (T, error) {
	if ll.IsEmpty() {
		return internal.ZeroValueOf[T](), emptyListError("LinkedList.Get")
	}

	curr, err := ll.traverseTo(index)
	if err != nil {
		return internal.ZeroValueOf[T](), invalidIndexError(index, "LinkedList.Get")
	}

	return curr.element, nil
}

func (ll *LinkedList[T]) GetFirst() (T, error) {
	if ll.IsEmpty() {
		return internal.ZeroValueOf[T](), emptyListError("LinkedList.GetFirst")
	}

	return ll.first.element, nil
}

func (ll *LinkedList[T]) GetLast() (T, error) {
	if ll.IsEmpty() {
		return internal.ZeroValueOf[T](), emptyListError("LinkedList.GetLast")
	}

	return ll.last.element, nil
}

func (ll *LinkedList[T]) Filter(predicate predicate.Predicate[T]) List[T] {
	curr := ll.first
	var elements []T

	for curr != nil {
		if ok := predicate.Test(curr.element); ok {
			elements = append(elements, curr.element)
		}
		curr = curr.next
	}

	return NewLinkedList[T](elements...)
}

//TODO: ADD TEST
func (ll *LinkedList[T]) FindFirst(predicate predicate.Predicate[T]) (T, error) {
	if ll.IsEmpty() {
		return internal.ZeroValueOf[T](), emptyListError("LinkedList.FindFirst")
	}

	curr := ll.first

	for curr != nil {
		if ok := predicate.Test(curr.element); ok {
			return curr.element, nil
		}
		curr = curr.next
	}

	return internal.ZeroValueOf[T](), noElementMatchFilterError("LinkedList.FindFirst")
}

func (ll *LinkedList[T]) IndexOf(element T) int64 {
	if ll.IsEmpty() {
		return invalidIndex
	}

	return newLinearFinder[T]().search(ll, element)
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}

func (ll *LinkedList[T]) Iterator() iterator.Iterator[T] {
	return newLinkedListIterator(ll)
}

func (ll *LinkedList[T]) DescendingIterator() iterator.Iterator[T] {
	return newLinkedListDescendingIterator(ll)
}

//TODO: OPTIMIZE, MAKE IT SHORT
func (ll *LinkedList[T]) LastIndexOf(element T) int64 {
	if ll.IsEmpty() {
		return invalidIndex
	}

	return newReverseFinder[T]().search(ll, element)
}

func (ll *LinkedList[T]) Remove(element T) error {
	if ll.IsEmpty() {
		return emptyListError("LinkedList.Remove")
	}

	i := ll.IndexOf(element)
	if i == -1 {
		return elementNotFoundError(element, "LinkedList.Remove")
	}

	if _, err := ll.removeAt(i); err != nil {
		return err
	}

	return nil
}

func (ll *LinkedList[T]) RemoveAt(index int64) (T, error) {
	element, err := ll.removeAt(index)
	if err != nil {
		return internal.ZeroValueOf[T](), err
	}

	return element, nil
}

//TODO: SHOULD IT FAIL OR NOT WHEN ARG LIST IS EMPTY
func (ll *LinkedList[T]) RemoveAll(elements ...T) error {
	return filterLinkedList(ll, false, elements...)
}

func (ll *LinkedList[T]) RemoveFirst() (T, error) {
	return ll.removeAt(0)
}

func (ll *LinkedList[T]) RemoveFirstOccurrence(element T) error {
	return ll.Remove(element)
}

func (ll *LinkedList[T]) RemoveLast() (T, error) {
	return ll.removeAt(ll.Size() - 1)
}

func (ll *LinkedList[T]) RemoveLastOccurrence(element T) (bool, error) {
	if ll.IsEmpty() {
		return false, emptyListError("LinkedList.RemoveLastOccurrence")
	}

	index := ll.LastIndexOf(element)
	if index == invalidIndex {
		return false, elementNotFoundError(element, "LinkedList.RemoveLastOccurrence")
	}

	if _, err := ll.removeAt(index); err != nil {
		return false, err
	}

	return true, nil
}

func (ll *LinkedList[T]) Replace(old, new T) error {
	if ll.IsEmpty() {
		return emptyListError("LinkedList.Replace")
	}

	curr := ll.first
	for curr != nil && curr.element != old {
		curr = curr.next
	}

	if curr == nil {
		return elementNotFoundError(old, "LinkedList.Replace")
	}

	curr.element = new
	return nil
}

func (ll *LinkedList[T]) ReplaceAll(uo operator.UnaryOperator[T]) {
	curr := ll.first

	for curr != nil {
		curr.element = uo.Apply(curr.element)
		curr = curr.next
	}
}

func (ll *LinkedList[T]) RetainAll(elements ...T) error {
	return filterLinkedList(ll, true, elements...)
}

func (ll *LinkedList[T]) Set(index int64, element T) (T, error) {
	if ll.IsEmpty() {
		return internal.ZeroValueOf[T](), emptyListError("LinkedList.Set")
	}

	if !ll.isValidIndex(index) {
		return internal.ZeroValueOf[T](), invalidIndexError(index, "LinkedList.Set")
	}

	curr, err := ll.traverseTo(index)
	if err != nil {
		return internal.ZeroValueOf[T](), err
	}

	curr.element = element
	return curr.element, nil
}

func (ll *LinkedList[T]) Size() int64 {
	return ll.size
}

//TODO: CHANGE TO LINKED LIST MERGE SORT
func (ll *LinkedList[T]) Sort(c comparator.Comparator[T]) {
	al := ll.ToArrayList()
	al.Sort(c)

	it := al.Iterator()
	temp := ll.first

	for temp != nil && it.HasNext() {
		temp.element, _ = it.Next()
		temp = temp.next
	}
}

//TODO: OPTIMIZE
func (ll *LinkedList[T]) SubList(start int64, end int64) (List[T], error) {
	if end < start {
		return nil, invalidArgsError("end cannot be smaller than start", "LinkedList.SubList")
	}

	if !ll.isValidIndex(start) {
		return nil, invalidIndexError(start, "LinkedList.SubList")
	}

	//TODO USE IS VALID INDEX METHOD HERE
	if end < 0 || end > ll.Size() {
		return nil, invalidIndexError(end, "LinkedList.SubList")
	}

	var tempLL *LinkedList[T] = NewLinkedList[T]()

	temp, _ := ll.traverseTo(start)

	n := start

	for temp != nil && n <= end {
		tempLL.addAt(temp.element, tempLL.size)
		temp = temp.next
		n++
	}

	return tempLL, nil
}

func filterLinkedList[T comparable](ll *LinkedList[T], inverse bool, elements ...T) error {
	if ll.IsEmpty() {
		return emptyListError("LinkedList.filterLinkedList")
	}

	elementSet := make(map[T]bool)
	for _, e := range elements {
		elementSet[e] = true
	}

	curr := ll.first

	rc := int64(0)

	for curr != nil {

		var shouldRemove bool

		if inverse {
			shouldRemove = !elementSet[curr.element]
		} else {
			shouldRemove = elementSet[curr.element]
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

	ll.size -= rc
	return nil
}

type linkedListIterator[T comparable] struct {
	currNode *node[T]
	ll       *LinkedList[T]
}

func newLinkedListIterator[T comparable](ll *LinkedList[T]) iterator.Iterator[T] {
	return &linkedListIterator[T]{
		currNode: ll.first,
		ll:       ll,
	}
}

func (ll *linkedListIterator[T]) HasNext() bool {
	return ll.currNode != nil
}

func (ll *linkedListIterator[T]) Next() (T, error) {
	if ll.currNode == nil {
		return internal.ZeroValueOf[T](), emptyIteratorError("linkedListIterator.Next")
	}

	e := ll.currNode.element
	ll.currNode = ll.currNode.next

	return e, nil
}

type linkedListDescendingIterator[T comparable] struct {
	currNode *node[T]
	ll       *LinkedList[T]
}

func newLinkedListDescendingIterator[T comparable](ll *LinkedList[T]) iterator.Iterator[T] {
	return &linkedListDescendingIterator[T]{
		currNode: ll.last,
		ll:       ll,
	}
}

func (ll *linkedListDescendingIterator[T]) HasNext() bool {
	return ll.currNode != nil
}

func (ll *linkedListDescendingIterator[T]) Next() (T, error) {
	if ll.currNode == nil {
		return internal.ZeroValueOf[T](), emptyIteratorError("linkedListDescendingIterator.Next")
	}

	e := ll.currNode.element
	ll.currNode = ll.currNode.prev

	return e, nil
}

func (ll *LinkedList[T]) ToArrayList() *ArrayList[T] {
	var e []T
	it := ll.Iterator()

	for it.HasNext() {
		v, _ := it.Next()
		e = append(e, v)
	}

	al := NewArrayList[T](e...)
	return al
}

// HELPER FUNCTION FROM HERE ON
func (ll *LinkedList[T]) addAt(element T, index int64) error {
	sz := ll.size

	if index < 0 || index > sz {
		return invalidIndexError(index, "LinkedList.addAt")
	}

	nd := newNode(element)

	if ll.IsEmpty() {
		ll.first = nd
		ll.last = nd
		ll.size++
		return nil
	}

	if index == 0 {
		nd.next = ll.first
		ll.first.prev = nd
		ll.first = nd
		ll.size++
		return nil
	}

	if index == sz {
		ll.last.next = nd
		nd.prev = ll.last
		ll.last = nd
		ll.size++
		return nil
	}

	curr, err := ll.traverseTo(index)
	if err != nil {
		return err
	}

	cp := curr.prev

	nd.next = curr
	curr.prev = nd

	cp.next = nd
	nd.prev = cp

	ll.size++

	return nil
}

func (ll *LinkedList[T]) removeAt(index int64) (T, error) {
	sz := ll.size

	if sz == 0 {
		return internal.ZeroValueOf[T](), emptyListError("LinkedList.removeAt")
	}

	if index < 0 || index >= sz {
		return internal.ZeroValueOf[T](), invalidIndexError(index, "LinkedList.removeAt")
	}

	if sz == 1 {
		element := ll.first.element
		ll.Clear()
		return element, nil
	}

	if index == 0 {
		element := ll.first.element
		ll.first = ll.first.next
		ll.first.prev = nil
		ll.size--
		return element, nil
	}

	if index == sz-1 {
		element := ll.last.element
		ll.last = ll.last.prev
		ll.last.next = nil
		ll.size--
		return element, nil
	}

	curr, err := ll.traverseTo(index)
	if err != nil {
		return internal.ZeroValueOf[T](), err
	}

	cp := curr.prev
	cn := curr.next

	cp.next = cn
	cn.prev = cp
	ll.size--

	return curr.element, nil
}

func (ll *LinkedList[T]) traverseTo(index int64) (*node[T], error) {
	traverseToFromFirst := func(index int64) *node[T] {
		curr := ll.first
		k := index

		for curr != nil && k > 0 {
			curr = curr.next
			k--
		}

		return curr
	}

	traverseToFromLast := func(sz, index int64) *node[T] {
		curr := ll.last
		k := sz - index - 1

		for curr != nil && k > 0 {
			curr = curr.prev
			k--
		}

		return curr
	}

	sz := ll.Size()

	if index < 0 || index >= sz {
		return nil, invalidIndexError(index, "LinkedList.traverseTo")
	}

	if index == 0 {
		return ll.first, nil
	}

	if index == sz-1 {
		return ll.last, nil
	}

	md := sz / 2
	df := sz - index

	var curr *node[T]

	if df > md {
		curr = traverseToFromFirst(index)
	} else {
		curr = traverseToFromLast(sz, index)
	}

	if curr == nil {
		return nil, invalidIndexError(index, "LinkedList.traverseTo")
	}

	return curr, nil
}

func (ll *LinkedList[T]) isValidIndex(i int64) bool {
	return i >= 0 && i < ll.size
}
