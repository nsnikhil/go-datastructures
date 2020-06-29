package consumer

type BiConsumer interface {
	Accept(e, f interface{})
}
