package consumer

type BiConsumer[T any, K any] interface {
	Accept(e T, f K)
}
