package src

import . "github.com/shopspring/decimal"

type ITransaction interface {
	Transfer(from string, to string, amount Decimal) bool
}

type IBalance interface {
	Amount(userId string) Decimal
}

type IRepo interface {
	withdraw(username string, amount Decimal) bool
	deposit(username string, amount Decimal) bool
	getBalance(username string) Decimal
}

type Account struct {
	username string
	balance  Decimal
	fullname string
}
