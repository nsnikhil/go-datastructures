package trie

import (
	"errors"
	"fmt"
)

const a = 97

type node struct {
	endOfWord bool
	links     map[rune]*node
}

func newNode() *node {
	return &node{
		links: make(map[rune]*node),
	}
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{newNode()}
}

func (t *Trie) Insert(s string) error {
	n := t.root
	if n == nil {
		return errors.New("root is nil")
	}

	for _, d := range s {
		if n.links[d-a] == nil {
			n.links[d-a] = newNode()
		}
		n = n.links[d-a]
	}

	n.endOfWord = true
	return nil
}

func (t *Trie) SearchPrefix(prefix string) bool {
	return search(prefix, t.root) != nil
}

func (t *Trie) SearchWord(word string) bool {
	n := search(word, t.root)
	if n == nil {
		return false
	}

	return n.endOfWord
}

func search(word string, n *node) *node {
	if n == nil {
		return nil
	}

	for _, q := range word {
		idx := q - a
		if n.links[idx] == nil {
			return nil
		}

		n = n.links[idx]
	}

	return n
}

func (t *Trie) Get(prefix string) []string {
	n := search(prefix, t.root)
	if n == nil {
		return []string{}
	}

	res := make([]string, 0)
	traverse(n, prefix, &res)

	return res
}

func traverse(n *node, prefix string, res *[]string) {
	if n == nil {
		return
	}

	if n.endOfWord {
		*res = append(*res, prefix)
	}

	for i, l := range n.links {
		traverse(l, fmt.Sprintf("%s%c", prefix, i+a), res)
	}
}
