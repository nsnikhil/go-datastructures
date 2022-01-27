package predicate

type Predicate[T any] interface {
	Test(e T) bool
}
