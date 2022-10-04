package routes

import (
	"github.com/dmedinao1/ejercicio-TM-04-10/models"
	"github.com/dmedinao1/ejercicio-TM-04-10/services"
	"github.com/gin-gonic/gin"
	"log"
)

func DefineTransactionsRoutes(group *gin.RouterGroup) {
	group.POST("/", createTransaction)
	group.GET("/", getAllTransactions)
}

func hasAValidToken(c *gin.Context) bool {
	token := c.GetHeader("token")
	return token == "1234"
}

func createTransaction(c *gin.Context) {

	if !hasAValidToken(c) {
		c.JSON(401, gin.H{
			"error": "Invalid token",
		})

		return
	}

	var toSave models.Transaction

	err := c.ShouldBind(&toSave)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	service := services.GetTransactionService()

	saved := service.SaveTransaction(&toSave)

	log.Printf("To save address: %p | Saved address: %v\n", &toSave, saved)

	c.JSON(200, *saved)
}

func getAllTransactions(c *gin.Context) {
	if !hasAValidToken(c) {
		c.JSON(401, gin.H{
			"error": "Invalid token",
		})

		return
	}

	service := services.GetTransactionService()

	allTransactions := service.GetAll()

	c.JSON(200, allTransactions)
}
