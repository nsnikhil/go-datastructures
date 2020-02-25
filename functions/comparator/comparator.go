package comparator

type Comparator interface {
	Compare(one interface{}, two interface{}) (int, error)
}
