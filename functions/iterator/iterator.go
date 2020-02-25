package iterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
