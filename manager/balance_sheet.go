package manager

func (e *ExpenseManager) updateBalanceSheet(paidBy string, amount float64, paidTo string) {
	//update Balance Sheet

	// updating balance sheet for payer
	if _, ok := e.BalanceSheet[paidBy]; !ok {
		e.BalanceSheet[paidBy] = make(map[string]float64)
	}

	if _, ok := e.BalanceSheet[paidBy][paidTo]; !ok {
		e.BalanceSheet[paidBy][paidTo] = 0.0
	}
	e.BalanceSheet[paidBy][paidTo] += amount

	// updating balance sheet for payee
	if _, ok := e.BalanceSheet[paidTo]; !ok {
		e.BalanceSheet[paidTo] = make(map[string]float64)
	}

	if _, ok := e.BalanceSheet[paidTo][paidBy]; !ok {
		e.BalanceSheet[paidTo][paidBy] = 0.0
	}
	e.BalanceSheet[paidTo][paidBy] -= amount
}
