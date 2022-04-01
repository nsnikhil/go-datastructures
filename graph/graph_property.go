package graph

type Property string

const (
	weighted   Property = "weighted"
	UnWeighted          = "UnWeighted"

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
