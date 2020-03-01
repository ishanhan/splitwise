package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/splitwise/dto"

	"github.com/splitwise/manager"
)

// Commands

const (
	show    = "SHOW"
	expense = "EXPENSE"
)

//EXPENSE <user-id-of-person-who-paid> <no-of-users> <space-separated-list-of-users> <EQUAL/EXACT/PERCENT> <space-separated-values-in-case-of-non-equal>

func main() {
	expenseManager := &manager.ExpenseManager{}
	initBalanceSheet(expenseManager)
	addUser(expenseManager)
	fetchCommands(expenseManager)
}

func initBalanceSheet(e *manager.ExpenseManager) {
	e.BalanceSheet = make(map[string]map[string]float64)
}

func addUser(e *manager.ExpenseManager) {
	e.AddUser("1", "Ishan", "xyz1@gmail.com", "123")
	e.AddUser("2", "Sidhima", "xyz2@gmail.com", "123")
	e.AddUser("3", "Pritika", "xyz3@gmail.com", "123")
	e.AddUser("4", "Anurag", "xyz4@gmail.com", "123")

}

func fetchCommands(e *manager.ExpenseManager) {

	in := bufio.NewReader(os.Stdin)
	for {
		input, err := in.ReadString('\n')
		if err != nil {
			// io.EOF is expected, anything else
			// should be handled/reported
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		input = strings.Replace(input, "\n", "", -1)
		expense := strings.Split(input, " ")
		command := expense[0]
		err = takeAction(command, expense, e)
		if err != nil {
			fmt.Printf("Error occured in the application %v", err)
		}
	}
}

func takeAction(command string, expenses []string, e *manager.ExpenseManager) (err error) {
	switch command {
	case show:
		if len(expenses) == 2 {
			var user dto.User
			for _, u := range e.Users {
				if u.ID == expenses[1] {
					user = u
				}
			}
			e.ShowUserBalances(user)
		} else {
			e.ShowBalances()
		}
		return nil
	case expense:
		var (
			paidBy    string
			amount    string
			noOfUsers string
			users     []string
			operation string
			division  []string
		)

		if len(expenses) == 7 {
			division = strings.Split(expenses[6], ",")
		}

		if len(expenses) == 6 || len(expenses) == 7 {
			paidBy = expenses[1]
			amount = expenses[2]
			noOfUsers = expenses[3]
			users = strings.Split(expenses[4], ",")
			operation = expenses[5]
			return e.AddExpense(paidBy, amount, noOfUsers, users, operation, division)
		}
		return errors.New("parameters in expense command are not proper")
	default:
		return errors.New("unrecognized command")
	}
}
