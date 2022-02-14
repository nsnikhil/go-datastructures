package operator

type BinaryOperator[T comparable, U comparable, R comparable] interface {
	Apply(t T, u U) R
}
