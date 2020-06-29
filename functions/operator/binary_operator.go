package operator

type BinaryOperator interface {
	Apply(t interface{}, u interface{}) interface{}
}
