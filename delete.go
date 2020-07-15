package main

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/tree"
)

func printTree(t tree.Tree) {
	it := t.InOrderIterator()
	for it.HasNext() {
		fmt.Printf("%v ", it.Next())
	}
	fmt.Println()
}

func main() {
	bst, _ := tree.NewBinarySearchTree(comparator.NewIntegerComparator())

	bst.Insert(6)
	bst.Insert(2)
	bst.Insert(8)
	bst.Insert(0)
	bst.Insert(4)
	bst.Insert(7)
	bst.Insert(9)
	bst.Insert(2)
	bst.Insert(6)

	//printTree(bst)

	nat, _ := tree.NewNAryTree(1)
	nat.Insert(1)
	nat.Insert(2)
	nat.Insert(3)
	nat.Insert(4)


}
