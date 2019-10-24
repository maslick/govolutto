package src

import (
	. "github.com/shopspring/decimal"
)

func CreateRepo() DummyRepo {
	return DummyRepo{users: map[string]Account{
		"donald":  {"donald", NewFromFloat(0), "Donald Duck"},
		"daisy":   {"daisy", NewFromFloat(100), "Daisy Duck"},
		"scrooge": {"scrooge", NewFromFloat(10000), "Scrooge McDuck"},
		"gyro":    {"gyro", NewFromFloat(10000), "Gyro Gearloose"},
	}}
}

func CreateTransaction(repo IRepo) Transaction {
	return Transaction{Repo: repo}
}

func CreateBalance(repo IRepo) Balance {
	return Balance{repo}
}

func CreateService() *Service {
	repo := CreateRepo()
	transaction := CreateTransaction(&repo)
	balance := CreateBalance(&repo)

	return &Service{
		Transaction: &transaction,
		Balance:     &balance,
	}
}
