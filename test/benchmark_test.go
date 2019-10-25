package test

import (
	. "github.com/maslick/govolutto/src"
	"github.com/shopspring/decimal"
	"testing"
)

func BenchmarkTransaction(b *testing.B) {
	var repo = CreateRepo()
	var transaction = CreateTransaction(&repo)

	for n := 0; n < b.N; n++ {
		transaction.Transfer("scrooge", "donald", decimal.NewFromFloat(1))
	}
}

func BenchmarkBalanceCheck(b *testing.B) {
	var repo = CreateRepo()
	var balance IBalance = &Balance{Repo: &repo}

	for n := 0; n < b.N; n++ {
		balance.Amount("scrooge")
	}
}
