package src

import (
	. "github.com/shopspring/decimal"
	"sync"
)

type Transaction struct {
	Repo IRepo
	sync.Mutex
}

func (it *Transaction) Transfer(from string, to string, amount Decimal) bool {
	it.Lock()
	defer it.Unlock()

	if from == to {
		return false
	}
	if !it.Repo.Withdraw(from, amount.Abs()) {
		return false
	}
	return it.Repo.Deposit(to, amount.Abs())
}
