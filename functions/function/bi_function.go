package function

type BiFunction interface {
	Apply(t interface{}, u interface{}) interface{}
}
