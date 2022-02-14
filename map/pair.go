package gmap

type Pair[F comparable, S comparable] struct {
	first  F
	second S
}

func NewPair[F comparable, S comparable](first F, second S) *Pair[F, S] {
	return &Pair[F, S]{
		first:  first,
		second: second,
	}
}

func (p Pair[F, S]) First() F {
	return p.first
}

func (p Pair[F, S]) Second() S {
	return p.second
}
