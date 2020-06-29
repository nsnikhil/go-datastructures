package consumer

type Consumer interface {
	Accept(e interface{})
}
