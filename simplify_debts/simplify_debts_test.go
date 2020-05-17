package simplify_debts

import (
	"fmt"
	"testing"

	"github.com/splitwise/manager"
)

//e.AddUser("1", "Alice", "xyz1@gmail.com", "123")
//e.AddUser("2", "Bob", "xyz2@gmail.com", "123")
//e.AddUser("3", "Charlie", "xyz3@gmail.com", "123")
//e.AddUser("4", "David", "xyz4@gmail.com", "123")
//e.AddUser("5", "Ema", "xyz4@gmail.com", "123")
//e.AddUser("6", "Fred", "xyz4@gmail.com", "123")
//e.AddUser("7", "Gabe", "xyz4@gmail.com", "123")

func TestSimplifyDebts(t *testing.T) {
	expenseManager := manager.ExpenseManager{
		BalanceSheet: map[string]map[string]float64{
			"1": nil,
			"2": {
				"7": 30,
				"3": -40,
				"6": 10,
			},
			"3": {
				"2": 40,
				"4": -20,
				"6": 30,
			},
			"4": {
				"7": 10,
				"3": 20,
				"5": -50,
				"6": 10,
			},
			"5": {
				"4": 50,
				"6": 10,
			},
			"6": {
				"5": -10,
				"2": -10,
				"3": -30,
				"4": -10,
			},
			"7": map[string]float64{
				"2": -30,
				"4": -10,
			},
		},
	}

	residualGraph := CopyToResidualGraph(expenseManager.BalanceSheet)
	fmt.Printf("residualGraph ------- %+v", residualGraph)
	FordFulkerson([]string{"2", "3", "4", "5", "6", "7"}, residualGraph)

}

func TestABC(t *testing.T) {
	input := [][]int{{1, 2}, {3}, {3}, {}}
	allPathsSourceTarget(input)
}

func allPathsSourceTarget(graph [][]int) [][]int {
	len := len(graph)
	start := 0
	end := len - 1
	return dfs(start, end, graph)
}

func dfsUtils(start, end int, graph [][]int, path []int, paths *[][]int, visited []bool) {
	visited[start] = true
	path = append(path, start)

	if start == end {
		*paths = append(*paths, path)
	}

	for i := 0; i < len(graph[start]); i++ {
		if !visited[graph[start][i]] {
			dfsUtils(graph[start][i], end, graph, path, paths, visited)
		}
	}

	visited[start] = false
	return

}

func dfs(start, end int, graph [][]int) [][]int {
	visited := make([]bool, end+1)
	paths := make([][]int, 0)
	path := make([]int, 0)
	dfsUtils(start, end, graph, path, &paths, visited)
	return paths
}
