package tree

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/nsnikhil/go-datastructures/queue"
	"github.com/nsnikhil/go-datastructures/stack"
	"github.com/nsnikhil/go-datastructures/utils"
)

type binaryNode struct {
	data interface{}
	//level  int // NOT IMPLEMENTED
	hd     int
	left   *binaryNode
	right  *binaryNode
	parent *binaryNode
}

func (bn *binaryNode) String() string {
	return fmt.Sprint(bn.data)
}

func (bn *binaryNode) isLeaf() bool {
	return bn.left == nil && bn.right == nil
}

func (bn *binaryNode) childCount() int {
	c := 0
	if bn.left != nil {
		c++
	}
	if bn.right != nil {
		c++
	}
	return c
}

func (bn *binaryNode) clone() *binaryNode {
	return &binaryNode{
		data:   bn.data,
		hd:     bn.hd,
		left:   bn.left,
		right:  bn.right,
		parent: bn.parent,
	}
}

func (bn *binaryNode) detach() {
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

func newBinaryNode(data interface{}) *binaryNode {
	return &binaryNode{data: data}
}

type BinaryTree struct {
	typeURL string
	count   int
	root    *binaryNode
}

func NewBinaryTree(e ...interface{}) (*BinaryTree, error) {
	bt := &BinaryTree{
		typeURL: utils.NA,
		count:   utils.Naught,
	}

	sz := len(e)

	if sz == 0 {
		return bt, nil
	}

	et := utils.GetTypeName(e[0])

	for i := 1; i < sz; i++ {
		if et != utils.GetTypeName(e[i]) {
			return nil, errors.New("all elements must be of same type")
		}
	}

	for i := 0; i < sz; i++ {
		if err := bt.Insert(e[i]); err != nil {
			return nil, err
		}
	}

	return bt, nil
}

func (bt *BinaryTree) Insert(e interface{}) error {
	et := utils.GetTypeName(e)

	if bt.typeURL == utils.NA {
		bt.typeURL = et
	}

	if bt.typeURL != et {
		return liberror.NewTypeMismatchError(bt.typeURL, et)
	}

	t := newBinaryNode(e)

	if bt.root == nil {
		bt.root = t
		bt.count++
		return nil
	}

	curr := bt.root

	q, err := queue.NewLinkedQueue()
	if err != nil {
		return err
	}

	if err := q.Add(curr); err != nil {
		return err
	}

	err = func() error {
		for !q.Empty() {
			sz := q.Count()

			for i := 0; i < sz; i++ {
				f, _ := q.Remove()

				if f.(*binaryNode).left == nil {
					f.(*binaryNode).left = t
					t.parent = f.(*binaryNode)
					return nil
				}

				if f.(*binaryNode).right == nil {
					f.(*binaryNode).right = t
					t.parent = f.(*binaryNode)
					return nil
				}

				if err := q.Add(f.(*binaryNode).left); err != nil {
					return err
				}

				if err := q.Add(f.(*binaryNode).right); err != nil {
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

func (bt *BinaryTree) InsertCompare(e interface{}, c comparator.Comparator) error {
	et := utils.GetTypeName(e)

	if bt.typeURL == utils.NA {
		bt.typeURL = et
	}

	if bt.typeURL != et {
		return liberror.NewTypeMismatchError(bt.typeURL, et)
	}

	t := newBinaryNode(e)

	if bt.root == nil {
		bt.root = t
		bt.count++
		return nil
	}

	curr := bt.root

	for {
		i, err := c.Compare(curr.data, t.data)
		if err != nil {
			return err
		}

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

	i, err := c.Compare(curr.data, t.data)
	if err != nil {
		return err
	}

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

func (bt *BinaryTree) Delete(e interface{}) error {
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

func (bt *BinaryTree) DeleteCompare(e interface{}, c comparator.Comparator) error {
	var deleteNode func(e interface{}, c comparator.Comparator, n *binaryNode) error

	deleteNode = func(e interface{}, c comparator.Comparator, n *binaryNode) error {
		if n == nil {
			return fmt.Errorf("%v not found in the tree", e)
		}

		i, _ := c.Compare(n.data, e)

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

func (bt *BinaryTree) Search(e interface{}) (bool, error) {
	if bt.Empty() {
		return false, errors.New("tree is empty")
	}

	et := utils.GetTypeName(e)
	if bt.typeURL != et {
		return false, liberror.NewTypeMismatchError(bt.typeURL, et)
	}

	_, err := search(e, bt.root)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (bt *BinaryTree) SearchCompare(e interface{}, c comparator.Comparator) (bool, error) {
	if bt.Empty() {
		return false, errors.New("tree is empty")
	}

	et := utils.GetTypeName(e)
	if et != bt.typeURL {
		return false, liberror.NewTypeMismatchError(bt.typeURL, et)
	}

	curr := bt.root

	for curr != nil {
		if curr.data == e {
			return true, nil
		}

		res, err := c.Compare(curr.data, e)
		if err != nil {
			return false, err
		}

		if res > 0 {
			curr = curr.left
		} else if res < 0 {
			curr = curr.right
		}
	}

	return false, fmt.Errorf("%v not found in the tree", e)
}

func (bt *BinaryTree) Count() int {
	return bt.count
}

func (bt *BinaryTree) Height() int {
	return calculateHeight(bt.root, nil)
}

func (bt *BinaryTree) Diameter() int {
	d := utils.Naught
	calculateHeight(bt.root, &d)
	return d
}

func (bt *BinaryTree) Empty() bool {
	return bt.count == utils.Naught
}

func (bt *BinaryTree) Clear() {
	bt.root = nil
	bt.count = utils.Naught
}

func (bt *BinaryTree) Clone() Tree {
	return &BinaryTree{
		typeURL: bt.typeURL,
		count:   bt.count,
		root:    cloneNodes(bt.root, nil),
	}
}

func (bt *BinaryTree) Mirror() (bool, error) {
	if bt.Empty() {
		return false, errors.New("tree is empty")
	}

	return bt.MirrorAt(bt.root.data)
}

func (bt *BinaryTree) MirrorAt(e interface{}) (bool, error) {
	if bt.Empty() {
		return false, errors.New("tree is empty")
	}

	et := utils.GetTypeName(e)
	if bt.typeURL != et {
		return false, liberror.NewTypeMismatchError(bt.typeURL, et)
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

func (bt *BinaryTree) RotateLeft() error {
	if bt.Empty() {
		return errors.New("tree is empty")
	}

	rotateLeft(bt.root, nil, bt)
	return nil
}

func (bt *BinaryTree) RotateRight() error {
	if bt.Empty() {
		return errors.New("tree is empty")
	}

	rotateRight(bt.root, nil, bt)
	return nil
}

func (bt *BinaryTree) RotateLeftAt(e interface{}) error {
	if bt.Empty() {
		return errors.New("tree is empty")
	}

	et := utils.GetTypeName(e)
	if bt.typeURL != et {
		return liberror.NewTypeMismatchError(bt.typeURL, et)
	}

	curr, err := search(e, bt.root)
	if err != nil {
		return err
	}

	rotateLeft(curr, curr.parent, bt)
	return nil
}

func (bt *BinaryTree) RotateRightAt(e interface{}) error {
	if bt.Empty() {
		return errors.New("tree is empty")
	}

	et := utils.GetTypeName(e)
	if bt.typeURL != et {
		return liberror.NewTypeMismatchError(bt.typeURL, et)
	}

	curr, err := search(e, bt.root)
	if err != nil {
		return err
	}

	rotateRight(curr, curr.parent, bt)
	return nil
}

func (bt *BinaryTree) IsBalanced() bool {
	return isBalancedAt(bt.root)
}

func (bt *BinaryTree) IsFull() bool {
	it := newBtLvOrderIterator(bt)
	it.(*btLvOrderIterator).v = true

	for it.HasNext() {
		if it.Next().(*binaryNode).childCount() == 1 {
			return false
		}
	}

	return true
}

func (bt *BinaryTree) IsPerfect() bool {
	curr := bt.root
	q, _ := queue.NewLinkedQueue()
	_ = q.Add(curr)
	c := 0
	fl := -1

	for !q.Empty() {
		sz := q.Count()

		for i := 0; i < sz; i++ {
			e, _ := q.Remove()

			if e.(*binaryNode).isLeaf() {
				if fl == -1 {
					fl = c
				}
				if fl != c {
					return false
				}
			}

			if e.(*binaryNode).left != nil {
				_ = q.Add(e.(*binaryNode).left)
			}

			if e.(*binaryNode).right != nil {
				_ = q.Add(e.(*binaryNode).right)
			}
		}
		c++
	}

	return true
}

func (bt *BinaryTree) IsComplete() bool {
	h := 0
	res := true
	isComplete(bt.root, &h, &res)
	return res
}

func (bt *BinaryTree) LowestCommonAncestor(a, b interface{}) (interface{}, error) {
	an, err := search(a, bt.root)
	if err != nil {
		return nil, err
	}

	bn, err := search(b, bt.root)
	if err != nil {
		return nil, err
	}

	return lowestCommonAncestor(an, bn, bt.root).data, nil
}

func (bt *BinaryTree) Paths() (list.List, error) {
	if bt.Empty() {
		return nil, errors.New("tree is empty")
	}

	temp, err := list.NewLinkedList()
	if err != nil {
		return nil, err
	}

	res, err := list.NewArrayList()
	if err != nil {
		return nil, err
	}

	if err := paths(bt.root, temp, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (bt *BinaryTree) Mode() (list.List, error) {
	if bt.Empty() {
		return nil, errors.New("tree is empty")
	}

	res, err := list.NewArrayList()
	if err != nil {
		return nil, err
	}

	cm := 0
	m := 0
	var p interface{}

	var mode func(n *binaryNode, l list.List, cm, m *int, p *interface{}) error
	mode = func(n *binaryNode, l list.List, cm, m *int, p *interface{}) error {
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

func (bt *BinaryTree) Equal(t Tree) (bool, error) {
	_, ok := t.(*BinaryTree)
	if !ok {
		return false, errors.New("TODO")
	}

	var equal func(a, b *binaryNode) bool

	equal = func(a, b *binaryNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil || a.data != b.data {
			return false
		}
		return equal(a.left, b.left) && equal(a.right, b.right)
	}

	return equal(bt.root, t.(*BinaryTree).root), nil
}

func (bt *BinaryTree) Symmetric() bool {
	var symmetric func(a, b *binaryNode) bool

	symmetric = func(a, b *binaryNode) bool {
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

func (bt *BinaryTree) Invert() {
	var inverter func(n *binaryNode) *binaryNode

	inverter = func(n *binaryNode) *binaryNode {
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

func (bt *BinaryTree) PreOrderSuccessor(e interface{}) (interface{}, error) {
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
			curr = top.(*binaryNode)
		}

	}

	if prev != nil && prev.data == e {
		return nil, fmt.Errorf("no pre order successor found for %v", e)
	}

	return nil, fmt.Errorf("%v not found in the tree", e)
}

func (bt *BinaryTree) PostOrderSuccessor(e interface{}) (interface{}, error) {
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
			curr = top.(*binaryNode)
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

func (bt *BinaryTree) InOrderSuccessor(e interface{}) (interface{}, error) {
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
			curr = top.(*binaryNode)
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

func (bt *BinaryTree) LevelOrderSuccessor(e interface{}) (interface{}, error) {
	q, err := queue.NewLinkedQueue()
	if err != nil {
		return nil, err
	}

	var prev *binaryNode
	_ = q.Add(bt.root)

	for !q.Empty() {

		sz := q.Count()

		for i := 0; i < sz; i++ {
			f, _ := q.Remove()

			if prev != nil && prev.data == e {
				return f.(*binaryNode).data, nil
			}

			prev = f.(*binaryNode)

			if f.(*binaryNode).left != nil {
				_ = q.Add(f.(*binaryNode).left)
			}

			if f.(*binaryNode).right != nil {
				_ = q.Add(f.(*binaryNode).right)
			}
		}
	}

	if prev != nil && prev.data == e {
		return nil, fmt.Errorf("no level order successor found for %v", e)
	}

	return nil, fmt.Errorf("%v not found in the tree", e)
}

func (bt *BinaryTree) PreOrderIterator() iterator.Iterator {
	return newBtPreOrderIterator(bt)
}

type btPreOrderIterator struct {
	curr *binaryNode
	s    *stack.Stack
	v    bool
}

func newBtPreOrderIterator(bt *BinaryTree) *btPreOrderIterator {
	s, _ := stack.NewStack()
	return &btPreOrderIterator{
		curr: bt.root,
		s:    s,
	}
}

func (bti *btPreOrderIterator) HasNext() bool {
	return bti.curr != nil || !bti.s.Empty()
}

func (bti *btPreOrderIterator) Next() interface{} {
	if bti.curr == nil {
		n, _ := bti.s.Pop()
		bti.curr = n.(*binaryNode)
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

func (bt *BinaryTree) PostOrderIterator() iterator.Iterator {
	return newBtPostOrderIterator(bt)
}

type btPostOrderIterator struct {
	curr *binaryNode
	last *binaryNode
	s    *stack.Stack
	v    bool
}

func newBtPostOrderIterator(bt *BinaryTree) *btPostOrderIterator {
	s, _ := stack.NewStack()
	return &btPostOrderIterator{
		curr: bt.root,
		s:    s,
	}
}

func (bto *btPostOrderIterator) HasNext() bool {
	return bto.curr != nil || !bto.s.Empty()
}

func (bto *btPostOrderIterator) Next() interface{} {
	get := func() interface{} {
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

		bto.curr = top.(*binaryNode)

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
			bto.curr = top.(*binaryNode)
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

func (bt *BinaryTree) InOrderIterator() iterator.Iterator {
	return newBtInOrderIterator(bt)
}

type btInOrderIterator struct {
	curr *binaryNode
	s    *stack.Stack
	v    bool
}

func newBtInOrderIterator(bt *BinaryTree) *btInOrderIterator {
	s, _ := stack.NewStack()
	return &btInOrderIterator{
		curr: bt.root,
		s:    s,
	}
}

func (bti *btInOrderIterator) HasNext() bool {
	return bti.curr != nil || !bti.s.Empty()
}

func (bti *btInOrderIterator) Next() interface{} {
	for bti.curr != nil {
		_ = bti.s.Push(bti.curr)
		bti.curr = bti.curr.left
	}

	if !bti.s.Empty() {
		top, _ := bti.s.Pop()
		bti.curr = top.(*binaryNode)
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

func (bt *BinaryTree) LevelOrderIterator() iterator.Iterator {
	return newBtLvOrderIterator(bt)
}

type btLvOrderIterator struct {
	curr *binaryNode
	q    queue.Queue
	v    bool
}

func newBtLvOrderIterator(bt *BinaryTree) iterator.Iterator {
	q, _ := queue.NewLinkedQueue()
	_ = q.Add(bt.root)

	return &btLvOrderIterator{
		curr: bt.root,
		q:    q,
	}
}

func (blv *btLvOrderIterator) HasNext() bool {
	return !blv.q.Empty()
}

func (blv *btLvOrderIterator) Next() interface{} {
	curr, _ := blv.q.Remove()

	if curr.(*binaryNode).left != nil {
		_ = blv.q.Add(curr.(*binaryNode).left)
	}

	if curr.(*binaryNode).right != nil {
		_ = blv.q.Add(curr.(*binaryNode).right)
	}

	if blv.v {
		return curr
	}

	return curr.(*binaryNode).data
}

//TODO FIX EXPENSIVE IMPLEMENTATION
func (bt *BinaryTree) VerticalViewIterator() iterator.Iterator {
	return newBtVerticalVOrderIterator(bt)
}

type btVerticalVOrderIterator struct {
	it iterator.Iterator
	v  bool
}

func newBtVerticalVOrderIterator(bt *BinaryTree) iterator.Iterator {
	return &btVerticalVOrderIterator{
		it: horizontalIterator(bt, 2),
	}
}

func (btv *btVerticalVOrderIterator) HasNext() bool {
	return btv.it.HasNext()
}

func (btv *btVerticalVOrderIterator) Next() interface{} {
	if btv.v {
		return btv.it.Next()
	}

	return btv.it.Next().(*binaryNode).data
}

func (bt *BinaryTree) LeftViewIterator() iterator.Iterator {
	return newBtLeftVOrderIterator(bt)
}

type btLeftVOrderIterator struct {
	curr *binaryNode
	q    queue.Queue
	v    bool
}

func newBtLeftVOrderIterator(bt *BinaryTree) iterator.Iterator {
	q, _ := queue.NewLinkedQueue()
	_ = q.Add(bt.root)

	return &btLeftVOrderIterator{
		curr: bt.root,
		q:    q,
	}
}

func (bfv *btLeftVOrderIterator) HasNext() bool {
	return !bfv.q.Empty()
}

func (bfv *btLeftVOrderIterator) Next() interface{} {
	sz := bfv.q.Count()

	var res *binaryNode = nil

	for i := 0; i < sz; i++ {
		curr, _ := bfv.q.Remove()

		if res == nil {
			res = curr.(*binaryNode)
		}

		if curr.(*binaryNode).left != nil {
			_ = bfv.q.Add(curr.(*binaryNode).left)
		}

		if curr.(*binaryNode).right != nil {
			_ = bfv.q.Add(curr.(*binaryNode).right)
		}

	}

	if bfv.v {
		return res
	}

	return res.data
}

func (bt *BinaryTree) RightViewIterator() iterator.Iterator {
	return newBtRightVOrderIterator(bt)
}

type btRightVOrderIterator struct {
	curr *binaryNode
	q    queue.Queue
	v    bool
}

func newBtRightVOrderIterator(bt *BinaryTree) iterator.Iterator {
	q, _ := queue.NewLinkedQueue()
	_ = q.Add(bt.root)

	return &btRightVOrderIterator{
		curr: bt.root,
		q:    q,
	}
}

func (brv *btRightVOrderIterator) HasNext() bool {
	return !brv.q.Empty()
}

func (brv *btRightVOrderIterator) Next() interface{} {
	sz := brv.q.Count()

	var res *binaryNode = nil

	for i := 0; i < sz; i++ {
		curr, _ := brv.q.Remove()

		if res == nil {
			res = curr.(*binaryNode)
		}

		if curr.(*binaryNode).right != nil {
			_ = brv.q.Add(curr.(*binaryNode).right)
		}

		if curr.(*binaryNode).left != nil {
			_ = brv.q.Add(curr.(*binaryNode).left)
		}
	}

	if brv.v {
		return res
	}

	return res.data
}

//TODO FIX EXPENSIVE IMPLEMENTATION
func (bt *BinaryTree) TopViewIterator() iterator.Iterator {
	return newBtTopVOrderIterator(bt.Clone().(*BinaryTree))
}

type btTopVOrderIterator struct {
	it iterator.Iterator
	v  bool
}

func newBtTopVOrderIterator(bt *BinaryTree) iterator.Iterator {
	return &btTopVOrderIterator{
		it: horizontalIterator(bt, 0),
	}
}

func (btv *btTopVOrderIterator) HasNext() bool {
	return btv.it.HasNext()
}

func (btv *btTopVOrderIterator) Next() interface{} {
	if btv.v {
		return btv.it.Next()
	}

	return btv.it.Next().(*binaryNode).data
}

//TODO FIX EXPENSIVE IMPLEMENTATION
func (bt *BinaryTree) BottomViewIterator() iterator.Iterator {
	return newBtBottomVOrderIterator(bt.Clone().(*BinaryTree))
}

type btBottomVOrderIterator struct {
	it iterator.Iterator
	v  bool
}

func newBtBottomVOrderIterator(bt *BinaryTree) iterator.Iterator {
	return &btBottomVOrderIterator{
		it: horizontalIterator(bt, 1),
	}
}

func (brv *btBottomVOrderIterator) HasNext() bool {
	return brv.it.HasNext()
}

func (brv *btBottomVOrderIterator) Next() interface{} {
	if brv.v {
		return brv.it.Next()
	}

	return brv.it.Next().(*binaryNode).data
}

func lastNode(bt *BinaryTree) (*binaryNode, *binaryNode) {
	if bt.root == nil {
		return nil, nil
	}

	var prev *binaryNode = nil
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
func mirrorAt(n *binaryNode) error {
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
			curr = top.(*binaryNode)
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

func calculateHeight(n *binaryNode, diameter *int) int {
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

func isBalancedAt(n *binaryNode) bool {
	if n == nil {
		return true
	}
	return absDiff(calculateHeight(n.left, nil), calculateHeight(n.right, nil)) <= 1 && isBalancedAt(n.left) && isBalancedAt(n.right)
}

func search(e interface{}, curr *binaryNode) (*binaryNode, error) {
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

func rotateRight(n, p *binaryNode, bt *BinaryTree) {
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

func rotateLeft(n, p *binaryNode, bt *BinaryTree) {
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

func isComplete(n *binaryNode, h *int, res *bool) {
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

func lowestCommonAncestor(a, b, r *binaryNode) *binaryNode {
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

func paths(n *binaryNode, temp *list.LinkedList, res list.List) error {
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

func cloneNodes(n *binaryNode, p *binaryNode) *binaryNode {
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
func horizontalIterator(bt *BinaryTree, kind int) iterator.Iterator {
	q, _ := queue.NewLinkedQueue()
	_ = q.Add(bt.root)

	chd := 0
	bt.root.hd = chd

	m := make(map[int][]interface{})
	keys, _ := list.NewArrayList()

	for !q.Empty() {

		t, _ := q.Remove()

		if m[t.(*binaryNode).hd] == nil {
			_ = keys.Add(t.(*binaryNode).hd)
			m[t.(*binaryNode).hd] = append(m[t.(*binaryNode).hd], t.(*binaryNode))
		} else {

			if kind == 1 {
				_ = keys.Add(t.(*binaryNode).hd)
				m[t.(*binaryNode).hd] = []interface{}{t.(*binaryNode)}
			} else if kind == 2 {
				_ = keys.Add(t.(*binaryNode).hd)
				m[t.(*binaryNode).hd] = append(m[t.(*binaryNode).hd], t.(*binaryNode))
			}

		}

		if t.(*binaryNode).left != nil {
			l := t.(*binaryNode).left
			l.hd = t.(*binaryNode).hd - 1
			_ = q.Add(l)
		}

		if t.(*binaryNode).right != nil {
			l := t.(*binaryNode).right
			l.hd = t.(*binaryNode).hd + 1
			_ = q.Add(l)
		}

	}

	l, _ := list.NewArrayList()
	keys.Sort(comparator.NewIntegerComparator())

	it := keys.Iterator()
	s := make(map[interface{}]bool)

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

func inOrderSuccessor(n *binaryNode) *binaryNode {
	c := n
	for c != nil && c.left != nil {
		c = c.left
	}
	return c
}
