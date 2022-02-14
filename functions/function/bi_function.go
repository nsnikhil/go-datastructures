package function

type BiFunction[K comparable, V comparable, R comparable] interface {
	Apply(t K, u V) R
}
