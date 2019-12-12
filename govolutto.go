//+build !test

package main

import (
	. "github.com/maslick/govolutto/src"
)

func main() {
	api := RestAPI{Service: *CreateService()}
	api.Start()
}
