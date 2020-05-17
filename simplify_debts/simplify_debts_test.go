package simplify_debts

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
	ExpectedBalanceSheet := map[string]map[string]ResidualValues{
		"1": map[string]ResidualValues{},
		"2": {
			"3": {
				30,
				false,
			},
		},
		"3": map[string]ResidualValues{},
		"4": {
			"5": {
				20,
				false,
			},
		},
		"5": map[string]ResidualValues{},
		"6": {
			"5": {
				40,
				false,
			},
			"3": {20,
				false,
			},
		},
		"7": {
			"2": {
				30,
				false,
			},
			"4": {
				10,
				false,
			},
		},
	}
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
			"7": {
				"2": -30,
				"4": -10,
			},
		},
	}

	residualGraph := CopyToResidualGraph(expenseManager.BalanceSheet)

	var edges []Edge
	for key, value := range residualGraph {
		for subKey, _ := range value {
			edges = append(edges, Edge{key, subKey})
		}
	}

	SimplifyDebts(edges, residualGraph)
	for _, value := range residualGraph {
		for subKey, subValue := range value {
			if subValue.flow == 0 {
				delete(value, subKey)
			}
		}
	}
	assert.Equal(t, ExpectedBalanceSheet, residualGraph)
}
