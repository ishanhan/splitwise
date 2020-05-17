package simplify_debts

import (
	"math"
)

type ResidualValues struct {
	flow     float64
	backward bool
}

type Edge struct {
	User      string
	OtherUser string
}

func SimplifyDebts(nodes []Edge, residualGraph map[string]map[string]ResidualValues) {

	for _, node := range nodes {

		user := node.User
		otherUser := node.OtherUser

		if user == otherUser {
			continue
		}

		maxFlow := 0.0
		for {
			var path []string
			path = DFS(user, otherUser, residualGraph)

			if path == nil {
				break
			}

			minFlow := math.MaxFloat64
			for i := 0; i < len(path)-1; i++ {
				residualValues := residualGraph[path[i]][path[i+1]]
				minFlow = math.Min(minFlow, residualValues.flow)
			}

			for i := 0; i < len(path)-1; i++ {
				ModifyForwardEdge(residualGraph, path[i], path[i+1], minFlow)
				//ModifyBackwardEdge(residualGraph, path[i], path[i+1], minFlow)
			}

			maxFlow += minFlow

		}

		if maxFlow > 0.0 {

			AddForwardEdge(residualGraph, user, otherUser, maxFlow)
			//AddZeroBackwardEdge(residualGraph, user, otherUser)
		}
	}
}
