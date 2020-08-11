package tree

import (
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/nsnikhil/go-datastructures/utils"
)

type NAryTree struct {
}

func NewNAryTree(e ...interface{}) (*NAryTree, error) {
	return nil, nil
}

func (nt *NAryTree) Insert(e interface{}) error {
	return nil
}

func (nt *NAryTree) Delete(e interface{}) error {
	return nil
}

func (nt *NAryTree) Search(e interface{}) (bool, error) {
	return false, nil
}

func (nt *NAryTree) Count() int {
	return utils.InvalidIndex
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
