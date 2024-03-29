//+build wireinject

package src

import (
	"github.com/google/wire"
	. "github.com/shopspring/decimal"
)

func RepoProducer() IRepo {
	return &DummyRepo{users: map[string]Account{
		"donald":  {"donald", NewFromFloat(0), "Donald Duck"},
		"daisy":   {"daisy", NewFromFloat(100), "Daisy Duck"},
		"scrooge": {"scrooge", NewFromFloat(10000), "Scrooge McDuck"},
		"gyro":    {"gyro", NewFromFloat(10000), "Gyro Gearloose"},
	}}
}

func TransactionProducer(repo IRepo) ITransaction {
	return &Transaction{Repo: repo}
}

func BalanceProducer(repo IRepo) IBalance {
	return &Balance{repo}
}

func ServiceProducer(transaction ITransaction, balance IBalance) *Service {
	return &Service{transaction, balance}
}

func CreateService() *Service {
	wire.Build(RepoProducer, TransactionProducer, BalanceProducer, ServiceProducer)
	return &Service{}
}
