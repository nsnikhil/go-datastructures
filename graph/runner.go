package main

func main() {
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

	dagShortestPath(g.(*listGraph))
}
