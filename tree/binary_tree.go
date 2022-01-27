package tree

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/liberr"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/nsnikhil/go-datastructures/queue"
	"github.com/nsnikhil/go-datastructures/stack"
	"github.com/nsnikhil/go-datastructures/utils"
)

type binaryNode[T comparable] struct {
	data T
	//level  int // NOT IMPLEMENTED
	hd     int
	left   *binaryNode[T]
	right  *binaryNode[T]
	parent *binaryNode[T]
}

func (bn *binaryNode[T]) String() string {
	return fmt.Sprint(bn.data)
}

func (bn *binaryNode[T]) isLeaf() bool {
	return bn.left == nil && bn.right == nil
}

func (bn *binaryNode[T]) childCount() int {
	c := 0
	if bn.left != nil {
		c++
	}
	if bn.right != nil {
		c++
	}
	return c
}

func (bn *binaryNode[T]) clone() *binaryNode[T] {
	return &binaryNode[T]{
		data:   bn.data,
		hd:     bn.hd,
		left:   bn.left,
		right:  bn.right,
		parent: bn.parent,
	}
}

func (bn *binaryNode[T]) detach() {
	p := bn.parent

	if p == nil {
		return
	}

	if p.left == bn {
		p.left = nil
	} else if p.right == bn {
		p.right = nil
	}

	bn.parent = nil
}

func newBinaryNode[T comparable](data T) *binaryNode[T] {
	return &binaryNode[T]{data: data}
}

type BinaryTree[T comparable] struct {
	typeURL string
	count   int
	root    *binaryNode[T]
}

func NewBinaryTree[T comparable](e ...T) (*BinaryTree[T], error) {
	bt := &BinaryTree[T]{
		typeURL: utils.NA,
		count:   utils.Naught,
	}

	sz := int64(len(e))

	if sz == 0 {
		return bt, nil
	}

	for i := int64(0); i < sz; i++ {
		if err := bt.Insert(e[i]); err != nil {
			return nil, err
		}
	}

	return bt, nil
}

func (bt *BinaryTree[T]) Insert(e T) error {
	et := utils.GetTypeName(e)

	if bt.typeURL == utils.NA {
		bt.typeURL = et
	}

	if bt.typeURL != et {
		return liberr.TypeMismatchError(bt.typeURL, et)
	}

	t := newBinaryNode(e)

	if bt.root == nil {
		bt.root = t
		bt.count++
		return nil
	}

	curr := bt.root

	q := queue.NewLinkedQueue[*binaryNode[T]]()

	if err := q.Add(curr); err != nil {
		return err
	}

	err := func() error {
		for !q.Empty() {
			sz := q.Size()

			for i := int64(0); i < sz; i++ {
				f, _ := q.Remove()

				if f.left == nil {
					f.left = t
					t.parent = f
					return nil
				}

				if f.right == nil {
					f.right = t
					t.parent = f
					return nil
				}

				if err := q.Add(f.left); err != nil {
					return err
				}

				if err := q.Add(f.right); err != nil {
					return err
				}
			}

		}

		return nil
	}()

	if err != nil {
		return err
	}

	bt.count++
	return nil
}

func (bt *BinaryTree[T]) InsertCompare(e T, c comparator.Comparator[T]) error {
	et := utils.GetTypeName(e)

	if bt.typeURL == utils.NA {
		bt.typeURL = et
	}

	if bt.typeURL != et {
		return liberr.TypeMismatchError(bt.typeURL, et)
	}

	t := newBinaryNode(e)

	if bt.root == nil {
		bt.root = t
		bt.count++
		return nil
	}

	curr := bt.root

	for {
		i := c.Compare(curr.data, t.data)

		if i > 0 {
			if curr.left != nil {
				curr = curr.left
			} else {
				break
			}
		} else {
			if curr.right != nil {
				curr = curr.right
			} else {
				break
			}
		}

	}

	i := c.Compare(curr.data, t.data)

	if i > 0 {
		curr.left = t
		t.parent = curr
	} else {
		curr.right = t
		t.parent = curr
	}

	bt.count++
	return nil
}

func (bt *BinaryTree[T]) Delete(e T) error {
	if bt.Empty() {
		return errors.New("tree is empty")
	}

	n, err := search(e, bt.root)
	if err != nil {
		return err
	}

	l, p := lastNode(bt)

	if l == n && p == nil {
		bt.Clear()
		return nil
	}

	if l != n {
		n.data = l.data
	}

	if p.left == l {
		l.parent = nil
		p.left = nil
	} else {
		l.parent = nil
		p.right = nil
	}

	bt.count--
	return nil
}

func (bt *BinaryTree[T]) DeleteCompare(e T, c comparator.Comparator[T]) error {
	var deleteNode func(e T, c comparator.Comparator[T], n *binaryNode[T]) error

	deleteNode = func(e T, c comparator.Comparator[T], n *binaryNode[T]) error {
		if n == nil {
			return fmt.Errorf("%v not found in the tree", e)
		}

		i := c.Compare(n.data, e)

		if i > 0 {
			return deleteNode(e, c, n.left)
		} else if i < 0 {
			return deleteNode(e, c, n.right)
		} else {

			if n.isLeaf() {
				n.detach()
			} else if n.left == nil {
				n.data = n.right.data
				n.right.detach()
			} else if n.right == nil {
				n.data = n.left.data
				n.left.detach()
			} else {
				sn := inOrderSuccessor(n.right)
				n.data = sn.data
				sn.detach()
			}
		}

		return nil
	}

	if bt.Empty() {
		return errors.New("tree is empty")
	}

	curr := bt.root
	if curr.data == e && curr.isLeaf() {
		bt.Clear()
		return nil
	}

	err := deleteNode(e, c, curr)
	if err != nil {
		return err
	}

	bt.count--
	return nil
}

func (bt *BinaryTree[T]) Search(e T) (bool, error) {
	if bt.Empty() {
		return false, errors.New("tree is empty")
	}

	et := utils.GetTypeName(e)
	if bt.typeURL != et {
		return false, liberr.TypeMismatchError(bt.typeURL, et)
	}

	_, err := search(e, bt.root)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (bt *BinaryTree[T]) SearchCompare(e T, c comparator.Comparator[T]) (bool, error) {
	if bt.Empty() {
		return false, errors.New("tree is empty")
	}

	et := utils.GetTypeName(e)
	if et != bt.typeURL {
		return false, liberr.TypeMismatchError(bt.typeURL, et)
	}

	curr := bt.root

	for curr != nil {
		if curr.data == e {
			return true, nil
		}

		res := c.Compare(curr.data, e)

		if res > 0 {
			curr = curr.left
		} else if res < 0 {
			curr = curr.right
		}
	}

	return false, fmt.Errorf("%v not found in the tree", e)
}

func (bt *BinaryTree[T]) Count() int {
	return bt.count
}

func (bt *BinaryTree[T]) Height() int {
	return calculateHeight(bt.root, nil)
}

func (bt *BinaryTree[T]) Diameter() int {
	d := utils.Naught
	calculateHeight(bt.root, &d)
	return d
}

func (bt *BinaryTree[T]) Empty() bool {
	return bt.count == utils.Naught
}

func (bt *BinaryTree[T]) Clear() {
	bt.root = nil
	bt.count = utils.Naught
}

func (bt *BinaryTree[T]) Clone() Tree[T] {
	return &BinaryTree[T]{
		typeURL: bt.typeURL,
		count:   bt.count,
		root:    cloneNodes(bt.root, nil),
	}
}

func (bt *BinaryTree[T]) Mirror() (bool, error) {
	if bt.Empty() {
		return false, errors.New("tree is empty")
	}

	return bt.MirrorAt(bt.root.data)
}

func (bt *BinaryTree[T]) MirrorAt(e T) (bool, error) {
	if bt.Empty() {
		return false, errors.New("tree is empty")
	}

	et := utils.GetTypeName(e)
	if bt.typeURL != et {
		return false, liberr.TypeMismatchError(bt.typeURL, et)
	}

	curr, err := search(e, bt.root)
	if err != nil {
		return false, err
	}

	if err := mirrorAt(curr); err != nil {
		return false, err
	}

	return true, nil
}

func (bt *BinaryTree[T]) RotateLeft() error {
	if bt.Empty() {
		return errors.New("tree is empty")
	}

	rotateLeft(bt.root, nil, bt)
	return nil
}

func (bt *BinaryTree[T]) RotateRight() error {
	if bt.Empty() {
		return errors.New("tree is empty")
	}

	rotateRight(bt.root, nil, bt)
	return nil
}

func (bt *BinaryTree[T]) RotateLeftAt(e T) error {
	if bt.Empty() {
		return errors.New("tree is empty")
	}

	et := utils.GetTypeName(e)
	if bt.typeURL != et {
		return liberr.TypeMismatchError(bt.typeURL, et)
	}

	curr, err := search(e, bt.root)
	if err != nil {
		return err
	}

	rotateLeft(curr, curr.parent, bt)
	return nil
}

func (bt *BinaryTree[T]) RotateRightAt(e T) error {
	if bt.Empty() {
		return errors.New("tree is empty")
	}

	et := utils.GetTypeName(e)
	if bt.typeURL != et {
		return liberr.TypeMismatchError(bt.typeURL, et)
	}

	curr, err := search(e, bt.root)
	if err != nil {
		return err
	}

	rotateRight(curr, curr.parent, bt)
	return nil
}

func (bt *BinaryTree[T]) IsBalanced() bool {
	return isBalancedAt(bt.root)
}

func (bt *BinaryTree[T]) IsFull() bool {
	it := newBtLvOrderIterator(bt)

	it.(*btLvOrderIterator[T]).v = true

	for it.HasNext() {
		v, _ := it.Next()
		if v.childCount() == 1 {
			return false
		}
	}

	return true
}

func (bt *BinaryTree[T]) IsPerfect() bool {
	curr := bt.root
	q := queue.NewLinkedQueue[*binaryNode[T]]()
	_ = q.Add(curr)
	c := 0
	fl := -1

	for !q.Empty() {
		sz := q.Size()

		for i := int64(0); i < sz; i++ {
			e, _ := q.Remove()

			if e.isLeaf() {
				if fl == -1 {
					fl = c
				}
				if fl != c {
					return false
				}
			}

			if e.left != nil {
				_ = q.Add(e.left)
			}

			if e.right != nil {
				_ = q.Add(e.right)
			}
		}
		c++
	}

	return true
}

func (bt *BinaryTree[T]) IsComplete() bool {
	h := 0
	res := true
	isComplete(bt.root, &h, &res)
	return res
}

func (bt *BinaryTree[T]) LowestCommonAncestor(a, b T) (T, error) {
	an, err := search(a, bt.root)
	if err != nil {
		return *new(T), err
	}

	bn, err := search(b, bt.root)
	if err != nil {
		return *new(T), err
	}

	return lowestCommonAncestor(an, bn, bt.root).data, nil
}

func (bt *BinaryTree[T]) Paths() (list.List[T], error) {
	if bt.Empty() {
		return nil, errors.New("tree is empty")
	}

	temp := list.NewLinkedList[T]()

	res := list.NewArrayList[T]()

	if err := paths(bt.root, temp, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (bt *BinaryTree[T]) Mode() (list.List[T], error) {
	if bt.Empty() {
		return nil, errors.New("tree is empty")
	}

	res := list.NewArrayList[T]()

	cm := 0
	m := 0
	var p T

	var mode func(n *binaryNode[T], l list.List[T], cm, m *int, p *T) error
	mode = func(n *binaryNode[T], l list.List[T], cm, m *int, p *T) error {
		if n == nil {
			return nil
		}

		if err := mode(n.left, l, cm, m, p); err != nil {
			return err
		}

		if n.data == p {
			*cm += 1
		} else {
			*cm = 1
		}

		if *cm > *m {
			l.Clear()
		}

		if *cm >= *m {
			if err := l.Add(n.data); err != nil {
				return err
			}
		}

		*p = n.data
		*m = max(*m, *cm)

		if err := mode(n.right, l, cm, m, p); err != nil {
			return err
		}

		return nil
	}

	if err := mode(bt.root, res, &cm, &m, &p); err != nil {
		return nil, err
	}

	return res, nil
}

func (bt *BinaryTree[T]) Equal(t Tree[T]) (bool, error) {
	_, ok := t.(*BinaryTree[T])
	if !ok {
		return false, errors.New("TODO")
	}

	var equal func(a, b *binaryNode[T]) bool

	equal = func(a, b *binaryNode[T]) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil || a.data != b.data {
			return false
		}
		return equal(a.left, b.left) && equal(a.right, b.right)
	}

	return equal(bt.root, t.root), nil
}

func (bt *BinaryTree[T]) Symmetric() bool {
	var symmetric func(a, b *binaryNode[T]) bool

	symmetric = func(a, b *binaryNode[T]) bool {
		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil || a.data != b.data {
			return false
		}

		return symmetric(a.left, b.right) && symmetric(a.right, b.left)
	}

	return symmetric(bt.root.left, bt.root.right)
}

func (bt *BinaryTree[T]) Invert() {
	var inverter func(n *binaryNode[T]) *binaryNode

	inverter = func(n *binaryNode[T]) *binaryNode {
		if n == nil {
			return nil
		}

		l := inverter(n.right)
		r := inverter(n.left)
		n.right = l
		n.left = r

		return n
	}

	bt.root = inverter(bt.root)
}

func (bt *BinaryTree[T]) PreOrderSuccessor(e T) (T, error) {
	st, err := stack.NewStack()
	if err != nil {
		return nil, err
	}
	curr := bt.root
	var prev *binaryNode

	for !st.Empty() || curr != nil {

		for curr != nil {

			if prev != nil && prev.data == e {
				return curr.data, nil
			}

			if curr.right != nil {
				_ = st.Push(curr.right)
			}

			prev = curr
			curr = curr.left

		}

		if !st.Empty() {
			top, _ := st.Pop()
			curr = top.(*binaryNode[T])
		}

	}

	if prev != nil && prev.data == e {
		return nil, fmt.Errorf("no pre order successor found for %v", e)
	}

	return nil, fmt.Errorf("%v not found in the tree", e)
}

func (bt *BinaryTree[T]) PostOrderSuccessor(e T) (T, error) {
	st, err := stack.NewStack()
	if err != nil {
		return nil, err
	}

	curr := bt.root
	var prev *binaryNode

	for !st.Empty() || curr != nil {

		for curr != nil {
			_ = st.Push(curr)
			curr = curr.left
		}

		if !st.Empty() {
			top, _ := st.Peek()
			curr = top.(*binaryNode[T])
		}

		if curr == nil {
			break
		}

		if curr.right != nil && curr.right != prev {
			curr = curr.right
		} else {

			if prev != nil && prev.data == e {
				return curr.data, nil
			}

			_, _ = st.Pop()
			prev = curr
			curr = nil
		}

	}

	if prev != nil && prev.data == e {
		return nil, fmt.Errorf("no post order successor found for %v", e)
	}

	return nil, fmt.Errorf("%v not found in the tree", e)
}

func (bt *BinaryTree[T]) InOrderSuccessor(e T) (T, error) {
	st, err := stack.NewStack()
	if err != nil {
		return nil, err
	}

	curr := bt.root
	var prev *binaryNode

	for !st.Empty() || curr != nil {

		for curr != nil {
			_ = st.Push(curr)
			curr = curr.left
		}

		if !st.Empty() {
			top, _ := st.Pop()
			curr = top.(*binaryNode[T])
		}

		if curr == nil {
			break
		}

		if prev != nil && prev.data == e {
			return curr.data, nil
		}

		prev = curr
		curr = curr.right
	}

	if prev != nil && prev.data == e {
		return nil, fmt.Errorf("no in order successor found for %v", e)
	}

	return nil, fmt.Errorf("%v not found in the tree", e)
}

func (bt *BinaryTree[T]) LevelOrderSuccessor(e T) (T, error) {
	q, err := queue.NewLinkedQueue()
	if err != nil {
		return nil, err
	}

	var prev *binaryNode[T]
	_ = q.Add(bt.root)

	for !q.Empty() {

		sz := q.Size()

		for i := int64(0); i < sz; i++ {
			f, _ := q.Remove()

			if prev != nil && prev.data == e {
				return f.data, nil
			}

			prev = f

			if f.left != nil {
				_ = q.Add(f.left)
			}

			if f.right != nil {
				_ = q.Add(f.right)
			}
		}
	}

	if prev != nil && prev.data == e {
		return nil, fmt.Errorf("no level order successor found for %v", e)
	}

	return nil, fmt.Errorf("%v not found in the tree", e)
}

func (bt *BinaryTree[T]) PreOrderIterator() iterator.Iterator[T] {
	return newBtPreOrderIterator[T](bt)
}

type btPreOrderIterator[T comparable] struct {
	curr *binaryNode[T]
	s    *stack.Stack[T]
	v    bool
}

func newBtPreOrderIterator[T comparable](bt *BinaryTree[T]) *btPreOrderIterator[T] {
	s := stack.NewStack()
	return &btPreOrderIterator[T]{
		curr: bt.root,
		s:    s,
	}
}

func (bti *btPreOrderIterator[T]) HasNext() bool {
	return bti.curr != nil || !bti.s.Empty()
}

func (bti *btPreOrderIterator[T]) Next() (T, error) {
	if bti.curr == nil {
		n, _ := bti.s.Pop()
		bti.curr = n.(*binaryNode[T])
	}

	temp := bti.curr

	if bti.curr.right != nil {
		_ = bti.s.Push(bti.curr.right)
	}

	bti.curr = bti.curr.left

	if bti.v {
		return temp
	}

	return temp.data
}

func (bt *BinaryTree[T]) PostOrderIterator() iterator.Iterator[T] {
	return newBtPostOrderIterator(bt)
}

type btPostOrderIterator[T comparable] struct {
	curr *binaryNode[T]
	last *binaryNode[T]
	s    *stack.Stack[T]
	v    bool
}

func newBtPostOrderIterator[T comparable](bt *BinaryTree[T]) *btPostOrderIterator[T] {
	s, _ := stack.NewStack()
	return &btPostOrderIterator{
		curr: bt.root,
		s:    s,
	}
}

func (bto *btPostOrderIterator[T]) HasNext() bool {
	return bto.curr != nil || !bto.s.Empty()
}

func (bto *btPostOrderIterator[T]) Next() (T, error) {
	get := func() T {
		_, _ = bto.s.Pop()

		temp := bto.curr
		bto.curr = nil

		bto.last = temp

		if bto.v {
			return temp
		}

		return temp.data
	}

	if bto.curr == nil {
		if bto.s.Empty() {
			return nil
		}

		top, _ := bto.s.Peek()

		bto.curr = top.(*binaryNode[T])

		if bto.curr.right != nil && bto.curr.right != bto.last {
			bto.curr = bto.curr.right
		} else {
			return get()
		}
	}

	left := func() {
		for bto.curr != nil {
			_ = bto.s.Push(bto.curr)
			bto.curr = bto.curr.left
		}

		if !bto.s.Empty() {
			top, _ := bto.s.Peek()
			bto.curr = top.(*binaryNode[T])
		}
	}

	left()

	if bto.curr == nil {
		return nil
	}

	for bto.curr != nil && bto.curr.right != nil && bto.curr.right != bto.last {
		bto.curr = bto.curr.right
		left()
	}

	return get()
}

func (bt *BinaryTree[T]) InOrderIterator() iterator.Iterator[T] {
	return newBtInOrderIterator(bt)
}

type btInOrderIterator[T comparable] struct {
	curr *binaryNode[T]
	s    *stack.Stack[T]
	v    bool
}

func newBtInOrderIterator[T comparable](bt *BinaryTree[T]) *btInOrderIterator[T] {
	s, _ := stack.NewStack()
	return &btInOrderIterator{
		curr: bt.root,
		s:    s,
	}
}

func (bti *btInOrderIterator[T]) HasNext() bool {
	return bti.curr != nil || !bti.s.Empty()
}

func (bti *btInOrderIterator[T]) Next() (T, error) {
	for bti.curr != nil {
		_ = bti.s.Push(bti.curr)
		bti.curr = bti.curr.left
	}

	if !bti.s.Empty() {
		top, _ := bti.s.Pop()
		bti.curr = top.(*binaryNode[T])
	}

	if bti.curr == nil {
		return nil
	}

	temp := bti.curr

	bti.curr = bti.curr.right

	if bti.v {
		return temp
	}

	return temp.data
}

func (bt *BinaryTree[T]) LevelOrderIterator() iterator.Iterator[T] {
	return newBtLvOrderIterator(bt)
}

type btLvOrderIterator[T comparable] struct {
	curr *binaryNode[T]
	q    queue.Queue[T]
	v    bool
}

func newBtLvOrderIterator[T comparable](bt *BinaryTree[T]) iterator.Iterator[T] {
	q, _ := queue.NewLinkedQueue()
	_ = q.Add(bt.root)

	return &btLvOrderIterator{
		curr: bt.root,
		q:    q,
	}
}

func (blv *btLvOrderIterator[T]) HasNext() bool {
	return !blv.q.Empty()
}

func (blv *btLvOrderIterator[T]) Next() (T, error) {
	curr, _ := blv.q.Remove()

	if curr.(*binaryNode[T]).left != nil {
		_ = blv.q.Add(curr.(*binaryNode[T]).left)
	}

	if curr.(*binaryNode[T]).right != nil {
		_ = blv.q.Add(curr.(*binaryNode[T]).right)
	}

	if blv.v {
		return curr
	}

	return curr.(*binaryNode[T]).data
}

//TODO FIX EXPENSIVE IMPLEMENTATION
func (bt *BinaryTree[T]) VerticalViewIterator() iterator.Iterator[T] {
	return newBtVerticalVOrderIterator(bt)
}

type btVerticalVOrderIterator[T comparable] struct {
	it iterator.Iterator[T]
	v  bool
}

func newBtVerticalVOrderIterator[T comparable](bt *BinaryTree[T]) iterator.Iterator[T] {
	return &btVerticalVOrderIterator{
		it: horizontalIterator(bt, 2),
	}
}

func (btv *btVerticalVOrderIterator[T]) HasNext() bool {
	return btv.it.HasNext()
}

func (btv *btVerticalVOrderIterator[T]) Next() (T, error) {
	if btv.v {
		return btv.it.Next()
	}

	return btv.it.Next().(*binaryNode[T]).data
}

func (bt *BinaryTree[T]) LeftViewIterator() iterator.Iterator[T] {
	return newBtLeftVOrderIterator(bt)
}

type btLeftVOrderIterator[T comparable] struct {
	curr *binaryNode[T]
	q    queue.Queue[T]
	v    bool
}

func newBtLeftVOrderIterator[T comparable](bt *BinaryTree[T]) iterator.Iterator[T] {
	q, _ := queue.NewLinkedQueue()
	_ = q.Add(bt.root)

	return &btLeftVOrderIterator{
		curr: bt.root,
		q:    q,
	}
}

func (bfv *btLeftVOrderIterator[T]) HasNext() bool {
	return !bfv.q.Empty()
}

func (bfv *btLeftVOrderIterator[T]) Next() (T, error) {
	sz := bfv.q.Size()

	var res *binaryNode = nil

	for i := int64(0); i < sz; i++ {
		curr, _ := bfv.q.Remove()

		if res == nil {
			res = curr.(*binaryNode[T])
		}

		if curr.(*binaryNode[T]).left != nil {
			_ = bfv.q.Add(curr.(*binaryNode[T]).left)
		}

		if curr.(*binaryNode[T]).right != nil {
			_ = bfv.q.Add(curr.(*binaryNode[T]).right)
		}

	}

	if bfv.v {
		return res
	}

	return res.data
}

func (bt *BinaryTree[T]) RightViewIterator() iterator.Iterator[T] {
	return newBtRightVOrderIterator(bt)
}

type btRightVOrderIterator[T comparable] struct {
	curr *binaryNode[T]
	q    queue.Queue[T]
	v    bool
}

func newBtRightVOrderIterator[T comparable](bt *BinaryTree[T]) iterator.Iterator[T] {
	q, _ := queue.NewLinkedQueue()
	_ = q.Add(bt.root)

	return &btRightVOrderIterator{
		curr: bt.root,
		q:    q,
	}
}

func (brv *btRightVOrderIterator[T]) HasNext() bool {
	return !brv.q.Empty()
}

func (brv *btRightVOrderIterator[T]) Next() (T, error) {
	sz := brv.q.Size()

	var res *binaryNode = nil

	for i := int64(0); i < sz; i++ {
		curr, _ := brv.q.Remove()

		if res == nil {
			res = curr.(*binaryNode[T])
		}

		if curr.(*binaryNode[T]).right != nil {
			_ = brv.q.Add(curr.(*binaryNode[T]).right)
		}

		if curr.(*binaryNode[T]).left != nil {
			_ = brv.q.Add(curr.(*binaryNode[T]).left)
		}
	}

	if brv.v {
		return res
	}

	return res.data
}

//TODO FIX EXPENSIVE IMPLEMENTATION
func (bt *BinaryTree[T]) TopViewIterator() iterator.Iterator[T] {
	return newBtTopVOrderIterator(bt.Clone().(*BinaryTree))
}

type btTopVOrderIterator[T comparable] struct {
	it iterator.Iterator[T]
	v  bool
}

func newBtTopVOrderIterator[T comparable](bt *BinaryTree[T]) iterator.Iterator[T] {
	return &btTopVOrderIterator{
		it: horizontalIterator(bt, 0),
	}
}

func (btv *btTopVOrderIterator[T]) HasNext() bool {
	return btv.it.HasNext()
}

func (btv *btTopVOrderIterator[T]) Next() (T, error) {
	if btv.v {
		return btv.it.Next()
	}

	return btv.it.Next().(*binaryNode[T]).data
}

//TODO FIX EXPENSIVE IMPLEMENTATION
func (bt *BinaryTree[T]) BottomViewIterator() iterator.Iterator[T] {
	return newBtBottomVOrderIterator(bt.Clone().(*BinaryTree))
}

type btBottomVOrderIterator[T comparable] struct {
	it iterator.Iterator[T]
	v  bool
}

func newBtBottomVOrderIterator[T comparable](bt *BinaryTree[T]) iterator.Iterator[T] {
	return &btBottomVOrderIterator{
		it: horizontalIterator(bt, 1),
	}
}

func (brv *btBottomVOrderIterator[T]) HasNext() bool {
	return brv.it.HasNext()
}

func (brv *btBottomVOrderIterator[T]) Next() (T, error) {
	if brv.v {
		return brv.it.Next()
	}

	return brv.it.Next().(*binaryNode[T]).data
}

func lastNode[T comparable](bt *BinaryTree[T]) (*binaryNode[T], *binaryNode[T]) {
	if bt.root == nil {
		return nil, nil
	}

	var prev *binaryNode[T] = nil
	curr := bt.root

	for {
		if curr.right != nil {
			prev = curr
			curr = curr.right
		} else if curr.left != nil {
			prev = curr
			curr = curr.left
		} else {
			break
		}
	}

	return curr, prev
}

//noinspection GoNilness
func mirrorAt[T comparable](n *binaryNode[T]) error {
	curr := n
	st, _ := stack.NewStack()
	var prev *binaryNode

	for curr != nil || !st.Empty() {

		for curr != nil {
			_ = st.Push(curr)
			curr = curr.left
		}

		if !st.Empty() {
			top, _ := st.Peek()
			curr = top.(*binaryNode[T])
		}

		if curr.right != nil && curr.right != prev {
			curr = curr.right
		} else {
			_, _ = st.Pop()
			prev = curr

			if !curr.isLeaf() {
				curr.left, curr.right = curr.right, curr.left
			}

			curr = nil
		}

	}

	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateHeight[T comparable](n *binaryNode[T], diameter *int) int {
	if n == nil {
		return 0
	}

	lh := calculateHeight(n.left, diameter)
	rh := calculateHeight(n.right, diameter)

	if diameter != nil {
		*diameter = max(*diameter, 1+lh+rh)
	}

	return 1 + max(lh, rh)
}

func absDiff(a, b int) int {
	diff := a - b
	if diff < 1 {
		diff *= -1
	}

	return diff
}

func isBalancedAt[T comparable](n *binaryNode[T]) bool {
	if n == nil {
		return true
	}
	return absDiff(calculateHeight(n.left, nil), calculateHeight(n.right, nil)) <= 1 && isBalancedAt(n.left) && isBalancedAt(n.right)
}

func search[T comparable](e T, curr *binaryNode[T]) (*binaryNode[T], error) {
	if curr == nil {
		return nil, fmt.Errorf("%v not found in the tree", e)
	}

	if curr.data == e {
		return curr, nil
	}

	if ele, err := search(e, curr.left); err == nil {
		return ele, nil
	}

	return search(e, curr.right)
}

func rotateRight[T comparable](n, p *binaryNode[T], bt *BinaryTree[T]) {
	if n.isLeaf() || n.left == nil {
		return
	}

	if n == bt.root {
		bt.root = n.left
		n.left.parent = nil
	} else {
		if p.left == n {
			p.left = n.left
		} else {
			p.right = n.left
		}
		n.left.parent = p
	}

	lc := n.left.right
	lp := n.left

	n.left.right = n
	n.left = lc

	n.parent = lp
	if lc != nil {
		lc.parent = n
	}
}

func rotateLeft[T comparable](n, p *binaryNode[T], bt *BinaryTree[T]) {
	if n.isLeaf() || n.right == nil {
		return
	}

	if n == bt.root {
		bt.root = n.right
		n.right.parent = nil
	} else {
		if p.left == n {
			p.left = n.right
		} else {
			p.right = n.right
		}
		n.right.parent = p
	}

	rc := n.right.left
	rp := n.right

	n.right.left = n
	n.right = rc

	n.parent = rp
	if rc != nil {
		rc.parent = n
	}
}

func isComplete[T comparable](n *binaryNode[T], h *int, res *bool) {
	if n == nil {
		return
	}

	if n.left == nil && n.right != nil {
		*res = false
		return
	}

	lh := 0
	rh := 0

	isComplete(n.left, &lh, res)
	isComplete(n.right, &rh, res)

	if rh > lh || lh-rh > 1 {
		*res = false
		return
	}

	*h = 1 + max(lh, rh)
}

func lowestCommonAncestor[T comparable](a, b, r *binaryNode[T]) *binaryNode[T] {
	if r == nil {
		return nil
	}

	if a == r || b == r {
		return r
	}

	lt := lowestCommonAncestor(a, b, r.left)
	rt := lowestCommonAncestor(a, b, r.right)

	if lt != nil && rt != nil {
		return r
	}

	if lt != nil {
		return lt
	}

	return rt
}

func paths[T comparable](n *binaryNode[T], temp *list.LinkedList[T], res list.List[T]) error {
	if n == nil {
		return nil
	}

	if err := temp.AddLast(n.data); err != nil {
		return err
	}

	if n.left == nil && n.right == nil {
		if err := res.Add(temp.ToArrayList()); err != nil {
			return err
		}

		if _, err := temp.RemoveLast(); err != nil {
			return err
		}

		return nil
	}

	if err := paths(n.left, temp, res); err != nil {
		return err
	}

	if err := paths(n.right, temp, res); err != nil {
		return err
	}

	if _, err := temp.RemoveLast(); err != nil {
		return err
	}

	return nil
}

func cloneNodes[T comparable](n *binaryNode[T], p *binaryNode[T]) *binaryNode[T] {
	if n == nil {
		return nil
	}

	bn := &binaryNode{}
	bn.data = n.data
	bn.parent = p
	bn.left = cloneNodes(n.left, bn)
	bn.right = cloneNodes(n.right, bn)

	return bn
}

//TODO FIX EXPENSIVE IMPLEMENTATION
func horizontalIterator[T comparable](bt *BinaryTree[T], kind int) iterator.Iterator[T] {
	q, _ := queue.NewLinkedQueue()
	_ = q.Add(bt.root)

	chd := 0
	bt.root.hd = chd

	m := make(map[int][]T)
	keys, _ := list.NewArrayList()

	for !q.Empty() {

		t, _ := q.Remove()

		if m[t.(*binaryNode[T]).hd] == nil {
			_ = keys.Add(t.(*binaryNode[T]).hd)
			m[t.(*binaryNode[T]).hd] = append(m[t.(*binaryNode[T]).hd], t.(*binaryNode[T]))
		} else {

			if kind == 1 {
				_ = keys.Add(t.(*binaryNode[T]).hd)
				m[t.(*binaryNode[T]).hd] = []T{t.(*binaryNode[T])}
			} else if kind == 2 {
				_ = keys.Add(t.(*binaryNode[T]).hd)
				m[t.(*binaryNode[T]).hd] = append(m[t.(*binaryNode[T]).hd], t.(*binaryNode[T]))
			}

		}

		if t.(*binaryNode[T]).left != nil {
			l := t.(*binaryNode[T]).left
			l.hd = t.(*binaryNode[T]).hd - 1
			_ = q.Add(l)
		}

		if t.(*binaryNode[T]).right != nil {
			l := t.(*binaryNode[T]).right
			l.hd = t.(*binaryNode[T]).hd + 1
			_ = q.Add(l)
		}

	}

	l, _ := list.NewArrayList()
	keys.Sort(comparator.NewIntegerComparator())

	it := keys.Iterator()
	s := make(map[T]bool)

	for it.HasNext() {
		e := it.Next()

		ele := m[e.(int)]

		for _, le := range ele {
			if !s[le] {
				_ = l.Add(le)
				s[le] = true
			}
		}
	}

	return l.Iterator()
}

func inOrderSuccessor[T comparable](n *binaryNode[T]) *binaryNode[T] {
	c := n
	for c != nil && c.left != nil {
		c = c.left
	}
	return c
}
