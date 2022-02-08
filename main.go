package main

import (
	"fcc-request-header-parser-ms/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/api/whoami", handlers.WhoAmIHandler)
	r.Run(":3000")
}
