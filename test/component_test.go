package test

import (
	. "github.com/maslick/govolutto/src"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestDummyTransfer(t *testing.T) {
	var repo IRepo = NewDummyRepo()
	var transaction = DefaultTransaction(&repo)
	var balance IBalance = &Balance{Repo: repo}

	assert.Equal(t, "100", balance.Amount("daisy").String())
	assert.Equal(t, "0", balance.Amount("donald").String())

	transaction.Transfer("daisy", "donald", decimal.NewFromFloat(50))

	assert.Equal(t, "50", balance.Amount("daisy").String())
	assert.Equal(t, "50", balance.Amount("donald").String())
}

func TestNegativeTransfer(t *testing.T) {
	var repo IRepo = NewDummyRepo()
	var transaction = DefaultTransaction(&repo)
	var balance IBalance = &Balance{Repo: repo}

	transaction.Transfer("daisy", "donald", decimal.NewFromFloat(-50))
	assert.Equal(t, "50", balance.Amount("daisy").String(), "Daisy should have 50")
	assert.Equal(t, "50", balance.Amount("donald").String(), "Donalnd should have 50")
}

func TestConcurrentTransfer(t *testing.T) {
	var wg sync.WaitGroup
	var count = 100
	var amount = decimal.NewFromFloat(float64(100.0 / float32(count)))
	wg.Add(count)

	var repo IRepo = NewDummyRepo()
	var transaction = DefaultTransaction(&repo)
	var balance IBalance = &Balance{Repo: repo}

	for i := 0; i < count; i++ {
		go func() {
			transaction.Transfer("daisy", "donald", amount)
			wg.Done()
		}()
	}

	wg.Wait()

	time.Sleep(3000)

	assert.Equal(t, "0", balance.Amount("daisy").String())
	assert.Equal(t, "100", balance.Amount("donald").String())
}
