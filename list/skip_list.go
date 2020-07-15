package list

import (
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/nsnikhil/go-datastructures/utils"
	"math/rand"
	"time"
)

const (
	heads = "heads"
	tails = "tails"
)

type skipNode struct {
	data   interface{}
	next   *skipNode
	prev   *skipNode
	top    *skipNode
	bottom *skipNode
}

func newSkipNode(e interface{}) *skipNode {
	return &skipNode{
		data: e,
	}
}

type skipList struct {
	typeURL string
	first   *skipNode
	last    *skipNode
}

func newSkipList() *skipList {
	return &skipList{}
}

func (sl *skipList) Add(e interface{}) error {

	return nil
}

func insert(e interface{}, sl *skipList) error {
	et := utils.GetTypeName(e)

	if sl.typeURL == utils.NA {
		sl.typeURL = et
	}

	if et != sl.typeURL {
		return liberror.NewTypeMismatchError(sl.typeURL, et)
	}

	temp := newSkipNode(e)

	if sl.IsEmpty() {
		sl.first = temp
		sl.last = temp
	}

	ct := sl.last

	for coinToss() == heads {
		ct.top = newSkipNode(e)
		ct = ct.top
	}

	return nil
}

func (sl *skipList) IsEmpty() bool {
	return false
}

func coinToss() string {
	rand.NewSource(time.Now().Unix())
	if rand.Intn(1-0+1)+0 == 1 {
		return heads
	}
	return tails
}
