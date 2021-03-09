package main

import (
	"github.com/gin-gonic/gin"
	// "strconv"
)

// CreditCard ...
type CreditCard struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Money int    `json:"money"`
}

var CreditCards []CreditCard

func getCreditCard(c *gin.Context) {
	for _, item := range CreditCards {
		c.JSON(200, gin.H{
			"id":    item.ID,
			"name":  item.Name,
			"money": item.Money,
		})
	}

	c.JSON(404, gin.H{
		"message": "No encontrado",
	})
}

func main() {
	CreditCards = append(CreditCards, CreditCard{
		ID:    1,
		Name:  "Paty maldonado",
		Money: 20000,
	})
	CreditCards = append(CreditCards, CreditCard{
		ID:    2,
		Name:  "Alexis Sanchez",
		Money: 406000570,
	})
	CreditCards = append(CreditCards, CreditCard{
		ID:    3,
		Name:  "Augusto Pinochet",
		Money: 1234567890,
	})

	r := gin.Default()
	userRoutes := r.Group("/api/examen")
	{
		userRoutes.GET("/CreditCard", getCreditCard)
		userRoutes.POST("/CreditCard", createCreditCard)
		userRoutes.PATCH("/:id", updateCreditCard)
		userRoutes.DELETE("/:id", deleteCreditCard)
	}

	r.Run()
}
