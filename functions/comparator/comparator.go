package comparator

type Comparator[T any] interface {
	Compare(one T, two T) int
}
