package src

import (
	. "github.com/shopspring/decimal"
	"sync"
)

type Transaction struct {
	Repo IRepo
	Lock sync.Mutex
}

func DefaultTransaction(repo *IRepo) ITransaction {
	var transaction = new(Transaction)
	transaction.Repo = *repo
	return transaction
}

func (it *Transaction) Transfer(from string, to string, amount Decimal) bool {
	it.Lock.Lock()
	defer it.Lock.Unlock()

	if from == to {
		return false
	}
	if !it.Repo.withdraw(from, amount.Abs()) {
		return false
	}
	return it.Repo.deposit(to, amount.Abs())
}
