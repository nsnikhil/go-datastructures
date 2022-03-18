package operator

type BinaryOperator[T any, U any, R any] interface {
	Apply(t T, u U) R
}
