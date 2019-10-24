package src

import (
	. "github.com/shopspring/decimal"
	"sync"
)

type Transaction struct {
	Repo IRepo
	Lock sync.Mutex
}

func (it *Transaction) Transfer(from string, to string, amount Decimal) bool {
	it.Lock.Lock()
	defer it.Lock.Unlock()

	if from == to {
		return false
	}
	if !it.Repo.Withdraw(from, amount.Abs()) {
		return false
	}
	return it.Repo.Deposit(to, amount.Abs())
}
