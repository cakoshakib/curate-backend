package main

import (
	"github.com/cakoshakib/curate-backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", handlers.HelloHandler)
	router.Run(":8080")
}
