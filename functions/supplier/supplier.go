package supplier

type Supplier[T any] interface {
	Get() T
}
