package test

import (
	. "github.com/maslick/govolutto/src"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

type falseWithdrawalRepo struct{}

func (repo *falseWithdrawalRepo) Withdraw(username string, amount decimal.Decimal) bool { return false }
func (repo *falseWithdrawalRepo) Deposit(username string, amount decimal.Decimal) bool  { return false }
func (repo *falseWithdrawalRepo) GetBalance(username string) decimal.Decimal {
	return decimal.NewFromFloat(0)
}

func TestFalseWithdrawal(t *testing.T) {
	var transaction = TransactionProducer(&falseWithdrawalRepo{})
	result := transaction.Transfer("daisy", "donald", decimal.NewFromFloat(-50))
	assert.False(t, result)
}

func TestWithdrawalAmount(t *testing.T) {
	var repo = RepoProducer()
	result := repo.Withdraw("donald", decimal.NewFromFloat(50))
	assert.False(t, result)
}

func TestUserNotExist(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("TestUserNotExist should have panicked!")
			}
		}()

		var repo = RepoProducer()
		repo.GetBalance("helloworld")
	}()
}
