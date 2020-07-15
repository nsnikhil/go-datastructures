package tree

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/nsnikhil/go-datastructures/queue"
	"github.com/nsnikhil/go-datastructures/set"
	"github.com/nsnikhil/go-datastructures/utils"
)

type node struct {
	data interface{}
	//level  int
	hd     int
	parent *node
	links  set.Set
}

func (n *node) childCount() int {
	if n.links == nil {
		return utils.Naught
	}

	return n.links.Size()
}

func (n *node) detach() error {
	if n == nil || n.parent == nil {
		return nil
	}

	if err := n.parent.links.Remove(n); err != nil {
		return err
	}

	n.parent = nil
	return nil
}

func (n *node) add(c *node) error {
	if n.links == nil {
		hs, err := set.NewHashSet()
		if err != nil {
			return err
		}

		n.links = hs
	}

	return n.links.Add(c)
}

func newNode(e interface{}) *node {
	return &node{data: e}
}

type NAryTree struct {
	maxChildren int
	typeURL     string
	count       int
	root        *node
}

func NewNAryTree(maxChildren int, e ...interface{}) (*NAryTree, error) {
	if maxChildren <= utils.Naught {
		return nil, errors.New("max child count cannot be below 1")
	}

	nt := &NAryTree{
		typeURL:     utils.NA,
		count:       utils.Naught,
		maxChildren: maxChildren,
	}

	sz := len(e)

	if sz == 0 {
		return nt, nil
	}

	et := utils.GetTypeName(e[0])

	for i := 1; i < sz; i++ {
		if et != utils.GetTypeName(e[i]) {
			return nil, errors.New("all elements must be of same type")
		}
	}

	for i := 0; i < sz; i++ {
		if err := nt.Insert(e[i]); err != nil {
			return nil, err
		}
	}

	return nt, nil
}

func (nt *NAryTree) Insert(e interface{}) error {
	et := utils.GetTypeName(e)

	if nt.typeURL == utils.NA {
		nt.typeURL = et
	}

	if nt.typeURL != et {
		return liberror.NewTypeMismatchError(nt.typeURL, et)
	}

	t := newNode(e)

	if nt.root == nil {
		nt.root = t
		nt.count++
		return nil
	}

	q, err := queue.NewLinkedQueue()
	if err != nil {
		return err
	}

	if err = q.Add(nt.root); err != nil {
		return err
	}

	err = func() error {
		for !q.Empty() {
			sz := q.Count()

			for i := 0; i < sz; i++ {
				f, err := q.Remove()
				if err != nil {
					return err
				}

				if f.(*node).childCount() < nt.maxChildren {
					if err := f.(*node).add(t); err != nil {
						return err
					}

					t.parent = f.(*node)
					return nil
				}

				if f.(*node).links == nil {
					continue
				}

				it := f.(*node).links.Iterator()
				for it.HasNext() {
					if err = q.Add(it.Next()); err != nil {
						return err
					}
				}
			}
		}

		return nil
	}()

	if err != nil {
		return err
	}

	nt.count++
	return nil
}

func (nt *NAryTree) Delete(e interface{}) error {
	return nil
}

func (nt *NAryTree) Search(e interface{}) (bool, error) {
	return false, nil
}

func (nt *NAryTree) Count() int {
	return nt.count
}

func (nt *NAryTree) Height() int {
	return utils.InvalidIndex
}

func (nt *NAryTree) Diameter() int {
	return utils.InvalidIndex
}

func (nt *NAryTree) Empty() bool {
	return false
}

func (nt *NAryTree) Clear() {

}

func (nt *NAryTree) Clone() Tree {
	return nil
}

func (nt *NAryTree) Mirror() (bool, error) {
	return false, nil
}

func (nt *NAryTree) MirrorAt(e interface{}) (bool, error) {
	return false, nil
}

func (nt *NAryTree) RotateLeft() error {
	return nil
}

func (nt *NAryTree) RotateRight() error {
	return nil
}

func (nt *NAryTree) RotateLeftAt(e interface{}) error {
	return nil
}

func (nt *NAryTree) RotateRightAt(e interface{}) error {
	return nil
}

func (nt *NAryTree) IsFull() bool {
	return false
}

func (nt *NAryTree) IsBalanced() bool {
	return false
}

func (nt *NAryTree) IsPerfect() bool {
	return false
}

func (nt *NAryTree) IsComplete() bool {
	return false
}

func (nt *NAryTree) LowestCommonAncestor(a, b interface{}) (interface{}, error) {
	return nil, nil
}

func (nt *NAryTree) Paths() (list.List, error) {
	return nil, nil
}

func (nt *NAryTree) Mode() (list.List, error) {
	return nil, nil
}

func (nt *NAryTree) Equal(t Tree) (bool, error) {
	return false, nil
}

func (nt *NAryTree) Symmetric() bool {
	return false
}

func (nt *NAryTree) Invert() {

}

func (nt *NAryTree) InOrderSuccessor(e interface{}) (interface{}, error) {
	return nil, nil
}

func (nt *NAryTree) PreOrderSuccessor(e interface{}) (interface{}, error) {
	return nil, nil
}

func (nt *NAryTree) PostOrderSuccessor(e interface{}) (interface{}, error) {
	return nil, nil
}

func (nt *NAryTree) LevelOrderSuccessor(e interface{}) (interface{}, error) {
	return nil, nil
}

func (nt *NAryTree) PreOrderIterator() iterator.Iterator {
	return nil
}

func (nt *NAryTree) PostOrderIterator() iterator.Iterator {
	return nil
}

func (nt *NAryTree) InOrderIterator() iterator.Iterator {
	return nil
}

func (nt *NAryTree) LevelOrderIterator() iterator.Iterator {
	return nil
}

func (nt *NAryTree) VerticalViewIterator() iterator.Iterator {
	return nil
}

func (nt *NAryTree) LeftViewIterator() iterator.Iterator {
	return nil
}

func (nt *NAryTree) RightViewIterator() iterator.Iterator {
	return nil
}

func (nt *NAryTree) TopViewIterator() iterator.Iterator {
	return nil
}

func (nt *NAryTree) BottomViewIterator() iterator.Iterator {
	return nil
}
