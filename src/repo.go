package src

import (
	"errors"
	"fmt"
	. "github.com/shopspring/decimal"
)

type DummyRepo struct {
	users map[string]Account
}

func (it *DummyRepo) Withdraw(username string, amount Decimal) bool {
	user := it.findUser(username)
	if amount.GreaterThan(user.balance) {
		return false
	}
	it.users[username] = Account{username: user.username, balance: user.balance.Sub(amount), fullname: user.fullname}
	return true
}

func (it *DummyRepo) Deposit(username string, amount Decimal) bool {
	return it.Withdraw(username, amount.Neg())
}

func (it *DummyRepo) GetBalance(username string) Decimal {
	return it.findUser(username).balance
}

func (it *DummyRepo) findUser(userId string) *Account {
	account, ok := it.users[userId]
	if !ok {
		panic(errors.New(fmt.Sprintf("account with username=%s not found", userId)))
	}
	return &account
}
