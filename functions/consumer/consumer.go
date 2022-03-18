package consumer

type Consumer[T comparable] interface {
	Accept(e T)
}
