package internal

func ZeroValueOf[T comparable]() T {
	return *new(T)
}
