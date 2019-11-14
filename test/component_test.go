package test

import (
	. "github.com/maslick/govolutto/src"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestDummyTransfer(t *testing.T) {
	var repo = RepoProducer()
	var transaction = TransactionProducer(repo)
	var balance IBalance = &Balance{Repo: repo}

	assert.Equal(t, "100", balance.Amount("daisy").String())
	assert.Equal(t, "0", balance.Amount("donald").String())

	transaction.Transfer("daisy", "donald", decimal.NewFromFloat(50))

	assert.Equal(t, "50", balance.Amount("daisy").String())
	assert.Equal(t, "50", balance.Amount("donald").String())
}

func TestInsufficientFundsTransfer(t *testing.T) {
	var repo = RepoProducer()
	var transaction = TransactionProducer(repo)
	var balance IBalance = &Balance{Repo: repo}

	result := transaction.Transfer("daisy", "daisy", decimal.NewFromFloat(110))

	assert.Equal(t, "100", balance.Amount("daisy").String())
	assert.Equal(t, "0", balance.Amount("donald").String())
	assert.False(t, result)
}

func TestTransferToMyself(t *testing.T) {
	var repo = RepoProducer()
	var transaction = TransactionProducer(repo)
	var balance IBalance = &Balance{Repo: repo}

	result := transaction.Transfer("daisy", "daisy", decimal.NewFromFloat(50))

	assert.Equal(t, "100", balance.Amount("daisy").String())
	assert.Equal(t, "0", balance.Amount("donald").String())
	assert.False(t, result)
}

func TestNegativeTransfer(t *testing.T) {
	var repo = RepoProducer()
	var transaction = TransactionProducer(repo)
	var balance IBalance = &Balance{Repo: repo}

	result := transaction.Transfer("daisy", "donald", decimal.NewFromFloat(-50))
	assert.Equal(t, "50", balance.Amount("daisy").String(), "Daisy should have 50")
	assert.Equal(t, "50", balance.Amount("donald").String(), "Donalnd should have 50")
	assert.True(t, result)
}

func TestConcurrentTransfer(t *testing.T) {
	var wg sync.WaitGroup
	var count = 100
	var amount = decimal.NewFromFloat(float64(100.0 / float32(count)))
	wg.Add(count)

	var repo = RepoProducer()
	var transaction = TransactionProducer(repo)
	var balance IBalance = &Balance{Repo: repo}

	for i := 0; i < count; i++ {
		go func() {
			transaction.Transfer("daisy", "donald", amount)
			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, "0", balance.Amount("daisy").String())
	assert.Equal(t, "100", balance.Amount("donald").String())
}
