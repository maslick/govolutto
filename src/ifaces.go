package src

import . "github.com/shopspring/decimal"

type ITransaction interface {
	Transfer(from string, to string, amount Decimal) bool
}

type IBalance interface {
	Amount(userId string) Decimal
}

type IRepo interface {
	Withdraw(username string, amount Decimal) bool
	Deposit(username string, amount Decimal) bool
	GetBalance(username string) Decimal
}

type Service struct {
	Transaction ITransaction
	Balance     IBalance
}
