package operator

type UnaryOperator interface {
	Apply(e interface{}) interface{}
}
