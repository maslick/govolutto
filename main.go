//+build !test

package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/maslick/govolutto/src"
	"log"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	go NewMetrics(5)
}

func main() {
	var service = CreateService()
	log.Fatal(SetupRouter(service).Run())
}
