package function

type Function[T comparable, R comparable] interface {
	Apply(e T) R
}
