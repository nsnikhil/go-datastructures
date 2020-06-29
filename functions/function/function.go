package function

type Function interface {
	Apply(e interface{}) interface{}
}
