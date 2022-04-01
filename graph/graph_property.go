package graph

type Property string

const (
	weighted   Property = "weighted"
	UnWeighted          = "UnWeighted"

	Directed   = "Directed"
	unDirected = "unDirected"

	cyclic  = "cyclic"
	ACyclic = "ACyclic"

	NonNegativeWeights = "NonNegativeWeights"
	negativeWeights    = "negativeWeights"
	negativeCycles     = "negativeCycles"

	connected    = "connected"
	disConnected = "disConnected"

	stronglyConnected = "stronglyConnected"
	weaklyConnected   = "weaklyConnected"
)
