package main

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

func graphOne() (graph, propertySet) {
	// ONE
	//
	//  6 <-- 4
	//   \   ^
	//	  \ /
	//     v
	//     5
	//

	four := newNode(4)
	five := newNode(5)
	six := newNode(6)

	g := newListGraph()
	g.createDiEdges(four, six)
	g.createDiEdges(six, five)
	g.createDiEdges(five, four)
	return g, newPropertySet(directed, unWeighted, cyclic, connected, stronglyConnected)
}

func graphTwo() (graph, propertySet) {
	// TWO
	//
	//  0  --> 1
	//    ^    |
	//      \  v
	//  3 <--> 2
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)

	g := newListGraph()
	g.createDiEdges(zero, one)
	g.createDiEdges(one, two)
	g.createDiEdges(two, zero)
	g.createBiEdges(two, three)
	return g, newPropertySet(directed, unWeighted, cyclic, connected, stronglyConnected)
}

func graphThree() (graph, propertySet) {
	// THREE
	//
	//  0  --> 1
	//    ^    |
	//      \  v
	//  3  --> 2
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)

	g := newListGraph()
	g.createDiEdges(zero, one)
	g.createDiEdges(one, two)
	g.createDiEdges(two, zero)
	g.createDiEdges(three, two)
	return g, newPropertySet(directed, unWeighted, cyclic, connected, weaklyConnected)
}

func graphFour() (graph, propertySet) {
	// FOUR
	//
	//  0  --> 1 <-- 2
	//  | \    | \   ^
	//  v   v  v   v |
	//  5      4 <-- 3
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	four := newNode(4)
	five := newNode(5)

	g := newListGraph()
	g.createDiEdges(zero, one, four, five)
	g.createDiEdges(one, three, four)
	g.createDiEdges(two, one)
	g.createDiEdges(three, two, four)
	return g, newPropertySet(directed, unWeighted, cyclic, connected, weaklyConnected)
}

func graphFive() (graph, propertySet) {
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

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	four := newNode(4)
	five := newNode(5)

	g := newListGraph()
	g.createDiEdges(zero, one)
	g.createDiEdges(one, two)
	g.createDiEdges(two, zero)
	g.createDiEdges(three, two, four)
	g.createDiEdges(four, five)
	g.createDiEdges(five, three)
	return g, newPropertySet(directed, unWeighted, cyclic, connected, weaklyConnected)
}

func graphSix() (graph, propertySet) {
	// SIX
	//
	//  6 --- 4
	//   \   /
	//	  \ /
	//     5
	//

	four := newNode(4)
	five := newNode(5)
	six := newNode(6)

	g := newListGraph()
	g.createBiEdges(four, six, five)
	g.createBiEdges(five, six)
	return g, newPropertySet(unDirected, unWeighted, cyclic, connected)
}

func graphSeven() (graph, propertySet) {
	// SEVEN
	//
	//  0  --- 1
	//    \    |
	//      \  |
	//  3  --- 2
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)

	g := newListGraph()
	g.createBiEdges(zero, one, two)
	g.createBiEdges(two, one, three)
	return g, newPropertySet(unDirected, unWeighted, cyclic, connected)
}

func graphEight() (graph, propertySet) {
	// EIGHT
	//
	//  0 --- 1 --- 2
	//  | \   | \   |
	//  |   \ |   \ |
	//  5     4 --- 3
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	four := newNode(4)
	five := newNode(5)

	g := newListGraph()
	g.createBiEdges(zero, one, four, five)
	g.createBiEdges(one, two, three, four)
	g.createBiEdges(three, two, four)
	return g, newPropertySet(unDirected, unWeighted, cyclic, connected)
}

func graphNine() (graph, propertySet) {
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

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	four := newNode(4)
	five := newNode(5)

	g := newListGraph()
	g.createBiEdges(zero, one, two)
	g.createBiEdges(one, two)
	g.createBiEdges(three, two, four, five)
	g.createBiEdges(four, five)
	return g, newPropertySet(unDirected, unWeighted, cyclic, connected)
}

func graphTen() (graph, propertySet) {
	// ONE
	//     2
	//  6 <-- 4
	//   \   ^
	// 3  \ /  4
	//     v
	//     5
	//

	four := newNode(4)
	five := newNode(5)
	six := newNode(6)

	g := newListGraph()
	g.createWeightedDiEdge(four, six, 2)
	g.createWeightedDiEdge(six, five, 3)
	g.createWeightedDiEdge(five, four, 4)
	return g, newPropertySet(directed, weighted, cyclic, connected, stronglyConnected)
}

func graphEleven() (graph, propertySet) {
	// TWO
	//
	//      2
	//  0  --> 1
	//    ^    | 4
	//    5 \  v
	//  3 <--> 2
	//      3
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)

	g := newListGraph()
	g.createWeightedDiEdge(zero, one, 2)
	g.createWeightedDiEdge(one, two, 4)
	g.createWeightedDiEdge(two, zero, 5)
	g.createWeightedBiEdge(two, three, 3)
	return g, newPropertySet(directed, weighted, cyclic, connected, stronglyConnected)
}

func graphTwelve() (graph, propertySet) {
	// THREE
	//
	//      2
	//  0  --> 1
	//    ^    | 4
	//    5 \  v
	//  3  --> 2
	//      3
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)

	g := newListGraph()
	g.createWeightedDiEdge(zero, one, 2)
	g.createWeightedDiEdge(one, two, 4)
	g.createWeightedDiEdge(two, zero, 5)
	g.createWeightedDiEdge(three, two, 3)
	return g, newPropertySet(directed, weighted, cyclic, connected, weaklyConnected)
}

func graphThirteen() (graph, propertySet) {
	// FOUR
	//       2      3
	//    0  --> 1 <-- 2
	// 4  | \ 5 2| \3  ^
	//    v   v  v   v | 1
	//    5      4 <-- 3
	//               4
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	four := newNode(4)
	five := newNode(5)

	g := newListGraph()
	g.createWeightedDiEdge(zero, one, 2)
	g.createWeightedDiEdge(zero, four, 5)
	g.createWeightedDiEdge(zero, five, 4)
	g.createWeightedDiEdge(one, four, 2)
	g.createWeightedDiEdge(one, three, 3)
	g.createWeightedDiEdge(two, one, 3)
	g.createWeightedDiEdge(three, two, 1)
	g.createWeightedDiEdge(three, four, 4)
	return g, newPropertySet(directed, weighted, cyclic, connected, weaklyConnected)
}

func graphFourteen() (graph, propertySet) {
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

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	four := newNode(4)
	five := newNode(5)

	g := newListGraph()
	g.createWeightedDiEdge(zero, one, 4)
	g.createWeightedDiEdge(one, two, 5)
	g.createWeightedDiEdge(two, zero, 3)
	g.createWeightedDiEdge(three, two, 2)
	g.createWeightedDiEdge(three, four, 4)
	g.createWeightedDiEdge(four, five, 5)
	g.createWeightedDiEdge(five, three, 6)
	return g, newPropertySet(directed, weighted, cyclic, connected, weaklyConnected)
}

func graphFifteen() (graph, propertySet) {
	// ONE
	//     2
	//  6 --- 4
	//   \   /
	// 3  \ /  4
	//     v
	//     5
	//

	four := newNode(4)
	five := newNode(5)
	six := newNode(6)

	g := newListGraph()
	g.createWeightedBiEdge(four, six, 2)
	g.createWeightedBiEdge(six, five, 3)
	g.createWeightedBiEdge(five, four, 4)
	return g, newPropertySet(unDirected, weighted, cyclic, connected)
}

func graphSixteen() (graph, propertySet) {
	// TWO
	//
	//      2
	//   0 --- 1
	//    \    | 4
	//    5 \  |
	//   3 --- 2
	//      3
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)

	g := newListGraph()
	g.createWeightedBiEdge(zero, one, 2)
	g.createWeightedBiEdge(one, two, 4)
	g.createWeightedBiEdge(two, zero, 5)
	g.createWeightedBiEdge(two, three, 3)
	return g, newPropertySet(unDirected, weighted, cyclic, connected)
}

func graphSeventeen() (graph, propertySet) {
	// FOUR
	//       2      3
	//    0  --- 1 --- 2
	// 4  | \ 5 2| \3  |
	//    |   \  |   \ | 1
	//    5      4 --- 3
	//               4
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	four := newNode(4)
	five := newNode(5)

	g := newListGraph()
	g.createWeightedBiEdge(zero, one, 2)
	g.createWeightedBiEdge(zero, four, 5)
	g.createWeightedBiEdge(zero, five, 4)
	g.createWeightedBiEdge(one, four, 2)
	g.createWeightedBiEdge(one, three, 3)
	g.createWeightedBiEdge(two, one, 3)
	g.createWeightedBiEdge(three, two, 1)
	g.createWeightedBiEdge(three, four, 4)
	return g, newPropertySet(unDirected, weighted, cyclic, connected)
}

func graphEighteen() (graph, propertySet) {
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

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	four := newNode(4)
	five := newNode(5)

	g := newListGraph()
	g.createWeightedBiEdge(zero, one, 4)
	g.createWeightedBiEdge(one, two, 5)
	g.createWeightedBiEdge(two, zero, 3)
	g.createWeightedBiEdge(three, two, 2)
	g.createWeightedBiEdge(three, four, 4)
	g.createWeightedBiEdge(four, five, 5)
	g.createWeightedBiEdge(five, three, 6)
	return g, newPropertySet(unDirected, weighted, cyclic, connected)
}

func graphNineTeen() (graph, propertySet) {
	// ONE
	//             2
	//          6 <-- 4
	//      1  / \ 3
	//        /   \
	//       v --> v
	//      7   1   5
	//
	//

	four := newNode(4)
	five := newNode(5)
	six := newNode(6)
	seven := newNode(7)

	g := newListGraph()
	g.createWeightedDiEdge(four, six, 2)
	g.createWeightedDiEdge(six, five, 3)
	g.createWeightedDiEdge(six, seven, 1)
	g.createWeightedDiEdge(seven, five, 1)
	return g, newPropertySet(directed, weighted, aCyclic, connected, weaklyConnected)
}

func graphTwenty() (graph, propertySet) {
	// TWO
	//
	//      2
	//   0 --> 1
	//     8 /   \ 4
	//      v     v
	//     3  <--  2
	//         3
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)

	g := newListGraph()
	g.createWeightedDiEdge(zero, one, 2)
	g.createWeightedDiEdge(one, two, 4)
	g.createWeightedDiEdge(one, three, 8)
	g.createWeightedBiEdge(two, three, 3)
	return g, newPropertySet(directed, weighted, aCyclic, connected, weaklyConnected)
}

func graphTwentyOne() (graph, propertySet) {
	// FOUR
	//       2      5
	//    0  --> 1 --> 2
	//  4 | \ 5 2| \3  ^
	//    v   v  v   v | 1
	//    5      4 <-- 3
	//               4
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	four := newNode(4)
	five := newNode(5)

	g := newListGraph()
	g.createWeightedDiEdge(zero, one, 2)
	g.createWeightedDiEdge(zero, four, 5)
	g.createWeightedDiEdge(zero, five, 4)
	g.createWeightedDiEdge(one, two, 5)
	g.createWeightedDiEdge(one, three, 3)
	g.createWeightedDiEdge(one, four, 2)
	g.createWeightedDiEdge(three, two, 1)
	g.createWeightedDiEdge(three, four, 4)
	return g, newPropertySet(directed, weighted, aCyclic, connected, weaklyConnected)
}

func graphTwentyTwo() (graph, propertySet) {
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

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	four := newNode(4)
	five := newNode(5)

	g := newListGraph()
	g.createWeightedDiEdge(zero, one, 4)
	g.createWeightedDiEdge(zero, three, 2)
	g.createWeightedDiEdge(one, two, 5)
	g.createWeightedDiEdge(three, two, 6)
	g.createWeightedDiEdge(three, four, 4)
	g.createWeightedDiEdge(four, five, 5)
	return g, newPropertySet(directed, weighted, aCyclic, connected, weaklyConnected)
}

func graphTwentyThree() (graph, propertySet) {
	// TWENTY THREE
	//
	//             -6
	//        ----------
	//       v           \
	//  0 --> 1 --> 2 --> 3 --> 4
	//     2     2     3     4
	//

	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)
	four := newNode(4)

	g := newListGraph()
	g.createWeightedDiEdge(zero, one, 2)
	g.createWeightedDiEdge(one, two, 2)
	g.createWeightedDiEdge(two, three, 3)
	g.createWeightedDiEdge(three, one, -6)
	g.createWeightedDiEdge(three, four, 4)
	return g, newPropertySet(directed, weighted, cyclic, connected, weaklyConnected, negativeWeights, negativeCycles)
}

func getGraphs(properties ...property) []graph {
	if len(properties) == 0 {
		return []graph{}
	}

	type graphSet struct {
		g  graph
		ps propertySet
	}

	toGS := func(g graph, ps propertySet) graphSet {
		return graphSet{
			g:  g,
			ps: ps,
		}
	}

	filter := func(gs []graphSet, p ...property) []graph {
		var res []graph

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

	gs := make([]graphSet, 0)
	gs = append(gs,
		toGS(graphOne()), toGS(graphTwo()), toGS(graphThree()), toGS(graphFour()),
		toGS(graphFive()), toGS(graphSix()), toGS(graphSeven()), toGS(graphEight()),
		toGS(graphNine()), toGS(graphTen()), toGS(graphEleven()), toGS(graphTwelve()),
		toGS(graphThirteen()), toGS(graphFourteen()), toGS(graphFifteen()), toGS(graphSixteen()),
		toGS(graphSeventeen()), toGS(graphEighteen()), toGS(graphNineTeen()), toGS(graphTwenty()),
		toGS(graphTwentyOne()), toGS(graphTwentyTwo()), toGS(graphTwentyThree()),
	)

	return filter(gs, properties...)
}
