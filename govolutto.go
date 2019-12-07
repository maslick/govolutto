//+build !test

package main

import (
	. "github.com/maslick/govolutto/src"
	"log"
)

func init() {
	go NewMetrics(5)
}

func main() {
	log.Fatal(SetupRouter(CreateService()).Run())
}
