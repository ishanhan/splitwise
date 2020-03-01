package dto

type User struct {
	ID      string
	Name    string
	Email   string
	PhoneNo string
}

type Expense struct {
	Amount float64
	PaidBy string
	Splits []Split
}

type Split struct {
	UserID   string
	Division interface{}
}

type ExactSplit struct {
	Amount float64
}

type PercentSplit struct {
	Percent int64
}

type EqualSplit struct {
	Amount float64
}
