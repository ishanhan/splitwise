package manager

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/splitwise/dto"
)

const (
	// split types
	equal   = "EQUAL"
	exact   = "EXACT"
	percent = "PERCENT"
)

type ExpenseManager struct {
	Expenses     []dto.Expense
	Users        []dto.User
	BalanceSheet map[string]map[string]float64
}

func (e *ExpenseManager) AddUser(userID string, name string, email string, phoneNo string) {
	user := dto.User{
		ID:      userID,
		Name:    name,
		Email:   email,
		PhoneNo: phoneNo,
	}
	e.Users = append(e.Users, user)
}

func (e *ExpenseManager) AddExpense(paidBy string, amount string, noOfUsers string, users []string, operation string, division []string) error {

	if paidBy == "" || len(users) == 0 || operation == "" {
		return errors.New("invalid expense added")
	}

	total, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return err
	}

	noOfSplits, err := strconv.Atoi(noOfUsers)
	if err != nil {
		return err
	}

	if operation != equal {
		if noOfSplits != len(division) {
			return errors.New("number of splits is not equal to users specified")
		}
	}

	if noOfSplits != len(users) {
		return errors.New("number of splits is not equal to users specified")
	}

	switch operation {
	case equal:
		err := e.createEqualExpense(paidBy, total, noOfSplits, users)
		return err
	case exact:
		err := e.createExactExpense(paidBy, total, noOfSplits, users, division)
		return err
	case percent:
		err := e.createPercentExpense(paidBy, total, noOfSplits, users, division)
		return err
	default:
		return errors.New("invalid expense type added")
	}
}

func (e *ExpenseManager) ShowUserBalances(user dto.User) {

	if e.BalanceSheet[user.ID] == nil {
		fmt.Printf("%v has no balances \n", user.Name)
	}

	for _, otherUser := range e.Users {
		if user == otherUser {
			continue
		}
		amount := e.BalanceSheet[user.ID][otherUser.ID]
		if amount == 0 {
			continue
		}
		if amount < 0 {
			fmt.Printf("%s owes %s: %f \n", user.Name, otherUser.Name, -1*amount)
		} else {
			fmt.Printf("%s gets back from %s: %f \n", user.Name, otherUser.Name, amount)
		}
	}
}

func (e *ExpenseManager) ShowBalances() {
	for _, user := range e.Users {
		totalAmount := 0.0
		if e.BalanceSheet[user.ID] == nil {
			fmt.Printf("%v is all settled up \n", user.Name)
		}

		for _, otherUser := range e.Users {
			if user == otherUser {
				continue
			}
			totalAmount += e.BalanceSheet[user.ID][otherUser.ID]
		}

		if totalAmount == 0 {
			fmt.Printf("%s is all settled up \n", user.Name)
		} else if totalAmount < 0 {
			fmt.Printf("%s owes %f \n", user.Name, -1*totalAmount)
		} else {
			fmt.Printf("%s gets back from %f \n", user.Name, totalAmount)
		}

	}
}
