package internal

func ZeroValueOf[T any]() T {
	return *new(T)
}
