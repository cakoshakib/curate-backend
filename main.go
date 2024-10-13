package main

import "github.com/gin-gonic/gin"

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, curate!"})
}

func main() {
	router := gin.Default()
	router.GET("/hello", helloHandler)
	router.Run(":8080")
}
