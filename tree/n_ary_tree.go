package tree

import (
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/nsnikhil/go-datastructures/utils"
)

type NAryTree[T comparable] struct {
}

func NewNAryTree[T comparable](e ...T) (*NAryTree[T], error) {
	return nil, nil
}

func (nt *NAryTree[T]) Insert(e T) error {
	return nil
}

func (nt *NAryTree[T]) Delete(e T) error {
	return nil
}

func (nt *NAryTree[T]) Search(e T) (bool, error) {
	return false, nil
}

func (nt *NAryTree[T]) Count() int {
	return utils.InvalidIndex
}

func (nt *NAryTree[T]) Height() int {
	return utils.InvalidIndex
}

func (nt *NAryTree[T]) Diameter() int {
	return utils.InvalidIndex
}

func (nt *NAryTree[T]) Empty() bool {
	return false
}

func (nt *NAryTree[T]) Clear() {

}

func (nt *NAryTree[T]) Clone() Tree[T] {
	return nil
}

func (nt *NAryTree[T]) Mirror() (bool, error) {
	return false, nil
}

func (nt *NAryTree[T]) MirrorAt(e T) (bool, error) {
	return false, nil
}

func (nt *NAryTree[T]) RotateLeft() error {
	return nil
}

func (nt *NAryTree[T]) RotateRight() error {
	return nil
}

func (nt *NAryTree[T]) RotateLeftAt(e T) error {
	return nil
}

func (nt *NAryTree[T]) RotateRightAt(e T) error {
	return nil
}

func (nt *NAryTree[T]) IsFull() bool {
	return false
}

func (nt *NAryTree[T]) IsBalanced() bool {
	return false
}

func (nt *NAryTree[T]) IsPerfect() bool {
	return false
}

func (nt *NAryTree[T]) IsComplete() bool {
	return false
}

func (nt *NAryTree[T]) LowestCommonAncestor(a, b T) (T, error) {
	return nil, nil
}

func (nt *NAryTree[T]) Paths() (list.List[T], error) {
	return nil, nil
}

func (nt *NAryTree[T]) Mode() (list.List[T], error) {
	return nil, nil
}

func (nt *NAryTree[T]) Equal(t Tree[T]) (bool, error) {
	return false, nil
}

func (nt *NAryTree[T]) Symmetric() bool {
	return false
}

func (nt *NAryTree[T]) Invert() {

}

func (nt *NAryTree[T]) InOrderSuccessor(e T) (T, error) {
	return nil, nil
}

func (nt *NAryTree[T]) PreOrderSuccessor(e T) (T, error) {
	return nil, nil
}

func (nt *NAryTree[T]) PostOrderSuccessor(e T) (T, error) {
	return nil, nil
}

func (nt *NAryTree[T]) LevelOrderSuccessor(e T) (T, error) {
	return nil, nil
}

func (nt *NAryTree[T]) PreOrderIterator() iterator.Iterator[T] {
	return nil
}

func (nt *NAryTree[T]) PostOrderIterator() iterator.Iterator[T] {
	return nil
}

func (nt *NAryTree[T]) InOrderIterator() iterator.Iterator[T] {
	return nil
}

func (nt *NAryTree[T]) LevelOrderIterator() iterator.Iterator[T] {
	return nil
}

func (nt *NAryTree[T]) VerticalViewIterator() iterator.Iterator[T] {
	return nil
}

func (nt *NAryTree[T]) LeftViewIterator() iterator.Iterator[T] {
	return nil
}

func (nt *NAryTree[T]) RightViewIterator() iterator.Iterator[T] {
	return nil
}

func (nt *NAryTree[T]) TopViewIterator() iterator.Iterator[T] {
	return nil
}

func (nt *NAryTree[T]) BottomViewIterator() iterator.Iterator[T] {
	return nil
}
