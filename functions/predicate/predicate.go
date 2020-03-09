package predicate

type Predicate interface {
	Test(e interface{}) bool
}
