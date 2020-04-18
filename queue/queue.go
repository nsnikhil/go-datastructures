package queue

type Queue interface {
	Add(e interface{}) error

	Remove() (interface{}, error)

	Peek() (interface{}, error)

	Empty() bool
	Count() int
	Clear()
}
