package manager

import (
	"strconv"

	"github.com/splitwise/dto"
)

func (e *ExpenseManager) createEqualExpense(paidBy string, amount float64, noOfUsers int, users []string) error {
	expense := initializeExpense(paidBy, amount)

	splitAmount := amount / float64(noOfUsers)

	for _, userID := range users {
		split := dto.Split{
			UserID:   userID,
			Division: dto.EqualSplit{Amount: splitAmount},
		}
		expense.Splits = append(expense.Splits, split)
		e.updateBalanceSheet(paidBy, splitAmount, userID)

	}
	e.Expenses = append(e.Expenses, expense)
	return nil
}

func (e *ExpenseManager) createExactExpense(paidBy string, amount float64, noOfUsers int, users []string, division []string) error {
	expense := initializeExpense(paidBy, amount)

	for index, userID := range users {
		splitAmount, err := strconv.ParseFloat(division[index], 64)
		if err != nil {
			return err
		}

		split := dto.Split{
			UserID:   userID,
			Division: dto.EqualSplit{Amount: splitAmount},
		}
		expense.Splits = append(expense.Splits, split)
		e.updateBalanceSheet(paidBy, splitAmount, userID)

	}
	e.Expenses = append(e.Expenses, expense)
	return nil
}

func (e *ExpenseManager) createPercentExpense(paidBy string, amount float64, noOfUsers int, users []string, division []string) error {
	expense := initializeExpense(paidBy, amount)

	for index, userID := range users {
		splitPercent, err := strconv.ParseFloat(division[index], 64)
		if err != nil {
			return err
		}
		splitAmount := (amount / 100) * splitPercent

		if userID == paidBy {
			continue
		}
		split := dto.Split{
			UserID:   userID,
			Division: dto.EqualSplit{Amount: splitAmount},
		}
		expense.Splits = append(expense.Splits, split)
		e.updateBalanceSheet(paidBy, splitAmount, userID)

	}
	e.Expenses = append(e.Expenses, expense)
	return nil
}

func initializeExpense(paidBy string, amount float64) dto.Expense {
	return dto.Expense{PaidBy: paidBy, Amount: amount}
}
