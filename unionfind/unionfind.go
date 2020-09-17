package unionfind

type UnionFind interface {
	Find(i int) int
	Union(i, j int)
}

type defaultUnionFind struct {
	par []int
}

func (duf *defaultUnionFind) Find(i int) int {
	if duf.par[i] < 0 {
		return i
	}

	t := duf.Find(duf.par[i])
	duf.par[i] = t
	return t
}

func (duf *defaultUnionFind) Union(i, j int) {
	x := duf.Find(i)
	y := duf.Find(j)

	if x == y {
		return
	}

	px := duf.par[x]
	py := duf.par[y]

	if (px * -1) < (py * -1) {
		duf.par[x] = y
		duf.par[y] += px
	} else {
		duf.par[y] = x
		duf.par[x] += py
	}
}

func NewUnionFind(sz int) UnionFind {
	par := make([]int, sz)

	for i := 0; i < sz; i++ {
		par[i] = -1
	}

	return &defaultUnionFind{
		par: par,
	}
}
