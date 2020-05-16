package gmap

type Pair struct {
	k interface{}
	v interface{}
}

func NewPair(k, v interface{}) *Pair {
	return &Pair{
		k: k,
		v: v,
	}
}

func (p Pair) GetKey() interface{} {
	return p.k
}

func (p Pair) GetValue() interface{} {
	return p.v
}
