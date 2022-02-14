package main

import (
	"fcc-request-header-parser-ms/handlers"
	"fcc-request-header-parser-ms/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.CORS())

	r.GET("/api/whoami", handlers.WhoAmIHandler)

	if err := r.Run(":3000"); err != nil {
		log.Fatal("can't connect to server at port 3000")
	}
}
