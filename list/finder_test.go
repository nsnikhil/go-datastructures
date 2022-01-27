package list

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/internal"
	"math"
	"testing"
)

func BenchmarkFinders(b *testing.B) {
	var lists []List[int64]

	for i := 1; i <= 8; i++ {
		sz := int64(math.Pow(10, float64(i)))

		al := NewArrayList(internal.SliceGenerator{Size: sz}.Generate()...)

		lists = append(lists, al)

	}

	lf := newLinearFinder[int64]()

	cf := newConcurrentFinder[int64]()

	for i, l := range lists {
		b.Run(fmt.Sprintf("lf %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				lf.search(l, math.MinInt32)
			}
		})
	}

	for i, l := range lists {
		b.Run(fmt.Sprintf("cf %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				cf.search(l, math.MinInt32)
			}
		})
	}
}
