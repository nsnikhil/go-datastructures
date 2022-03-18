package queue

type Queue[T comparable] interface {
	Add(e T)

	Remove() (T, error)

	Peek() (T, error)

	Empty() bool
	Size() int64
	Clear()
}
