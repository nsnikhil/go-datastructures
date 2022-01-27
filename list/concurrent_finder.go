package list

func searchUtil[T comparable](l List[T], e interface{}, start, end int64, res chan<- int64) {
	for start <= end {
		el, _ := l.Get(start)
		if el == e {
			res <- start
		}

		el, _ = l.Get(end)
		if el == e {
			res <- end
		}
		start++
		end--
	}
	res <- -1
}

type concurrentFinder[T comparable] struct {
	searchFactor  int64
	partitionSize int64
}

func newConcurrentFinder[T comparable]() finder[T] {
	return concurrentFinder[T]{
		searchFactor:  10000,
		partitionSize: 1000,
	}
}

/*
	concurrent finder splits the list into partitions and searches in each partition in a new go routine.

	concurrent finder is only invoked if the number of elements in list are greater than search factor else
	the linear finder is invoked

	below are the benchmark comparing linear and concurrent finder:

	linear finder:
		element count      iterations      iterations/sec
		10         	        3288336	        368 ns/op
		100         	    856701	        1376 ns/op
		1000         	    108962	        10903 ns/op
		10000         	    10000	        107407 ns/op
		100000         	    1130	        1042636 ns/op
		1000000         	100	            10808345 ns/op
		10000000         	10	            110666584 ns/op
		100000000         	1	            1074130085 ns/op

	concurrent finder:
		element count      iterations      iterations/sec
		10         	        2913409	        415 ns/op
		100         	    881508	        1337 ns/op
		1000         	    113041	        10654 ns/op
		10000         	    26298	        44269 ns/op
		100000         	    3853	        303841 ns/op
		1000000        	    435	            2731591 ns/op
		10000000         	40	            26074049 ns/op
		100000000         	4	            280420052 ns/op

*/
func (cf concurrentFinder[T]) search(l List[T], e T) int64 {
	sz := l.Size()
	if sz == 0 {
		return -1
	}

	if sz < cf.searchFactor {
		return newLinearFinder[T]().search(l, e)
	}

	res := make(chan int64, cf.partitionSize)

	for i := int64(0); i < sz; i += cf.partitionSize {
		go searchUtil(l, e, i, i+cf.partitionSize, res)
	}

	for i := int64(0); i < sz; i += cf.partitionSize {
		if idx := <-res; idx != -1 {
			return idx
		}
	}

	return -1
}
