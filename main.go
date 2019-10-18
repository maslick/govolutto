package main

import (
	"fmt"
	. "github.com/maslick/govolutto/src"
	. "github.com/shopspring/decimal"
)

func main() {
	var repo IRepo = NewDummyRepo()
	var transaction = DefaultTransaction(&repo)
	var balance IBalance = &Balance{Repo: repo}

	fmt.Println("Daisy's balance: " + balance.Amount("daisy").String())
	fmt.Println("Donald's balance: " + balance.Amount("donald").String())

	transaction.Transfer("daisy", "donald", NewFromFloat(50))

	fmt.Println("Daisy's balance: " + balance.Amount("daisy").String())
	fmt.Println("Donald's balance: " + balance.Amount("donald").String())
}
