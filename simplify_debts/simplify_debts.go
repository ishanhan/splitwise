package simplify_debts

import (
	"log"
	"math"
)

type ResidualValues struct {
	flow    float64
	visited bool
}

func SimplifyDebts(nodes []string, residualGraph map[string]map[string]ResidualValues) {

	for _, user := range nodes {
		for _, otherUser := range nodes {

			if user == otherUser {
				continue
			}

			maxFlow := 0.0
			for {
				var path []string
				if !residualGraph[user][otherUser].visited {
					path = DFS(user, otherUser, residualGraph)
					log.Printf("user %v and other user %v", user, otherUser)
				}

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
					ModifyBackwardEdge(residualGraph, path[i], path[i+1], minFlow)
				}

				maxFlow += minFlow
			}

			MarkVisited(residualGraph, user, otherUser)
			if maxFlow > 0.0 {
				AddForwardEdge(residualGraph, user, otherUser, maxFlow)
				AddZeroBackwardEdge(residualGraph, user, otherUser)
			}

		}

	}
}
