package graph

type propertySet struct {
	hm map[Property]bool
}

func (ps propertySet) hasProperty(p Property) bool {
	return ps.hm[p]
}

func newPropertySet(properties ...Property) propertySet {
	data := make(map[Property]bool)

	for _, p := range properties {
		data[p] = true
	}

	return propertySet{
		hm: data,
	}
}

func graphOne() (Graph[int], propertySet) {
	// ONE
	//
	//  6 <-- 4
	//   \   ^
	//	  \ /
	//     v
	//     5
	//

	four := NewNode[int](4)
	five := NewNode[int](5)
	six := NewNode[int](6)

	g := NewListGraph[int]()
	createEdge(g, false, four, six)
	createEdge(g, false, six, five)
	createEdge(g, false, five, four)

	return g, newPropertySet(Directed, UnWeighted, cyclic, connected, stronglyConnected)
}

func graphOneReverse() (Graph[int], propertySet) {
	// ONE
	//
	//  6 --> 4
	//   ^   /
	//	  \ /
	//     v
	//     5
	//

	four := NewNode[int](4)
	five := NewNode[int](5)
	six := NewNode[int](6)

	g := NewListGraph[int]()
	createEdge(g, false, six, four)
	createEdge(g, false, five, six)
	createEdge(g, false, four, five)

	return g, newPropertySet(Directed, UnWeighted, cyclic, connected, stronglyConnected)
}

func graphTwo() (Graph[int], propertySet) {
	// TWO
	//
	//  0  --> 1
	//    ^    |
	//      \  v
	//  3 <--> 2
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)

	g := NewListGraph[int]()
	createEdge(g, false, zero, one)
	createEdge(g, false, one, two)
	createEdge(g, false, two, zero)
	createEdge(g, true, two, three)
	return g, newPropertySet(Directed, UnWeighted, cyclic, connected, stronglyConnected)
}

func graphThree() (Graph[int], propertySet) {
	// THREE
	//
	//  0  --> 1
	//    ^    |
	//      \  v
	//  3  --> 2
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)

	g := NewListGraph[int]()
	createEdge(g, false, zero, one)
	createEdge(g, false, one, two)
	createEdge(g, false, two, zero)
	createEdge(g, false, three, two)
	return g, newPropertySet(Directed, UnWeighted, cyclic, connected, weaklyConnected)
}

func graphFour() (Graph[int], propertySet) {
	// FOUR
	//
	//  0  --> 1 <-- 2
	//  | \    | \   ^
	//  v   v  v   v |
	//  5      4 <-- 3
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)
	four := NewNode[int](4)
	five := NewNode[int](5)

	g := NewListGraph[int]()
	createEdge(g, false, zero, one, four, five)
	createEdge(g, false, one, three, four)
	createEdge(g, false, two, one)
	createEdge(g, false, three, two, four)
	return g, newPropertySet(Directed, UnWeighted, cyclic, connected, weaklyConnected)
}

func graphFive() (Graph[int], propertySet) {
	// FIVE
	//
	//            0 --> 1
	//             ^    |
	//               \  v
	//    4 <-- 3  -->  2
	//    |   ^
	//    v  /
	//    5
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)
	four := NewNode[int](4)
	five := NewNode[int](5)

	g := NewListGraph[int]()
	createEdge(g, false, zero, one)
	createEdge(g, false, one, two)
	createEdge(g, false, two, zero)
	createEdge(g, false, three, two, four)
	createEdge(g, false, four, five)
	createEdge(g, false, five, three)
	return g, newPropertySet(Directed, UnWeighted, cyclic, connected, weaklyConnected)
}

func graphSix() (Graph[int], propertySet) {
	// SIX
	//
	//  6 --- 4
	//   \   /
	//	  \ /
	//     5
	//

	four := NewNode[int](4)
	five := NewNode[int](5)
	six := NewNode[int](6)

	g := NewListGraph[int]()
	createEdge(g, true, four, six, five)
	createEdge(g, true, five, six)
	return g, newPropertySet(unDirected, UnWeighted, cyclic, connected)
}

func graphSeven() (Graph[int], propertySet) {
	// SEVEN
	//
	//  0  --- 1
	//    \    |
	//      \  |
	//  3  --- 2
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)

	g := NewListGraph[int]()
	createEdge(g, true, zero, one, two)
	createEdge(g, true, two, one, three)
	return g, newPropertySet(unDirected, UnWeighted, cyclic, connected)
}

func graphEight() (Graph[int], propertySet) {
	// EIGHT
	//
	//  0 --- 1 --- 2
	//  | \   | \   |
	//  |   \ |   \ |
	//  5     4 --- 3
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)
	four := NewNode[int](4)
	five := NewNode[int](5)

	g := NewListGraph[int]()
	createEdge(g, true, zero, one, four, five)
	createEdge(g, true, one, two, three, four)
	createEdge(g, true, three, two, four)
	return g, newPropertySet(unDirected, UnWeighted, cyclic, connected)
}

func graphNine() (Graph[int], propertySet) {
	// NINE
	//
	//            0 --- 1
	//             \    |
	//               \  |
	//    4 --- 3  ---  2
	//    |   /
	//    |  /
	//     5
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)
	four := NewNode[int](4)
	five := NewNode[int](5)

	g := NewListGraph[int]()
	createEdge(g, true, zero, one, two)
	createEdge(g, true, one, two)
	createEdge(g, true, three, two, four, five)
	createEdge(g, true, four, five)
	return g, newPropertySet(unDirected, UnWeighted, cyclic, connected)
}

func graphTen() (Graph[int], propertySet) {
	// ONE
	//     2
	//  6 <-- 4
	//   \   ^
	// 3  \ /  4
	//     v
	//     5
	//

	four := NewNode[int](4)
	five := NewNode[int](5)
	six := NewNode[int](6)

	g := NewListGraph[int]()
	createWeightedEdge(g, false, four, six, 2)
	createWeightedEdge(g, false, six, five, 3)
	createWeightedEdge(g, false, five, four, 4)
	return g, newPropertySet(Directed, weighted, cyclic, connected, stronglyConnected)
}

func graphEleven() (Graph[int], propertySet) {
	// TWO
	//
	//      2
	//  0  --> 1
	//    ^    | 4
	//    5 \  v
	//  3 <--> 2
	//      3
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)

	g := NewListGraph[int]()
	createWeightedEdge(g, false, zero, one, 2)
	createWeightedEdge(g, false, one, two, 4)
	createWeightedEdge(g, false, two, zero, 5)
	createWeightedEdge(g, true, two, three, 3)
	return g, newPropertySet(Directed, weighted, cyclic, connected, stronglyConnected)
}

func graphTwelve() (Graph[int], propertySet) {
	// THREE
	//
	//      2
	//  0  --> 1
	//    ^    | 4
	//    5 \  v
	//  3  --> 2
	//      3
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)

	g := NewListGraph[int]()
	createWeightedEdge(g, false, zero, one, 2)
	createWeightedEdge(g, false, one, two, 4)
	createWeightedEdge(g, false, two, zero, 5)
	createWeightedEdge(g, false, three, two, 3)
	return g, newPropertySet(Directed, weighted, cyclic, connected, weaklyConnected)
}

func graphThirteen() (Graph[int], propertySet) {
	// FOUR
	//       2      3
	//    0  --> 1 <-- 2
	// 4  | \ 5 2| \3  ^
	//    v   v  v   v | 1
	//    5      4 <-- 3
	//               4
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)
	four := NewNode[int](4)
	five := NewNode[int](5)

	g := NewListGraph[int]()
	createWeightedEdge(g, false, zero, one, 2)
	createWeightedEdge(g, false, zero, four, 5)
	createWeightedEdge(g, false, zero, five, 4)
	createWeightedEdge(g, false, one, four, 2)
	createWeightedEdge(g, false, one, three, 3)
	createWeightedEdge(g, false, two, one, 3)
	createWeightedEdge(g, false, three, two, 1)
	createWeightedEdge(g, false, three, four, 4)
	return g, newPropertySet(Directed, weighted, cyclic, connected, weaklyConnected)
}

func graphFourteen() (Graph[int], propertySet) {
	// FIVE
	//               4
	//            0 --> 1
	//          3  ^    | 5
	//       4       \  v
	//    4 <-- 3  -->  2
	//  5 |   ^    2
	//    v  /  6
	//    5
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)
	four := NewNode[int](4)
	five := NewNode[int](5)

	g := NewListGraph[int]()
	createWeightedEdge(g, false, zero, one, 4)
	createWeightedEdge(g, false, one, two, 5)
	createWeightedEdge(g, false, two, zero, 3)
	createWeightedEdge(g, false, three, two, 2)
	createWeightedEdge(g, false, three, four, 4)
	createWeightedEdge(g, false, four, five, 5)
	createWeightedEdge(g, false, five, three, 6)
	return g, newPropertySet(Directed, weighted, cyclic, connected, weaklyConnected)
}

func graphFifteen() (Graph[int], propertySet) {
	// ONE
	//     2
	//  6 --- 4
	//   \   /
	// 3  \ /  4
	//     v
	//     5
	//

	four := NewNode[int](4)
	five := NewNode[int](5)
	six := NewNode[int](6)

	g := NewListGraph[int]()
	createWeightedEdge(g, true, four, six, 2)
	createWeightedEdge(g, true, six, five, 3)
	createWeightedEdge(g, true, five, four, 4)
	return g, newPropertySet(unDirected, weighted, cyclic, connected)
}

func graphSixteen() (Graph[int], propertySet) {
	// TWO
	//
	//      2
	//   0 --- 1
	//    \    | 4
	//    5 \  |
	//   3 --- 2
	//      3
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)

	g := NewListGraph[int]()
	createWeightedEdge(g, true, zero, one, 2)
	createWeightedEdge(g, true, one, two, 4)
	createWeightedEdge(g, true, two, zero, 5)
	createWeightedEdge(g, true, two, three, 3)
	return g, newPropertySet(unDirected, weighted, cyclic, connected)
}

func graphSeventeen() (Graph[int], propertySet) {
	// FOUR
	//       2      3
	//    0  --- 1 --- 2
	// 4  | \ 5 2| \3  |
	//    |   \  |   \ | 1
	//    5      4 --- 3
	//               4
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)
	four := NewNode[int](4)
	five := NewNode[int](5)

	g := NewListGraph[int]()
	createWeightedEdge(g, true, zero, one, 2)
	createWeightedEdge(g, true, zero, four, 5)
	createWeightedEdge(g, true, zero, five, 4)
	createWeightedEdge(g, true, one, four, 2)
	createWeightedEdge(g, true, one, three, 3)
	createWeightedEdge(g, true, two, one, 3)
	createWeightedEdge(g, true, three, two, 1)
	createWeightedEdge(g, true, three, four, 4)
	return g, newPropertySet(unDirected, weighted, cyclic, connected)
}

func graphEighteen() (Graph[int], propertySet) {
	// FIVE
	//               4
	//            0 --- 1
	//          3  \    | 5
	//       4       \  |
	//    4 --- 3  ---  2
	//  5 |   /    2
	//    |  /  6
	//    5
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)
	four := NewNode[int](4)
	five := NewNode[int](5)

	g := NewListGraph[int]()
	createWeightedEdge(g, true, zero, one, 4)
	createWeightedEdge(g, true, one, two, 5)
	createWeightedEdge(g, true, two, zero, 3)
	createWeightedEdge(g, true, three, two, 2)
	createWeightedEdge(g, true, three, four, 4)
	createWeightedEdge(g, true, four, five, 5)
	createWeightedEdge(g, true, five, three, 6)
	return g, newPropertySet(unDirected, weighted, cyclic, connected)
}

func graphNineTeen() (Graph[int], propertySet) {
	// ONE
	//             2
	//          6 <-- 4
	//      1  / \ 3
	//        /   \
	//       v --> v
	//      7   1   5
	//
	//

	four := NewNode[int](4)
	five := NewNode[int](5)
	six := NewNode[int](6)
	seven := NewNode[int](7)

	g := NewListGraph[int]()
	createWeightedEdge(g, false, four, six, 2)
	createWeightedEdge(g, false, six, five, 3)
	createWeightedEdge(g, false, six, seven, 1)
	createWeightedEdge(g, false, seven, five, 1)
	return g, newPropertySet(Directed, weighted, ACyclic, connected, weaklyConnected)
}

func graphTwenty() (Graph[int], propertySet) {
	// TWO
	//
	//      2
	//   0 --> 1
	//     8 /   \ 4
	//      v     v
	//     3  <--  2
	//         3
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)

	g := NewListGraph[int]()
	createWeightedEdge(g, false, zero, one, 2)
	createWeightedEdge(g, false, one, two, 4)
	createWeightedEdge(g, false, one, three, 8)
	createWeightedEdge(g, false, two, three, 3)
	return g, newPropertySet(Directed, weighted, ACyclic, connected, weaklyConnected)
}

func graphTwentyOne() (Graph[int], propertySet) {
	// FOUR
	//       2      5
	//    0  --> 1 --> 2
	//  4 | \ 5 2| \3  ^
	//    v   v  v   v | 1
	//    5      4 <-- 3
	//               4
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)
	four := NewNode[int](4)
	five := NewNode[int](5)

	g := NewListGraph[int]()
	createWeightedEdge(g, false, zero, one, 2)
	createWeightedEdge(g, false, zero, four, 5)
	createWeightedEdge(g, false, zero, five, 4)
	createWeightedEdge(g, false, one, two, 5)
	createWeightedEdge(g, false, one, three, 3)
	createWeightedEdge(g, false, one, four, 2)
	createWeightedEdge(g, false, three, two, 1)
	createWeightedEdge(g, false, three, four, 4)
	return g, newPropertySet(Directed, weighted, ACyclic, connected, weaklyConnected)
}

func graphTwentyTwo() (Graph[int], propertySet) {
	// FIVE
	//               4
	//            0 --> 1
	//         2 /      | 5
	//          v       v
	//    4 <-- 3  -->  2
	//  5 |  4     6
	//    v
	//    5
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)
	four := NewNode[int](4)
	five := NewNode[int](5)

	g := NewListGraph[int]()
	createWeightedEdge(g, false, zero, one, 4)
	createWeightedEdge(g, false, zero, three, 2)
	createWeightedEdge(g, false, one, two, 5)
	createWeightedEdge(g, false, three, two, 6)
	createWeightedEdge(g, false, three, four, 4)
	createWeightedEdge(g, false, four, five, 5)
	return g, newPropertySet(Directed, weighted, ACyclic, connected, weaklyConnected)
}

func graphTwentyThree() (Graph[int], propertySet) {
	// TWENTY THREE
	//
	//             -6
	//        ----------
	//       v           \
	//  0 --> 1 --> 2 --> 3 --> 4
	//     2     2     3     4
	//

	zero := NewNode[int](0)
	one := NewNode[int](1)
	two := NewNode[int](2)
	three := NewNode[int](3)
	four := NewNode[int](4)

	g := NewListGraph[int]()
	createWeightedEdge(g, false, zero, one, 2)
	createWeightedEdge(g, false, one, two, 2)
	createWeightedEdge(g, false, two, three, 3)
	createWeightedEdge(g, false, three, one, -6)
	createWeightedEdge(g, false, three, four, 4)
	return g, newPropertySet(Directed, weighted, cyclic, connected, weaklyConnected, negativeWeights, negativeCycles)
}

func graphTwentyFour() (Graph[int], propertySet) {
	// TWENTY FOUR
	//
	// 0 ---- 1 ---- 2 ---- 3      8 ---- 9      10 ---- 11
	//        |      |  \                        | \      |
	//        |      |     \                     |    \   |
	//        5 ---- 4      6                    12      13
	//                \     |                    |     /
	//                   \  |                    |   /
	//                      7                    14
	//
	//

	g := NewListGraph[int]()

	sz := 15
	nodes := make([]*Node[int], sz)
	for i := 0; i < sz; i++ {
		nodes[i] = NewNode[int](i)
		g.AddNode(nodes[i])
	}

	createEdge(g, true, nodes[0], nodes[1])
	createEdge(g, true, nodes[1], nodes[2])
	createEdge(g, true, nodes[1], nodes[5])
	createEdge(g, true, nodes[2], nodes[3])
	createEdge(g, true, nodes[2], nodes[4])
	createEdge(g, true, nodes[2], nodes[6])
	createEdge(g, true, nodes[4], nodes[5])
	createEdge(g, true, nodes[4], nodes[7])
	createEdge(g, true, nodes[6], nodes[7])

	createEdge(g, true, nodes[8], nodes[9])

	createEdge(g, true, nodes[10], nodes[11])
	createEdge(g, true, nodes[10], nodes[12])
	createEdge(g, true, nodes[10], nodes[13])
	createEdge(g, true, nodes[11], nodes[13])
	createEdge(g, true, nodes[12], nodes[14])
	createEdge(g, true, nodes[13], nodes[14])

	return g, newPropertySet(unDirected, UnWeighted, cyclic, disConnected)
}

func graphTwentyFive() (Graph[int], propertySet) {
	// TWENTY FOUR
	//
	// 0 ---> 1 ---> 2 ---> 3      8 ---> 9      10 <--- 11
	//        |      ^   \                       ^ \      ^
	//        v      |     v                     |    v   |
	//        5 <--- 4      6                    12      13
	//                ^     |                    ^     /
	//                   \  v                    |   v
	//                      7                    14
	//
	//

	g := NewListGraph[int]()

	sz := 15
	nodes := make([]*Node[int], sz)
	for i := 0; i < sz; i++ {
		nodes[i] = NewNode[int](i)
		g.AddNode(nodes[i])
	}

	createEdge(g, false, nodes[0], nodes[1])
	createEdge(g, false, nodes[1], nodes[2])
	createEdge(g, false, nodes[1], nodes[5])
	createEdge(g, false, nodes[2], nodes[3])
	createEdge(g, false, nodes[2], nodes[6])
	createEdge(g, false, nodes[4], nodes[2])
	createEdge(g, false, nodes[4], nodes[5])
	createEdge(g, false, nodes[6], nodes[7])
	createEdge(g, false, nodes[7], nodes[4])

	createEdge(g, false, nodes[8], nodes[9])

	createEdge(g, false, nodes[10], nodes[13])
	createEdge(g, false, nodes[11], nodes[10])
	createEdge(g, false, nodes[12], nodes[10])
	createEdge(g, false, nodes[13], nodes[11])
	createEdge(g, false, nodes[13], nodes[14])
	createEdge(g, false, nodes[14], nodes[12])

	return g, newPropertySet(Directed, UnWeighted, cyclic, disConnected)
}

func graphTwentyFiveReverse() (Graph[int], propertySet) {
	// TWENTY FOUR
	//
	// 0 <--- 1 <--- 2 <--- 3      8 <--- 9      10 ---> 11
	//        ^      |   ^                       | ^      |
	//        |      v     \                     v    \   v
	//        5 ---> 4      6                    12      13
	//                \     ^                    |     ^
	//                   v  |                    v   /
	//                      7                    14
	//
	//

	g := NewListGraph[int]()

	sz := 15
	nodes := make([]*Node[int], sz)
	for i := 0; i < sz; i++ {
		nodes[i] = NewNode[int](i)
		g.AddNode(nodes[i])
	}

	createEdge(g, false, nodes[1], nodes[0])
	createEdge(g, false, nodes[2], nodes[1])
	createEdge(g, false, nodes[5], nodes[1])
	createEdge(g, false, nodes[3], nodes[2])
	createEdge(g, false, nodes[6], nodes[2])
	createEdge(g, false, nodes[2], nodes[4])
	createEdge(g, false, nodes[5], nodes[4])
	createEdge(g, false, nodes[7], nodes[6])
	createEdge(g, false, nodes[4], nodes[7])

	createEdge(g, false, nodes[9], nodes[8])

	createEdge(g, false, nodes[13], nodes[10])
	createEdge(g, false, nodes[10], nodes[11])
	createEdge(g, false, nodes[10], nodes[12])
	createEdge(g, false, nodes[11], nodes[13])
	createEdge(g, false, nodes[14], nodes[13])
	createEdge(g, false, nodes[12], nodes[14])

	return g, newPropertySet(Directed, UnWeighted, cyclic, disConnected)
}

type graphSet[T comparable] struct {
	g  Graph[T]
	ps propertySet
}

func getGraphs[T comparable](properties ...Property) []Graph[int] {
	if len(properties) == 0 {
		return []Graph[int]{}
	}

	toGS := func(g Graph[int], ps propertySet) graphSet[int] {
		return graphSet[int]{
			g:  g,
			ps: ps,
		}
	}

	filter := func(gs []graphSet[int], p ...Property) []Graph[int] {
		var res []Graph[int]

		for _, g := range gs {

			isValid := true
			for _, v := range p {
				if !g.ps.hasProperty(v) {
					isValid = false
					break
				}
			}

			if isValid {
				res = append(res, g.g)
			}
		}

		return res
	}

	gs := make([]graphSet[int], 0)
	gs = append(gs,
		toGS(graphOne()), toGS(graphOneReverse()), toGS(graphTwo()), toGS(graphThree()), toGS(graphFour()),
		toGS(graphFive()), toGS(graphSix()), toGS(graphSeven()), toGS(graphEight()),
		toGS(graphNine()), toGS(graphTen()), toGS(graphEleven()), toGS(graphTwelve()),
		toGS(graphThirteen()), toGS(graphFourteen()), toGS(graphFifteen()), toGS(graphSixteen()),
		toGS(graphSeventeen()), toGS(graphEighteen()), toGS(graphNineTeen()), toGS(graphTwenty()),
		toGS(graphTwentyOne()), toGS(graphTwentyTwo()), toGS(graphTwentyThree()), toGS(graphTwentyFour()),
		toGS(graphTwentyFive()), toGS(graphTwentyFiveReverse()),
	)

	return filter(gs, properties...)
}

func getAllGraphs() []Graph[int] {
	f := func(g Graph[int], ps propertySet) Graph[int] { return g }

	return []Graph[int]{
		f(graphOne()), f(graphOneReverse()), f(graphTwo()), f(graphThree()), f(graphFour()),
		f(graphFive()), f(graphSix()), f(graphSeven()), f(graphEight()),
		f(graphNine()), f(graphTen()), f(graphEleven()), f(graphTwelve()),
		f(graphThirteen()), f(graphFourteen()), f(graphFifteen()), f(graphSixteen()),
		f(graphSeventeen()), f(graphEighteen()), f(graphNineTeen()), f(graphTwenty()),
		f(graphTwentyOne()), f(graphTwentyTwo()), f(graphTwentyThree()), f(graphTwentyFour()),
		f(graphTwentyFive()), f(graphTwentyFiveReverse()),
	}
}

func createEdge(g Graph[int], isBidirected bool, source *Node[int], targets ...*Node[int]) {
	for _, target := range targets {
		createWeightedEdge(g, isBidirected, source, target, 0)
	}
}

func createWeightedEdge(g Graph[int], isBidirected bool, source *Node[int], target *Node[int], weight int64) {
	if !g.Contains(source) {
		g.AddNode(source)
	}

	if !g.Contains(target) {
		g.AddNode(target)
	}

	var err error

	if isBidirected {
		err = g.CreateWeightedBiEdge(source, target, weight)
	} else {
		err = g.CreateWeightedDiEdge(source, target, weight)
	}

	if err != nil {
		panic(err)
	}
}
