package simplify_detbs

import (
	"fmt"
	"math"
)

type ResidualValues struct {
	flow    float64
	visited bool
}

func CopyToResidualGraph(graph map[string]map[string]float64) map[string]map[string]ResidualValues {
	residualGraph := make(map[string]map[string]ResidualValues)
	for user, otherUser := range graph {

		if residualGraph[user] == nil {
			residualGraph[user] = make(map[string]ResidualValues)
		}

		for subKey, subValue := range otherUser {
			if subValue <= 0 {
				residualGraph[user][subKey] = ResidualValues{flow: -1 * subValue, visited: false}
			}

		}
	}
	return residualGraph
}

func MarkVisited(residualGraph map[string]map[string]ResidualValues, user, otherUser string) {
	residualInfo := residualGraph[user][otherUser]
	residualInfo.visited = true
	residualGraph[user][otherUser] = residualInfo
}

func AddBackwardEdge(residualGraph map[string]map[string]ResidualValues, user, otherUser string, minflow float64) {

	if residualGraph[otherUser] == nil {
		residualGraph[otherUser] = make(map[string]ResidualValues)
	}

	residualValue := residualGraph[otherUser][user]
	residualValue.flow += minflow
	residualGraph[otherUser][user] = residualValue
}

func AddForwardEdge(residualGraph map[string]map[string]ResidualValues, user string, otherUser string, maxFlow float64) {
	if residualGraph[user] == nil {
		residualGraph[user] = make(map[string]ResidualValues)
	}

	residualGraph[user][otherUser] = ResidualValues{flow: maxFlow, visited: false}
}

func RemoveBackwardEdge(residualGraph map[string]map[string]ResidualValues, user, otherUser string) {
	if residualGraph[otherUser] == nil {
		residualGraph[otherUser] = make(map[string]ResidualValues)
	}

	residualGraph[otherUser][user] = ResidualValues{flow: 0, visited: false}
}

func FordFulkerson(nodes []string, residualGraph map[string]map[string]ResidualValues) {

	for _, user := range nodes {
		for _, otherUser := range nodes {

			if user == otherUser {
				continue
			}

			visited := make(map[string]bool)

			maxFlow := 0.0
			for {
				path := make([]string, 0)
				if !residualGraph[user][otherUser].visited {
					DFS(user, otherUser, residualGraph, &path, visited)
					fmt.Printf("user %v and other user %v", user, otherUser)
				}

				if len(path) == 1 {
					break
				}
				minFlow := math.MaxFloat64
				for i := 0; i < len(path)-1; i++ {
					residualValues := residualGraph[path[i]][path[i+1]]
					minFlow = math.Min(minFlow, residualValues.flow)
				}

				for i := 0; i < len(path)-1; i++ {
					residualValues := residualGraph[path[i]][path[i+1]]
					residualValues.flow -= minFlow
					residualGraph[path[i]][path[i+1]] = residualValues
					AddBackwardEdge(residualGraph, path[i], path[i+1], minFlow)
				}

				maxFlow += minFlow
			}

			MarkVisited(residualGraph, user, otherUser)
			if maxFlow > 0.0 {
				AddForwardEdge(residualGraph, user, otherUser, maxFlow)
				RemoveBackwardEdge(residualGraph, user, otherUser)
			}

		}

	}
}

func DFS(user string, endUser string, residualGraph map[string]map[string]ResidualValues, path *[]string, visited map[string]bool) {
	*path = append(*path, user)
	if user == endUser {
		return
	}

	for otherUser, residualValue := range residualGraph[user] {

		if residualValue.flow > 0 {
			visited[otherUser] = true
			DFS(otherUser, endUser, residualGraph, path, visited)
			visited[otherUser] = false
			return
		}

	}

}
