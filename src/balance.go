package src

import . "github.com/shopspring/decimal"

type Balance struct {
	Repo IRepo
}

func (it *Balance) Amount(userId string) Decimal {
	return it.Repo.getBalance(userId)
}
