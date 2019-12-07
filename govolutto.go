//+build !test

package main

import (
	. "github.com/maslick/govolutto/src"
	"log"
)

func main() {
	log.Fatal(SetupRouter(CreateService()).Run())
}
