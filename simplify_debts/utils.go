package simplify_debts

func CopyToResidualGraph(graph map[string]map[string]float64) map[string]map[string]ResidualValues {
	residualGraph := make(map[string]map[string]ResidualValues)
	for user, otherUser := range graph {

		if residualGraph[user] == nil {
			residualGraph[user] = make(map[string]ResidualValues)
		}

		for subKey, subValue := range otherUser {
			if subValue < 0 {
				residualGraph[user][subKey] = ResidualValues{flow: -1 * subValue}
			}

		}
	}
	return residualGraph
}

func ModifyForwardEdge(residualGraph map[string]map[string]ResidualValues, user, otherUser string, minflow float64) {
	if residualGraph[user] == nil {
		residualGraph[otherUser] = make(map[string]ResidualValues)
	}

	residualValue := residualGraph[user][otherUser]
	residualValue.flow -= minflow
	residualGraph[user][otherUser] = residualValue
}

func ModifyBackwardEdge(residualGraph map[string]map[string]ResidualValues, user, otherUser string, minflow float64) {

	if residualGraph[otherUser] == nil {
		residualGraph[otherUser] = make(map[string]ResidualValues)
	}

	residualValue := residualGraph[otherUser][user]
	if residualValue == (ResidualValues{}) {
		residualValue.backward = true
	}
	residualValue.flow += minflow
	residualGraph[otherUser][user] = residualValue
}

func AddForwardEdge(residualGraph map[string]map[string]ResidualValues, user string, otherUser string, maxFlow float64) {
	if residualGraph[user] == nil {
		residualGraph[user] = make(map[string]ResidualValues)
	}

	residualGraph[user][otherUser] = ResidualValues{flow: maxFlow}
}

func AddZeroBackwardEdge(residualGraph map[string]map[string]ResidualValues, user, otherUser string) {
	if residualGraph[otherUser] == nil {
		residualGraph[otherUser] = make(map[string]ResidualValues)
	}

	residualGraph[otherUser][user] = ResidualValues{flow: 0}
}
