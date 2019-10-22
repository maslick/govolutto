package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/maslick/govolutto/src"
	"log"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	var service = CreateService()
	log.Fatal(SetupRouter(service).Run())
}
