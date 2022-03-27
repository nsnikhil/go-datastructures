package graph

type property string

const (
	weighted   property = "weighted"
	unWeighted          = "unWeighted"

	directed   = "directed"
	unDirected = "unDirected"

	cyclic  = "cyclic"
	aCyclic = "aCyclic"

	negativeWeights = "negativeWeights"
	negativeCycles  = "negativeCycles"

	connected    = "connected"
	disConnected = "disConnected"

	stronglyConnected = "stronglyConnected"
	weaklyConnected   = "weaklyConnected"
)

type propertySet struct {
	hm map[property]bool
}

func (ps propertySet) hasProperty(p property) bool {
	return ps.hm[p]
}

func newPropertySet(properties ...property) propertySet {
	data := make(map[property]bool)

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

	return g, newPropertySet(directed, unWeighted, cyclic, connected, stronglyConnected)
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

	return g, newPropertySet(directed, unWeighted, cyclic, connected, stronglyConnected)
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
	return g, newPropertySet(directed, unWeighted, cyclic, connected, stronglyConnected)
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
	return g, newPropertySet(directed, unWeighted, cyclic, connected, weaklyConnected)
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
	return g, newPropertySet(directed, unWeighted, cyclic, connected, weaklyConnected)
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
	return g, newPropertySet(directed, unWeighted, cyclic, connected, weaklyConnected)
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
	return g, newPropertySet(unDirected, unWeighted, cyclic, connected)
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
	return g, newPropertySet(unDirected, unWeighted, cyclic, connected)
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
	return g, newPropertySet(unDirected, unWeighted, cyclic, connected)
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
	return g, newPropertySet(unDirected, unWeighted, cyclic, connected)
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
	g.CreateWeightedDiEdge(four, six, 2)
	g.CreateWeightedDiEdge(six, five, 3)
	g.CreateWeightedDiEdge(five, four, 4)
	return g, newPropertySet(directed, weighted, cyclic, connected, stronglyConnected)
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
	g.CreateWeightedDiEdge(zero, one, 2)
	g.CreateWeightedDiEdge(one, two, 4)
	g.CreateWeightedDiEdge(two, zero, 5)
	g.CreateWeightedBiEdge(two, three, 3)
	return g, newPropertySet(directed, weighted, cyclic, connected, stronglyConnected)
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
	g.CreateWeightedDiEdge(zero, one, 2)
	g.CreateWeightedDiEdge(one, two, 4)
	g.CreateWeightedDiEdge(two, zero, 5)
	g.CreateWeightedDiEdge(three, two, 3)
	return g, newPropertySet(directed, weighted, cyclic, connected, weaklyConnected)
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
	g.CreateWeightedDiEdge(zero, one, 2)
	g.CreateWeightedDiEdge(zero, four, 5)
	g.CreateWeightedDiEdge(zero, five, 4)
	g.CreateWeightedDiEdge(one, four, 2)
	g.CreateWeightedDiEdge(one, three, 3)
	g.CreateWeightedDiEdge(two, one, 3)
	g.CreateWeightedDiEdge(three, two, 1)
	g.CreateWeightedDiEdge(three, four, 4)
	return g, newPropertySet(directed, weighted, cyclic, connected, weaklyConnected)
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
	g.CreateWeightedDiEdge(zero, one, 4)
	g.CreateWeightedDiEdge(one, two, 5)
	g.CreateWeightedDiEdge(two, zero, 3)
	g.CreateWeightedDiEdge(three, two, 2)
	g.CreateWeightedDiEdge(three, four, 4)
	g.CreateWeightedDiEdge(four, five, 5)
	g.CreateWeightedDiEdge(five, three, 6)
	return g, newPropertySet(directed, weighted, cyclic, connected, weaklyConnected)
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
	g.CreateWeightedBiEdge(four, six, 2)
	g.CreateWeightedBiEdge(six, five, 3)
	g.CreateWeightedBiEdge(five, four, 4)
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
	g.CreateWeightedBiEdge(zero, one, 2)
	g.CreateWeightedBiEdge(one, two, 4)
	g.CreateWeightedBiEdge(two, zero, 5)
	g.CreateWeightedBiEdge(two, three, 3)
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
	g.CreateWeightedBiEdge(zero, one, 2)
	g.CreateWeightedBiEdge(zero, four, 5)
	g.CreateWeightedBiEdge(zero, five, 4)
	g.CreateWeightedBiEdge(one, four, 2)
	g.CreateWeightedBiEdge(one, three, 3)
	g.CreateWeightedBiEdge(two, one, 3)
	g.CreateWeightedBiEdge(three, two, 1)
	g.CreateWeightedBiEdge(three, four, 4)
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
	g.CreateWeightedBiEdge(zero, one, 4)
	g.CreateWeightedBiEdge(one, two, 5)
	g.CreateWeightedBiEdge(two, zero, 3)
	g.CreateWeightedBiEdge(three, two, 2)
	g.CreateWeightedBiEdge(three, four, 4)
	g.CreateWeightedBiEdge(four, five, 5)
	g.CreateWeightedBiEdge(five, three, 6)
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
	g.CreateWeightedDiEdge(four, six, 2)
	g.CreateWeightedDiEdge(six, five, 3)
	g.CreateWeightedDiEdge(six, seven, 1)
	g.CreateWeightedDiEdge(seven, five, 1)
	return g, newPropertySet(directed, weighted, aCyclic, connected, weaklyConnected)
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
	g.CreateWeightedDiEdge(zero, one, 2)
	g.CreateWeightedDiEdge(one, two, 4)
	g.CreateWeightedDiEdge(one, three, 8)
	g.CreateWeightedBiEdge(two, three, 3)
	return g, newPropertySet(directed, weighted, aCyclic, connected, weaklyConnected)
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
	g.CreateWeightedDiEdge(zero, one, 2)
	g.CreateWeightedDiEdge(zero, four, 5)
	g.CreateWeightedDiEdge(zero, five, 4)
	g.CreateWeightedDiEdge(one, two, 5)
	g.CreateWeightedDiEdge(one, three, 3)
	g.CreateWeightedDiEdge(one, four, 2)
	g.CreateWeightedDiEdge(three, two, 1)
	g.CreateWeightedDiEdge(three, four, 4)
	return g, newPropertySet(directed, weighted, aCyclic, connected, weaklyConnected)
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
	g.CreateWeightedDiEdge(zero, one, 4)
	g.CreateWeightedDiEdge(zero, three, 2)
	g.CreateWeightedDiEdge(one, two, 5)
	g.CreateWeightedDiEdge(three, two, 6)
	g.CreateWeightedDiEdge(three, four, 4)
	g.CreateWeightedDiEdge(four, five, 5)
	return g, newPropertySet(directed, weighted, aCyclic, connected, weaklyConnected)
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
	g.CreateWeightedDiEdge(zero, one, 2)
	g.CreateWeightedDiEdge(one, two, 2)
	g.CreateWeightedDiEdge(two, three, 3)
	g.CreateWeightedDiEdge(three, one, -6)
	g.CreateWeightedDiEdge(three, four, 4)
	return g, newPropertySet(directed, weighted, cyclic, connected, weaklyConnected, negativeWeights, negativeCycles)
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

	g.CreateBiEdge(nodes[0], nodes[1])
	g.CreateBiEdge(nodes[1], nodes[2])
	g.CreateBiEdge(nodes[1], nodes[5])
	g.CreateBiEdge(nodes[2], nodes[3])
	g.CreateBiEdge(nodes[2], nodes[4])
	g.CreateBiEdge(nodes[2], nodes[6])
	g.CreateBiEdge(nodes[4], nodes[5])
	g.CreateBiEdge(nodes[4], nodes[7])
	g.CreateBiEdge(nodes[6], nodes[7])

	g.CreateBiEdge(nodes[8], nodes[9])

	g.CreateBiEdge(nodes[10], nodes[11])
	g.CreateBiEdge(nodes[10], nodes[12])
	g.CreateBiEdge(nodes[10], nodes[13])
	g.CreateBiEdge(nodes[11], nodes[13])
	g.CreateBiEdge(nodes[12], nodes[14])
	g.CreateBiEdge(nodes[13], nodes[14])

	return g, newPropertySet(unDirected, unWeighted, cyclic, disConnected)
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

	g.CreateDiEdge(nodes[0], nodes[1])
	g.CreateDiEdge(nodes[1], nodes[2])
	g.CreateDiEdge(nodes[1], nodes[5])
	g.CreateDiEdge(nodes[2], nodes[3])
	g.CreateDiEdge(nodes[2], nodes[6])
	g.CreateDiEdge(nodes[4], nodes[2])
	g.CreateDiEdge(nodes[4], nodes[5])
	g.CreateDiEdge(nodes[6], nodes[7])
	g.CreateDiEdge(nodes[7], nodes[4])

	g.CreateDiEdge(nodes[8], nodes[9])

	g.CreateDiEdge(nodes[10], nodes[13])
	g.CreateDiEdge(nodes[11], nodes[10])
	g.CreateDiEdge(nodes[12], nodes[10])
	g.CreateDiEdge(nodes[13], nodes[11])
	g.CreateDiEdge(nodes[13], nodes[14])
	g.CreateDiEdge(nodes[14], nodes[12])

	return g, newPropertySet(directed, unWeighted, cyclic, disConnected)
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

	g.CreateDiEdge(nodes[1], nodes[0])
	g.CreateDiEdge(nodes[2], nodes[1])
	g.CreateDiEdge(nodes[5], nodes[1])
	g.CreateDiEdge(nodes[3], nodes[2])
	g.CreateDiEdge(nodes[6], nodes[2])
	g.CreateDiEdge(nodes[2], nodes[4])
	g.CreateDiEdge(nodes[5], nodes[4])
	g.CreateDiEdge(nodes[7], nodes[6])
	g.CreateDiEdge(nodes[4], nodes[7])

	g.CreateDiEdge(nodes[9], nodes[8])

	g.CreateDiEdge(nodes[13], nodes[10])
	g.CreateDiEdge(nodes[10], nodes[11])
	g.CreateDiEdge(nodes[10], nodes[12])
	g.CreateDiEdge(nodes[11], nodes[13])
	g.CreateDiEdge(nodes[14], nodes[13])
	g.CreateDiEdge(nodes[12], nodes[14])

	return g, newPropertySet(directed, unWeighted, cyclic, disConnected)
}

type graphSet[T comparable] struct {
	g  Graph[T]
	ps propertySet
}

func getGraphs[T comparable](properties ...property) []Graph[int] {
	if len(properties) == 0 {
		return []Graph[int]{}
	}

	toGS := func(g Graph[int], ps propertySet) graphSet[int] {
		return graphSet[int]{
			g:  g,
			ps: ps,
		}
	}

	filter := func(gs []graphSet[int], p ...property) []Graph[int] {
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

func createEdge(g Graph[int], isBidirected bool, source *Node[int], targets ...*Node[int]) {
	for _, target := range targets {
		if isBidirected {
			g.CreateBiEdge(source, target)
		} else {
			g.CreateDiEdge(source, target)
		}
	}
}
