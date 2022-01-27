package operator

type UnaryOperator[T any] interface {
	Apply(e T) T
}
