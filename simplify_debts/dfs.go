package simplify_debts

func DFSUtils(user, endUser string, graph map[string]map[string]ResidualValues, path []string, visited map[string]bool) []string {
	visited[user] = true
	path = append(path, user)

	if user == endUser {
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		return pathCopy
	}

	for otherUser, residualValue := range graph[user] {
		if !visited[otherUser] && residualValue.flow > 0 {
			path := DFSUtils(otherUser, endUser, graph, path, visited)
			if path == nil {
				continue
			}
			return path
		}
	}

	return nil
}

func DFS(user, otherUser string, graph map[string]map[string]ResidualValues) []string {
	visited := make(map[string]bool)
	path := make([]string, 0)
	return DFSUtils(user, otherUser, graph, path, visited)
}
