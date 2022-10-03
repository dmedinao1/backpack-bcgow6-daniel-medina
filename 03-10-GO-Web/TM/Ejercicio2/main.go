package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		log.Println("Starting hello world handler")
		c.JSON(200, gin.H{
			"message": "Hola Daniel",
		})
	})

	router.Run()
}
