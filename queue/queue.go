package queue

type Queue[T comparable] interface {
	Add(e T) error

	Remove() (T, error)

	Peek() (T, error)

	Empty() bool
	Size() int64
	Clear()
}
